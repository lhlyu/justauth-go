package main

import (
	"encoding/json"
	"github.com/lhlyu/justauth-go/config"
	"github.com/lhlyu/justauth-go/model"
	"github.com/lhlyu/justauth-go/request"
	"github.com/lhlyu/justauth-go/source"
	"log"
	"net/http"
	"net/url"
)

const (
	CLIENT_ID     = "a71fffee035f0ec7dc3689651561df09a30fbace6f956f31a9e67bb77de55de1x" // 自行申请
	CLIENT_SECRET = "024de2d0fc304b0294a1fc2cead10a45b9b0f18cf6af2dcafce73729171813d4y" // 自行申请
	REDIRECT_URL  = "http://localhost:8080/login"

	STATE = "test"
)

var authRequest request.AuthRequest

func init() {
	authRequest, _ = request.NewAuthRequest(config.AuthConfig{
		ClientId:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		RedirectUrl:  REDIRECT_URL,
	}, source.GITEE)
}

func main() {
	http.HandleFunc("/", Authorize)
	http.HandleFunc("/login", Login)
	log.Println("localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}

// 生成授权URL
func Authorize(w http.ResponseWriter, r *http.Request) {
	rs := authRequest.AuthorizeWithState(STATE)
	w.Write([]byte(rs.Val()))
}

// 登录并返回用户的个人信息
func Login(w http.ResponseWriter, r *http.Request) {
	u, _ := url.ParseRequestURI(r.RequestURI)
	q := u.Query()
	// 获取参数
	code := q.Get("code")
	state := q.Get("state") // or STATE
	rs := authRequest.Login(&model.Callback{
		Code:  code,
		State: state,
	})
	bts, _ := json.Marshal(rs.Val())
	w.Write(bts)
}
