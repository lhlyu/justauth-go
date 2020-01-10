package request

import (
	"github.com/lhlyu/justauth-go/config"
	"github.com/lhlyu/justauth-go/enum"
	"github.com/lhlyu/justauth-go/model"
	"github.com/lhlyu/justauth-go/result"
	"github.com/lhlyu/justauth-go/source"
	"github.com/lhlyu/justauth-go/utils"
)

type alipayRequest struct {
	BaseRequest
}

// todo
func newAlipayRequest(cfg config.AuthConfig, src source.AuthSource) AuthRequest {
	var authRequest AuthRequest
	authRequest = &alipayRequest{
		BaseRequest: BaseRequest{
			Source: src,
			Config: cfg,
		},
	}
	return authRequest
}

// Override 返回授权url，可自行跳转页面
func (this *alipayRequest) Authorize() *result.UrlResult {
	return this.AuthorizeWithState("")
}

// Override
func (this *alipayRequest) AuthorizeWithState(state string) *result.UrlResult {
	url := utils.NewUrlBuilder(this.Source.Authorize()).
		AddParam("scope", "auth_user").
		AddParam("app_id", this.Config.ClientId).
		AddParam("redirect_uri", this.Config.RedirectUrl).
		AddParam("state", this.GetState(state)).Build()
	return result.Success.WithVal(url).ToUrlResult()
}

// todo 需要接入ali sdk
// Override 统一的登录入口
func (this *alipayRequest) Login(callback *model.Callback) *result.UserResult {
	rs := this.getAccessToken(callback)
	if !rs.Ok() {
		return rs.ToUserResult()
	}
	rs = this.getUserInfo(rs.ToTokenResult().Val())
	if !rs.Ok() {
		return rs.ToUserResult()
	}
	return rs.ToUserResult()
}

// todo 需要接入ali sdk
func (this *alipayRequest) getAccessToken(callback *model.Callback) *result.Result {
	url := utils.NewUrlBuilder(this.Source.AccessToken()).
		AddParam("code", callback.Code).
		AddParam("client_id", this.Config.ClientId).
		AddParam("client_secret", this.Config.ClientSecret).
		AddParam("grant_type", "authorization_code").
		AddParam("redirect_uri", this.Config.RedirectUrl).Build()
	body, err := utils.Post(url)
	if err != nil {
		return result.Failure.WithErr(err)
	}
	m := utils.StrToMSS(body)
	if _, ok := m["error"]; ok {
		desc := m["error_description"]
		return result.Failure.WithMsg(desc)
	}
	token := &model.AuthToken{
		AccessToken: m["access_token"],
		Scope:       m["scope"],
		TokenType:   m["token_type"],
	}
	return result.Success.WithVal(token)
}

// todo 需要接入ali sdk
func (this *alipayRequest) getUserInfo(authToken *model.AuthToken) *result.Result {
	url := utils.NewUrlBuilder(this.Source.UserInfo()).
		AddParam("access_token", authToken.AccessToken).Build()
	body, err := utils.Get(url)
	if err != nil {
		return result.Failure.WithErr(err)
	}
	m := utils.JsonToMSS(body)
	if _, ok := m["error"]; ok {
		desc := m["error_description"]
		return result.Failure.WithMsg(desc)
	}
	user := &model.AuthUser{
		UUID:     m["id"],
		UserName: m["login"],
		Avatar:   m["avatar_url"],
		Blog:     m["blog"],
		NickName: m["name"],
		Company:  m["company"],
		Location: m["location"],
		Email:    m["email"],
		Remark:   m["bio"],
		Gender:   enum.GetRealGender("").Desc,
		Token:    authToken,
		Source:   this.Source.ToString(),
	}
	return result.Success.WithVal(user)
}
