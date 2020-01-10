package request

import (
	"github.com/lhlyu/justauth-go/config"
	"github.com/lhlyu/justauth-go/model"
	"github.com/lhlyu/justauth-go/result"
	"github.com/lhlyu/justauth-go/source"
	"github.com/lhlyu/justauth-go/utils"
)

type AuthRequest interface {
	// 返回授权url，可自行跳转页面
	// 不建议使用该方式获取授权地址，不带{@code state}的授权地址，容易受到csrf攻击。
	Authorize() *result.UrlResult
	// 返回带{@code state}参数的授权url，授权回调时会带上这个{@code state}
	AuthorizeWithState(state string) *result.UrlResult
	// 第三方登录
	Login(callback *model.Callback) *result.UserResult
	// 撤销授权
	Revoke(token *model.AuthToken) *result.StatusResult
	// 刷新access token （续期）
	Refresh(token *model.AuthToken) *result.TokenResult
}

func NewAuthRequest(cfg config.AuthConfig, src source.AuthSource) (AuthRequest, *result.Result) {
	if rs := utils.CheckAuth(cfg, src); !rs.Ok() {
		return nil, rs
	}
	switch src {
	case source.GITHUB:
		return newGiteeRequest(cfg, src), nil
	case source.GITEE:
		return newGiteeRequest(cfg, src), nil
	}
	return nil, result.ParameterIncomplete
}
