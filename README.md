# justauth-go
第三方平台的授权登录

#### 开发进度

- [X] Github
- [X] Gitee
- [X] Gitlab
- [X] Coding
- [X] CSDN
- [X] QQ

## 例子

- github完整例子

```go
package main

import (
	"log"
        "net/url"
        "net/http"
	"encoding/json"
	
	"github.com/lhlyu/justauth-go/config"
	"github.com/lhlyu/justauth-go/model"
	"github.com/lhlyu/justauth-go/request"
	"github.com/lhlyu/justauth-go/source"

)

const (
	CLIENT_ID     = "Iv1.094eec991d1dsdad"                       // 自行申请
	CLIENT_SECRET = "26074c03ea0167590039f3fb175078a14dsadd23"   // 自行申请
	REDIRECT_URL  = "http://localhost:8080/login"                // 授权回调地址

	STATE = "test" // 自定义
)

var authRequest request.AuthRequest

func init(){
	authRequest, _ = request.NewAuthRequest(config.AuthConfig{
		ClientId:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		RedirectUrl:  REDIRECT_URL,
	}, source.GITHUB)    
}

func main(){
	// 创建两个路由
	http.HandleFunc("/",Authorize)
	http.HandleFunc("/login",Login)
	log.Println("localhost:8080")
	// 启动一个服务
	http.ListenAndServe("localhost:8080",nil)
}

// 生成授权URL
func Authorize(w http.ResponseWriter, r *http.Request){
	rs := authRequest.AuthorizeWithState(STATE)
	w.Write([]byte(rs.Val()))
}

// 登录并返回用户的个人信息
func Login(w http.ResponseWriter, r *http.Request){
	u,_ := url.ParseRequestURI(r.RequestURI)
	q := u.Query()
	// 获取参数
	code := q.Get("code")
	state := q.Get("test")  // or STATE
	rs := authRequest.Login(&model.Callback{
		Code: code,
		State: state,
	})
	bts,_ := json.Marshal(rs.Val())
	w.Write(bts)
}

```
