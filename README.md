# justauth-go
第三方平台的授权登录

#### 开发进度

- 进行中...

## 例子

- github

```go
// 生成授权url
func TestAuthorize(t *testing.T) {
	authRequest,err := request.NewAuthRequest(config.AuthConfig{
		ClientId:     CLIENT_ID, 
		ClientSecret: CLIENT_SECRET,
		RedirectUrl:  REDIRECT_URL,
	},source.GITHUB)
	if err != nil{
		t.Error(err)
		return
	}
	// 生成授权页面
	url, err := authRequest.AuthorizeWithState(STATE)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(url)
}
```

```go
// 登录，返回用户信息
func TestLogin(t *testing.T) {
	authRequest, err := request.NewAuthRequest(config.AuthConfig{
		ClientId:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		RedirectUrl:  REDIRECT_URL,
	}, source.GITHUB)
	if err != nil {
		t.Error(err)
		return
	}
	user, err := authRequest.Login(&model.Callback{
		Code:  CODE,
		State: STATE,
	})
	if err != nil {
		t.Error(err)
		return
	}
	bts, _ := json.Marshal(user)
	t.Log(string(bts))
}
```