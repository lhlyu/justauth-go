package request

import (
	"github.com/lhlyu/justauth-go/config"
	"github.com/lhlyu/justauth-go/model"
	"github.com/lhlyu/justauth-go/result"
	"github.com/lhlyu/justauth-go/source"
	"github.com/lhlyu/justauth-go/utils"
)

type BaseRequest struct {
	Config config.AuthConfig
	Source source.AuthSource
}

func (*BaseRequest) GetState(state string) string {
	if state == "" {
		return utils.GetUUID()
	}
	return state
}

// 生成授权URL
func (*BaseRequest) Authorize() *result.UrlResult {
	return result.NotImplemented.ToUrlResult()
}

// 生成授权URL
func (*BaseRequest) AuthorizeWithState(state string) *result.UrlResult {
	return result.NotImplemented.ToUrlResult()
}

// 登录返回用户信息
func (*BaseRequest) Login(callback *model.Callback) *result.UserResult {
	return result.NotImplemented.ToUserResult()
}

// 撤销授权,返回状态
func (*BaseRequest) Revoke(token *model.AuthToken) *result.StatusResult {
	return result.NotImplemented.ToStatusResult()
}

// 刷新token，返回新token
func (*BaseRequest) Refresh(token *model.AuthToken) *result.TokenResult {
	return result.NotImplemented.ToTokenResult()
}
