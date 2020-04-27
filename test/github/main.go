package main

import (
	"encoding/json"
	"fmt"
	"github.com/lhlyu/justauth-go/config"
	"github.com/lhlyu/justauth-go/entity"
	"github.com/lhlyu/justauth-go/request"
	"github.com/lhlyu/justauth-go/source"
	"log"
	"net/http"
	"net/url"
)

const (
	CLIENT_ID     = "Iv1.094eec991d1d290d2"                     // 自行申请
	CLIENT_SECRET = "26074c03ea0167590039f3fb175078a14d864ce1a" // 自行申请
	REDIRECT_URL  = "http://localhost:8080/login"               // 回调地址

	STATE = "test"
)

var authRequest request.AuthRequest

func init() {
	// 初始化
	authRequest, _ = request.NewAuthRequest(config.AuthConfig{
		ClientId:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		RedirectUrl:  REDIRECT_URL,
	}, source.GITHUB) // 这里指定源
}

func main() {
	// http://localhost:8080/     // 返回授权的地址
	http.HandleFunc("/", Authorize)
	// 授权成功回调地址，登陆成功会返回用户的信息
	http.HandleFunc("/login", Login)

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
		fmt.Println(rs.Err())
		w.Write([]byte(rs.Msg()))
		return
	}
	// 成功
	bts, _ := json.Marshal(rs.Val())
	w.Write(bts)
}
