package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lhlyu/justauth-go/model"
	"github.com/lhlyu/justauth-go/utils"
	"net/url"
)

// github
// doc: https://developer.github.com/apps/building-oauth-apps/authorizing-oauth-apps/#web-application-flow
//      https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/
//      https://docs.github.com/en/developers/apps/authorizing-oauth-apps
type Github struct {
	BaseAuth
}

func NewGithub(config model.AuthConfig) *Github {
	return &Github{
		BaseAuth: BaseAuth{
			authorizeApi:   "https://github.com/login/oauth/authorize",
			accessTokenApi: "https://github.com/login/oauth/access_token",
			userInfoApi:    "https://api.github.com/user",
			revokeApi:      "",
			refreshApi:     "",
			name:           "Github",
			Config:         config,
		},
	}
}

// 生成授权URL
func (self Github) Authorize() *model.UrlResult {
	return self.AuthorizeWithState(utils.GetRandomString(16))
}

// 生成授权URL
func (self Github) AuthorizeWithState(state string) *model.UrlResult {
	uv := &url.Values{}
	uv.Add("client_id", self.Config.ClientId)
	uv.Add("redirect_uri", self.Config.RedirectUrl)
	uv.Add("state", state)
	uv.Add("scope", "user")
	authorizeApi := fmt.Sprintf("%s?%s", self.authorizeApi, uv.Encode())
	return model.Success.WithVal(authorizeApi).ToUrlResult()
}

// 登录返回用户信息
func (self Github) Login(callback model.Callback) *model.UserResult {
	// 获取token
	data := map[string]string{
		"client_id":     self.Config.ClientId,
		"client_secret": self.Config.ClientSecret,
		"code":          callback.Code,
		"redirect_uri":  self.Config.RedirectUrl,
		"state":         callback.State,
	}
	headers := map[string]string{
		"Accept": "application/json",
	}
	body, err := utils.Post(self.accessTokenApi, data, headers, nil)
	if err != nil {
		return model.Failure.WithErr(err).ToUserResult()
	}
	// success : {"access_token":"eb9b93ae90f8d95c8ed3b997a7af593692c540da","token_type":"bearer","scope":""}
	// failure : {"error":"bad_verification_code","error_description":"The code passed is incorrect or expired.","error_uri":"https://developer.github.com/apps/managing-oauth-apps/troubleshooting-oauth-app-access-token-request-errors/#bad-verification-code"}
	resp := make(map[string]interface{})
	if err := json.Unmarshal([]byte(body), &resp); err != nil {
		return model.Failure.WithErr(err).ToUserResult()
	}
	if _, ok := resp["error"]; ok {
		return model.Failure.
			WithErr(errors.New(utils.ToString(resp, "error"))).
			WithMsg(utils.ToString(resp, "error_description")).
			ToUserResult()
	}
	// 获取用户信息
	headers["Authorization"] = fmt.Sprintf("token %s", resp["access_token"])
	body, err = utils.Get(self.userInfoApi, nil, headers)
	if err != nil {
		return model.Failure.WithErr(err).ToUserResult()
	}
	resp2 := make(map[string]interface{})
	if err := json.Unmarshal([]byte(body), &resp2); err != nil {
		return model.Failure.WithErr(err).ToUserResult()
	}
	if _, ok := resp2["id"]; !ok {
		return model.Failure.
			WithErr(errors.New(utils.ToString(resp2, "message"))).
			WithMsg(utils.ToString(resp2, "message")).
			ToUserResult()
	}
	user := self.getUser(resp, resp2)
	return model.Success.WithVal(user).ToUserResult()
}

// user
func (self Github) getUser(tokenMap, userMap map[string]interface{}) *model.User {
	byts, _ := json.Marshal(userMap)
	return &model.User{
		Token: model.Token{
			AccessToken: utils.ToString(tokenMap, "access_token"),
		},
		UUID:      utils.ToString(userMap, "id"),
		UserName:  utils.ToString(userMap, "login"),
		NickName:  utils.ToString(userMap, "name"),
		Avatar:    utils.ToString(userMap, "avatar_url"),
		Company:   utils.ToString(userMap, "company"),
		Blog:      utils.ToString(userMap, "blog"),
		Location:  utils.ToString(userMap, "location"),
		Email:     utils.ToString(userMap, "email"),
		Remark:    utils.ToString(userMap, "bio"),
		Url:       utils.ToString(userMap, "html_url"),
		Source:    self.name,
		CreatedAt: utils.ToString(userMap, "created_at"),
		UpdatedAt: utils.ToString(userMap, "updated_at"),
		Original:  string(byts),
	}
}
