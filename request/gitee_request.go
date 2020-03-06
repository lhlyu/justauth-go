package request

import (
	"github.com/lhlyu/justauth-go/config"
	"github.com/lhlyu/justauth-go/entity"
	"github.com/lhlyu/justauth-go/result"
	"github.com/lhlyu/justauth-go/source"
	"github.com/lhlyu/justauth-go/utils"
)

type giteeRequest struct {
	BaseRequest
}

func newGiteeRequest(cfg config.AuthConfig, src source.AuthSource) AuthRequest {
	authRequest := &giteeRequest{}
	authRequest.Set(cfg, src)
	return authRequest
}

// Override 返回授权url
func (this *giteeRequest) Authorize() *result.UrlResult {
	return this.AuthorizeWithState("")
}

// Override 返回授权url + state
func (this *giteeRequest) AuthorizeWithState(state string) *result.UrlResult {
	url := utils.NewUrlBuilder(this.Source.Authorize()).
		AddParam("response_type", "code").
		AddParam("client_id", this.Config.ClientId).
		AddParam("redirect_uri", this.Config.RedirectUrl).
		AddParam("state", this.GetState(state)).
		Build()
	return result.Success.WithVal(url).ToUrlResult()
}

// Override 登录返回用户信息
func (this *giteeRequest) Login(callback *entity.Callback) *result.UserResult {
	url := utils.NewUrlBuilder(this.Source.AccessToken()).
		AddParam("grant_type", "authorization_code").
		AddParam("code", callback.Code).
		AddParam("client_id", this.Config.ClientId).
		AddParam("client_secret", this.Config.ClientSecret).
		AddParam("redirect_uri", this.Config.RedirectUrl).
		Build()
	rs := this.getAccessToken(url)
	if !rs.Ok() {
		return rs.ToUserResult()
	}
	token := rs.ToTokenResult().Val()
	url = utils.NewUrlBuilder(this.Source.UserInfo()).
		AddParam("access_token", token.AccessToken).
		Build()
	rs = this.getUserInfo(url)
	if !rs.Ok() {
		return rs.ToUserResult()
	}
	userResult := rs.ToUserResult()
	userResult.Val().Token = token
	return userResult
}

// Override 刷新token，返回新token
func (this *giteeRequest) Refresh(token *entity.Token) *result.TokenResult {
	url := utils.NewUrlBuilder(this.Source.Refresh()).
		AddParam("grant_type", "refresh_token").
		AddParam("refresh_token", token.RefreshToken).
		Build()
	return this.getAccessToken(url).ToTokenResult()
}

// ------------------------------------------------------------------

// get access token
func (this *giteeRequest) getAccessToken(url string) *result.Result {
	body, err := utils.Post(url)
	if err != nil {
		return result.Failure.WithErr(err)
	}
	m := utils.JsonToMSS(body)
	if _, ok := m["error"]; ok {
		return result.Failure.WithMsg(m["error_description"])
	}
	token := &entity.Token{
		AccessToken:  m["access_token"],
		RefreshToken: m["refresh_token"],
		ExpireIn:     m["expires_in"],
		Scope:        m["scope"],
		TokenType:    m["token_type"],
	}
	return result.Success.WithVal(token)
}

// get user info
func (this *giteeRequest) getUserInfo(url string) *result.Result {
	body, err := utils.Get(url)
	if err != nil {
		return result.Failure.WithErr(err)
	}
	m := utils.JsonToMSS(body)
	if _, ok := m["error"]; ok {
		return result.Failure.WithMsg(m["error_description"])
	}
	user := &entity.User{
		UUID:      m["id"],
		UserName:  m["login"],
		Avatar:    m["avatar_url"],
		Blog:      m["blog"],
		NickName:  m["name"],
		Company:   m["company"],
		Location:  m["address"],
		Email:     m["email"],
		Remark:    m["bio"],
		Url:       m["html_url"],
		CreatedAt: m["created_at"],
		UpdatedAt: m["updated_at"],
		Source:    this.Source.Name(),
		Gender:    utils.GetRealGender("").Desc,
	}
	return result.Success.WithVal(user)
}
