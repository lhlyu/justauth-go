# justauth-go
第三方平台的授权登录

#### 开发进度

- 进行中...

## 例子

- github

```go
// 生成授权URL
func TestAuthorize(t *testing.T) {
	authRequest, reqRs := request.NewAuthRequest(config.AuthConfig{
		ClientId:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		RedirectUrl:  REDIRECT_URL,
	}, source.GITHUB)
	if !reqRs.Ok() {
		t.Error(reqRs.Msg())
		return
	}
	rs := authRequest.AuthorizeWithState(STATE)
	if !rs.Ok() {
		t.Error(rs.Msg())
		return
	}
	t.Log(rs.Val())
}

// 登录获取用户信息
func TestLogin(t *testing.T) {
	authRequest, reqRs := request.NewAuthRequest(config.AuthConfig{
		ClientId:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		RedirectUrl:  REDIRECT_URL,
	}, source.GITHUB)
	if !reqRs.Ok() {
		t.Error(reqRs.Msg())
		return
	}
	rs := authRequest.Login(&model.Callback{
		Code:  CODE,
		State: STATE,
	})
	if !rs.Ok() {
		t.Error(rs.Msg())
		return
	}
	bts, _ := json.Marshal(rs.Val())
	t.Log(string(bts))
}
```