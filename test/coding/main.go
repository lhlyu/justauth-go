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
	CLIENT_ID     = "259880f51e07f079d6f1d492bd2619723"         // 自行申请
	CLIENT_SECRET = "c59978915682568f9c6e59b91a44e9cc7be7a2f61" // 自行申请
	REDIRECT_URL  = "http://xxxx.com/login"

	STATE = "test"
)

// coding 已改版 todo
var authRequest request.AuthRequest

func init() {
	authRequest, _ = request.NewAuthRequest(config.AuthConfig{
		ClientId:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		RedirectUrl:  REDIRECT_URL,
	}, source.CODING)
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
