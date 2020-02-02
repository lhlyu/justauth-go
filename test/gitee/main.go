package main

import (
	"encoding/json"
	"github.com/lhlyu/justauth-go/config"
	"github.com/lhlyu/justauth-go/entity"
	"github.com/lhlyu/justauth-go/request"
	"github.com/lhlyu/justauth-go/source"
	"log"
	"net/http"
	"net/url"
)

const (
	CLIENT_ID     = "a71fffee035f0ec7dc3689651561df09a30fbace6f956f31a9e67bb77de55de14" // 自行申请
	CLIENT_SECRET = "024de2d0fc304b0294a1fc2cead10a45b9b0f18cf6af2dcafce73729171813d4d" // 自行申请
	REDIRECT_URL  = "http://localhost:8080/login"                                       // 回调地址

	STATE = "test"
)

var authRequest request.AuthRequest

func init() {
	// 初始化
	authRequest, _ = request.NewAuthRequest(config.AuthConfig{
		ClientId:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		RedirectUrl:  REDIRECT_URL,
	}, source.GITEE)
}

func main() {
	// http://localhost:8080/     // 返回授权的地址
	http.HandleFunc("/", Authorize)
	// 授权成功回调地址，登陆成功会返回用户的信息
	http.HandleFunc("/login", Login)
	// 刷新token，获取新的token
	// GET: http://localhost:8080/refresh?refreshToken=
	http.HandleFunc("/refresh", Refresh)

	log.Println("localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}

// 生成授权URL
func Authorize(w http.ResponseWriter, r *http.Request) {
	rs := authRequest.AuthorizeWithState(STATE)
	// 如果失败的话，返回错误信息
	if !rs.Ok() {
		w.Write([]byte(rs.Msg()))
		return
	}
	w.Write([]byte(rs.Val()))
}

// 登录并返回用户的个人信息
func Login(w http.ResponseWriter, r *http.Request) {
	u, _ := url.ParseRequestURI(r.RequestURI)
	q := u.Query()
	// 获取参数
	code := q.Get("code")
	state := q.Get("state")
	rs := authRequest.Login(&entity.Callback{
		Code:  code,
		State: state,
	})
	// 如果失败的话，返回错误信息
	if !rs.Ok() {
		w.Write([]byte(rs.Msg()))
		return
	}
	// 成功
	bts, _ := json.Marshal(rs.Val())
	w.Write(bts)
}

// GET: http://localhost:8080/refresh?refreshToken=
func Refresh(w http.ResponseWriter, r *http.Request) {
	u, _ := url.ParseRequestURI(r.RequestURI)
	q := u.Query()
	// 获取参数
	refreshToken := q.Get("refreshToken")

	rs := authRequest.Refresh(&entity.Token{
		RefreshToken: refreshToken,
	})
	if !rs.Ok() {
		w.Write([]byte(rs.Msg()))
		return
	}
	bts, _ := json.Marshal(rs.Val())
	w.Write(bts)
}
