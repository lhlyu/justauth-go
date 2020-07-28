package justauth

import (
	"github.com/lhlyu/justauth-go/api"
	"github.com/lhlyu/justauth-go/model"
	"strings"
)

// 接口定义
type IAuth interface {
	// 返回授权URL
	Authorize() *model.UrlResult
	// 自定义state,返回授权URL
	AuthorizeWithState(state string) *model.UrlResult
	// 登录并返回用户信息
	Login(callback model.Callback) *model.UserResult
	// 撤销授权
	Revoke(token model.Token) *model.StatusResult
	// 刷新access token （续期）
	Refresh(token model.Token) *model.TokenResult
}

func NewAuth(source string, config model.AuthConfig) IAuth {
	source = strings.ToLower(source)
	switch source {
	case "github":
		return api.NewGithub(config)
	}
	return &api.BaseAuth{}
}
