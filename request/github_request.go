package request

import (
	"github.com/lhlyu/justauth-go/config"
	"github.com/lhlyu/justauth-go/enum"
	"github.com/lhlyu/justauth-go/errcode"
	"github.com/lhlyu/justauth-go/model"
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

// Override 返回授权url，可自行跳转页面
func (this *githubRequest) Authorize() (string, *errcode.ErrCode) {
	return this.AuthorizeWithState("")
}

// Override
func (this *githubRequest) AuthorizeWithState(state string) (string, *errcode.ErrCode) {
	return utils.NewUrlBuilder(this.Source.Authorize()).
		AddParam("response_type", "code").
		AddParam("client_id", this.Config.ClientId).
		AddParam("redirect_uri", this.Config.RedirectUrl).
		AddParam("state", this.GetState(state)).Build(), nil
}

// Override 统一的登录入口
func (this *githubRequest) Login(callback *model.Callback) (*model.AuthUser, *errcode.ErrCode) {
	authToken, err := this.getAccessToken(callback)
	if err != nil {
		return nil, err
	}
	authUser, err := this.getUserInfo(authToken)
	if err != nil {
		return nil, err
	}
	return authUser, nil
}

func (this *githubRequest) getAccessToken(callback *model.Callback) (*model.AuthToken, *errcode.ErrCode) {
	url := utils.NewUrlBuilder(this.Source.AccessToken()).
		AddParam("code", callback.Code).
		AddParam("client_id", this.Config.ClientId).
		AddParam("client_secret", this.Config.ClientSecret).
		AddParam("grant_type", "authorization_code").
		AddParam("redirect_uri", this.Config.RedirectUrl).Build()
	body, err := utils.Post(url)
	if err != nil {
		return nil, err
	}
	m := utils.StrToMSS(body)
	if _, ok := m["error"]; ok {
		desc := m["error_description"]
		return nil, errcode.Failure.WithMsg(desc)
	}
	return &model.AuthToken{
		AccessToken: m["access_token"],
		Scope:       m["scope"],
		TokenType:   m["token_type"],
	}, nil
}

func (this *githubRequest) getUserInfo(authToken *model.AuthToken) (*model.AuthUser, *errcode.ErrCode) {
	url := utils.NewUrlBuilder(this.Source.UserInfo()).
		AddParam("access_token", authToken.AccessToken).Build()
	body, err := utils.Get(url)
	if err != nil {
		return nil, err
	}
	m := utils.JsonToMSS(body)
	if _, ok := m["error"]; ok {
		desc := m["error_description"]
		return nil, errcode.Failure.WithMsg(desc)
	}
	return &model.AuthUser{
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
	}, nil
}
