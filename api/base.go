package api

import (
	"github.com/lhlyu/justauth-go/model"
)

type BaseAuth struct {
	// 授权的api
	authorizeApi string
	// 获取accessToken的api
	accessTokenApi string
	// 获取用户信息的api
	userInfoApi string
	// 取消授权的api
	revokeApi string
	// 刷新授权的api
	refreshApi string
	// 当前授权来源
	name string
	// 配置
	Config model.AuthConfig
}

// 生成授权URL
func (BaseAuth) Authorize() *model.UrlResult {
	return model.NotImplemented.ToUrlResult()
}

// 生成授权URL
func (BaseAuth) AuthorizeWithState(state string) *model.UrlResult {
	return model.NotImplemented.ToUrlResult()
}

// 登录返回用户信息
func (BaseAuth) Login(callback model.Callback) *model.UserResult {
	return model.NotImplemented.ToUserResult()
}

// 撤销授权,返回状态
func (BaseAuth) Revoke(token model.Token) *model.StatusResult {
	return model.NotImplemented.ToStatusResult()
}

// 刷新token，返回新token
func (BaseAuth) Refresh(token model.Token) *model.TokenResult {
	return model.NotImplemented.ToTokenResult()
}
