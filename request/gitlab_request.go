package request

import (
	"github.com/lhlyu/justauth-go/config"
	"github.com/lhlyu/justauth-go/enum"
	"github.com/lhlyu/justauth-go/model"
	"github.com/lhlyu/justauth-go/result"
	"github.com/lhlyu/justauth-go/source"
	"github.com/lhlyu/justauth-go/utils"
)

type gitlabRequest struct {
	BaseRequest
}

func newGitlabRequest(cfg config.AuthConfig, src source.AuthSource) AuthRequest {
	var authRequest AuthRequest
	authRequest = &gitlabRequest{
		BaseRequest: BaseRequest{
			Source: src,
			Config: cfg,
		},
	}
	return authRequest
}

// Override 返回授权url，可自行跳转页面
func (this *gitlabRequest) Authorize() *result.UrlResult {
	return this.AuthorizeWithState("")
}

// Override
func (this *gitlabRequest) AuthorizeWithState(state string) *result.UrlResult {
	url := utils.NewUrlBuilder(this.Source.Authorize()).
		AddParam("scope", "read_user+openid+profile+email").
		AddParam("response_type", "code").
		AddParam("client_id", this.Config.ClientId).
		AddParam("redirect_uri", this.Config.RedirectUrl).
		AddParam("state", this.GetState(state)).Build()
	return result.Success.WithVal(url).ToUrlResult()
}

// Override 统一的登录入口
func (this *gitlabRequest) Login(callback *model.Callback) *result.UserResult {
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

func (this *gitlabRequest) getAccessToken(callback *model.Callback) *result.Result {
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
		AccessToken:  m["access_token"],
		Scope:        m["scope"],
		TokenType:    m["token_type"],
		RefreshToken: m["refresh_token"],
		IdToken:      m["id_token"],
	}
	return result.Success.WithVal(token)
}

func (this *gitlabRequest) getUserInfo(authToken *model.AuthToken) *result.Result {
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
		UserName: m["username"],
		Avatar:   m["avatar_url"],
		Blog:     m["web_url"],
		NickName: m["name"],
		Company:  m["organization"],
		Location: m["location"],
		Email:    m["email"],
		Remark:   m["bio"],
		Gender:   enum.GetRealGender("").Desc,
		Token:    authToken,
		Source:   this.Source.ToString(),
	}
	return result.Success.WithVal(user)
}
