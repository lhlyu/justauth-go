package request

import (
	"github.com/lhlyu/justauth-go/config"
	"github.com/lhlyu/justauth-go/enum"
	"github.com/lhlyu/justauth-go/model"
	"github.com/lhlyu/justauth-go/result"
	"github.com/lhlyu/justauth-go/source"
	"github.com/lhlyu/justauth-go/utils"
)

type githubRequest struct {
	BaseRequest
}

func newGithubRequest(cfg config.AuthConfig, src source.AuthSource) AuthRequest {
	var authRequest AuthRequest
	authRequest = &githubRequest{
		BaseRequest: BaseRequest{
			Source: src,
			Config: cfg,
		},
	}
	return authRequest
}

// Override 返回授权url
func (this *githubRequest) Authorize() *result.UrlResult {
	return this.AuthorizeWithState("")
}

// Override 返回授权url + state
func (this *githubRequest) AuthorizeWithState(state string) *result.UrlResult {
	url := utils.NewUrlBuilder(this.Source.Authorize()).
		AddParam("response_type", "code").
		AddParam("client_id", this.Config.ClientId).
		AddParam("redirect_uri", this.Config.RedirectUrl).
		AddParam("state", this.GetState(state)).Build()
	return result.Success.WithVal(url).ToUrlResult()
}

// Override 登录返回用户信息
func (this *githubRequest) Login(callback *model.Callback) *result.UserResult {
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

// ------------------------------------------------------------------

func (this *githubRequest) getAccessToken(callback *model.Callback) *result.Result {
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
	return this.getToken(body)
}

func (this *githubRequest) getUserInfo(authToken *model.AuthToken) *result.Result {
	url := utils.NewUrlBuilder(this.Source.UserInfo()).
		AddParam("access_token", authToken.AccessToken).Build()
	body, err := utils.Get(url)
	if err != nil {
		return result.Failure.WithErr(err)
	}
	m := utils.JsonToMSS(body)
	if _, ok := m["error"]; ok {
		return result.Failure.WithMsg(m["error_description"])
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

func (this *githubRequest) getToken(body string) *result.Result {
	m := utils.StrToMSS(body)
	if _, ok := m["error"]; ok {
		return result.Failure.WithMsg(m["error_description"])
	}
	token := &model.AuthToken{
		AccessToken: m["access_token"],
		Scope:       m["scope"],
		TokenType:   m["token_type"],
	}
	return result.Success.WithVal(token)
}
