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
	CLIENT_ID     = "Iv1.4r3tg45g3g3g"                     // 自行申请
	CLIENT_SECRET = "45g4g4g53456645654645643564646546353" // 自行申请
	REDIRECT_URL  = "http://localhost:8080/login"

	STATE = "test"
)

var authRequest request.AuthRequest

func init() {
	authRequest, _ = request.NewAuthRequest(config.AuthConfig{
		ClientId:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		RedirectUrl:  REDIRECT_URL,
	}, source.GITHUB)
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
	state := q.Get("test") // or STATE
	rs := authRequest.Login(&model.Callback{
		Code:  code,
		State: state,
	})
	bts, _ := json.Marshal(rs.Val())
	w.Write(bts)
}
