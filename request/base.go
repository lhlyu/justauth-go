package request

import (
	"github.com/lhlyu/justauth-go/config"
	"github.com/lhlyu/justauth-go/errcode"
	"github.com/lhlyu/justauth-go/model"
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

func (*BaseRequest) Authorize() (string, *errcode.ErrCode) {
	return "", errcode.NotImplemented
}

func (*BaseRequest) AuthorizeWithState(state string) (string, *errcode.ErrCode) {
	return "", errcode.NotImplemented
}

func (*BaseRequest) Login(callback *model.Callback) (*model.AuthResponse, *errcode.ErrCode) {
	return nil, errcode.NotImplemented
}

func (*BaseRequest) Revoke(token *model.AuthToken) (*model.AuthResponse, *errcode.ErrCode) {
	return nil, errcode.NotImplemented
}

func (*BaseRequest) Refresh(token *model.AuthToken) (*model.AuthResponse, *errcode.ErrCode) {
	return nil, errcode.NotImplemented
}
