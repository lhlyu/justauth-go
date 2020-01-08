package request

import (
	"github.com/lhlyu/justauth-go/errcode"
	"github.com/lhlyu/justauth-go/model"
)

type AuthRequest interface {
	// 返回授权url，可自行跳转页面
	// 不建议使用该方式获取授权地址，不带{@code state}的授权地址，容易受到csrf攻击。
	Authorize() (string, *errcode.ErrCode)
	// 返回带{@code state}参数的授权url，授权回调时会带上这个{@code state}
	AuthorizeWithState(state string) (string, *errcode.ErrCode)
	// 第三方登录
	Login(authCallback *model.Callback) (*model.AuthResponse, *errcode.ErrCode)
	// 撤销授权
	Revoke(authToken *model.AuthToken) (*model.AuthResponse, *errcode.ErrCode)
	// 刷新access token （续期）
	Refresh(authToken *model.AuthToken) (*model.AuthResponse, *errcode.ErrCode)
}
