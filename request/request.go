package request

import (
	"errors"
	"github.com/lhlyu/justauth-go/config"
	"github.com/lhlyu/justauth-go/entity"
	"github.com/lhlyu/justauth-go/result"
	"github.com/lhlyu/justauth-go/source"
	"github.com/lhlyu/justauth-go/utils"
)

type AuthRequest interface {
	// 返回授权URL
	Authorize() *result.UrlResult
	// 自定义state,返回授权URL
	AuthorizeWithState(state string) *result.UrlResult
	// 登录并返回用户信息
	Login(callback *entity.Callback) *result.UserResult
	// 撤销授权
	Revoke(token *entity.Token) *result.StatusResult
	// 刷新access token （续期）
	Refresh(token *entity.Token) *result.TokenResult
}

var param_error = errors.New("Parameter incomplete")

func NewAuthRequest(cfg config.AuthConfig, src source.AuthSource) (AuthRequest, error) {
	if !utils.IsSupport(cfg, src) {
		return nil, param_error
	}
	switch src {
	case source.GITHUB:
		return newGithubRequest(cfg, src), nil
	case source.GITEE:
		return newGiteeRequest(cfg, src), nil
	}
	return nil, param_error
}
