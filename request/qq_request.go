package request

import (
	"github.com/lhlyu/justauth-go/config"
	"github.com/lhlyu/justauth-go/enum"
	"github.com/lhlyu/justauth-go/model"
	"github.com/lhlyu/justauth-go/result"
	"github.com/lhlyu/justauth-go/source"
	"github.com/lhlyu/justauth-go/utils"
	"strings"
)

type qqRequest struct {
	BaseRequest
}

const qq_oauth = "https://graph.qq.com/oauth2.0/me"

func newQqRequest(cfg config.AuthConfig, src source.AuthSource) AuthRequest {
	var authRequest AuthRequest
	authRequest = &qqRequest{
		BaseRequest: BaseRequest{
			Source: src,
			Config: cfg,
		},
	}
	return authRequest
}

// Override 返回授权url
func (this *qqRequest) Authorize() *result.UrlResult {
	return this.AuthorizeWithState("")
}

// Override 返回授权url + state
func (this *qqRequest) AuthorizeWithState(state string) *result.UrlResult {
	url := utils.NewUrlBuilder(this.Source.Authorize()).
		AddParam("response_type", "code").
		AddParam("client_id", this.Config.ClientId).
		AddParam("redirect_uri", this.Config.RedirectUrl).
		AddParam("state", this.GetState(state)).Build()
	return result.Success.WithVal(url).ToUrlResult()
}

// Override 登录返回用户信息
func (this *qqRequest) Login(callback *model.Callback) *result.UserResult {
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

// Override 刷新token，返回新token
func (this *qqRequest) Refresh(token *model.AuthToken) *result.TokenResult {
	url := utils.NewUrlBuilder(this.Source.Refresh()).
		AddParam("client_id", this.Config.ClientId).
		AddParam("client_secret", this.Config.ClientSecret).
		AddParam("refresh_token", token).
		AddParam("grant_type", "refresh_token").
		AddParam("redirect_uri", this.Config.RedirectUrl).Build()
	body, err := utils.Get(url)
	if err != nil {
		return result.Failure.WithErr(err).ToTokenResult()
	}
	rs := this.getToken(body)
	return rs.ToTokenResult()
}

// ------------------------------------------------------------------

func (this *qqRequest) getAccessToken(callback *model.Callback) *result.Result {
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

func (this *qqRequest) getUserInfo(token *model.AuthToken) *result.Result {
	rs := this.getOpenId(token)
	if !rs.Ok() {
		return rs
	}
	url := utils.NewUrlBuilder(this.Source.UserInfo()).
		AddParam("access_token", token.AccessToken).Build()
	body, err := utils.Get(url)
	if err != nil {
		return result.Failure.WithErr(err)
	}
	m := utils.JsonToMSS(body)
	if _, ok := m["ret"]; ok {
		return result.Failure.WithMsg(m["msg"])
	}
	avatar := m["figureurl_qq_2"]
	if avatar == "" {
		avatar = m["figureurl_qq_1"]
	}
	user := &model.AuthUser{
		UUID:     rs.Val().(string),
		UserName: m["nickname"],
		Avatar:   avatar,
		Blog:     m["blog"],
		NickName: m["nickname"],
		Company:  m["company"],
		Location: m["location"],
		Remark:   m["bio"],
		Gender:   enum.GetRealGender(m["gender"]).Desc,
		Token:    token,
		Source:   this.Source.ToString(),
	}
	return result.Success.WithVal(user)
}

func (this *qqRequest) getToken(body string) *result.Result {
	m := utils.StrToMSS(body)
	_, hasAccessToken := m["access_token"]
	_, hasCode := m["code"]
	if !hasAccessToken || hasCode {
		return result.Failure.WithMsg(m["msg"])
	}
	token := &model.AuthToken{
		AccessToken:  m["access_token"],
		ExpireIn:     m["expires_in"],
		RefreshToken: m["refresh_token"],
	}
	return result.Success.WithVal(token)
}

func (this *qqRequest) getOpenId(token *model.AuthToken) *result.Result {
	isUnionId := 0
	if this.Config.UnionId {
		isUnionId = 1
	}
	url := utils.NewUrlBuilder(qq_oauth).
		AddParam("access_token", token.AccessToken).
		AddParam("unionid", isUnionId).Build()
	body, err := utils.Get(url)
	if err != nil {
		return result.Failure.WithErr(err)
	}
	body = strings.Replace(body, "callback(", "", -1)
	body = strings.Replace(body, ");", "", -1)
	openId := strings.TrimSpace(body)
	m := utils.JsonToMSS(openId)
	if _, has := m["error"]; has {
		return result.Failure.WithMsg(m["error"]).WithMsg(m["error_description"])
	}
	token.OpenId = m["openid"]
	if v, has := m["unionid"]; has {
		token.UnionId = v
	}
	val := token.UnionId
	if val == "" {
		val = token.OpenId
	}
	return result.Success.WithVal(val)
}
