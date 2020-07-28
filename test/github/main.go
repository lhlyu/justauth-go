package main

import (
	"encoding/json"
	"github.com/lhlyu/justauth-go"
	"github.com/lhlyu/justauth-go/model"
	"log"
	"net/http"
	"net/url"
)

const (
	CLIENT_ID     = "6713787xxxxxxxxxx"               // 自行申请
	CLIENT_SECRET = "99cb9efbe5xxxxxxxxxxxxxxxxxxxxx" // 自行申请
	REDIRECT_URL  = "http://localhost:8080/api/login" // 回调地址

	STATE = "test"
)

var auth justauth.IAuth

func init() {
	auth = justauth.NewAuth("github", model.AuthConfig{
		ClientId:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		RedirectUrl:  REDIRECT_URL,
	})
}

func main() {
	// http://localhost:8080/     // 返回授权的地址
	http.HandleFunc("/", Authorize)
	// 授权成功回调地址，登陆成功会返回用户的信息
	http.HandleFunc("/api/login", Login)

	log.Println("localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}

// 生成授权URL
func Authorize(w http.ResponseWriter, r *http.Request) {
	rs := auth.AuthorizeWithState(STATE)
	w.Write([]byte(rs.Val()))
}

// 登录并返回用户的个人信息
func Login(w http.ResponseWriter, r *http.Request) {
	u, _ := url.ParseRequestURI(r.RequestURI)
	q := u.Query()
	// 获取参数
	code := q.Get("code")
	state := q.Get("state")
	rs := auth.Login(model.Callback{
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
