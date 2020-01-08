package request

import (
	"github.com/lhlyu/justauth-go/config"
	"github.com/lhlyu/justauth-go/enums"
	"github.com/lhlyu/justauth-go/errcode"
	"github.com/lhlyu/justauth-go/model"
)

type BaseRequest struct {
	Config config.AuthConfig
	Source config.AuthSource
}

func (this *BaseRequest) SetConfig(cfg config.AuthConfig) *BaseRequest {
	this.Config = cfg
	return this
}

func (*BaseRequest) Authorize() (string, *errcode.ErrCode) {
	return "", errcode.NewErrCode(enums.NOT_IMPLEMENTED)
}

func (*BaseRequest) AuthorizeWithState(state string) (string, *errcode.ErrCode) {
	return "", errcode.NewErrCode(enums.NOT_IMPLEMENTED)
}

func (*BaseRequest) Login(callback *model.Callback) (*model.AuthResponse, *errcode.ErrCode) {
	return nil, errcode.NewErrCode(enums.NOT_IMPLEMENTED)
}

func (*BaseRequest) Revoke(authToken *model.AuthToken) (*model.AuthResponse, *errcode.ErrCode) {
	return nil, errcode.NewErrCode(enums.NOT_IMPLEMENTED)
}

func (*BaseRequest) Refresh(authToken *model.AuthToken) (*model.AuthResponse, *errcode.ErrCode) {
	return nil, errcode.NewErrCode(enums.NOT_IMPLEMENTED)
}
