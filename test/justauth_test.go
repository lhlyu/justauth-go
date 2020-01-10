package test

import (
	"encoding/json"
	"github.com/lhlyu/justauth-go/config"
	"github.com/lhlyu/justauth-go/model"
	"github.com/lhlyu/justauth-go/request"
	"github.com/lhlyu/justauth-go/source"
	"testing"
)

const (
	CLIENT_ID     = "Iv1.094eec991d1d290dx"
	CLIENT_SECRET = "26074c03ea0167590039f3fb175078a14d864ce1x"
	REDIRECT_URL  = "http://localhost:8080/login"

	CODE  = "85d27e8e172ecfbbbc6d"
	STATE = "test"
)

// 生成授权URL
func TestAuthorize(t *testing.T) {
	authRequest, reqRs := request.NewAuthRequest(config.AuthConfig{
		ClientId:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		RedirectUrl:  REDIRECT_URL,
	}, source.GITHUB)
	if !reqRs.Ok() {
		t.Error(reqRs.Msg())
		return
	}
	rs := authRequest.AuthorizeWithState(STATE)
	if !rs.Ok() {
		t.Error(rs.Msg())
		return
	}
	t.Log(rs.Val())
}

// 登录获取用户信息
func TestLogin(t *testing.T) {
	authRequest, reqRs := request.NewAuthRequest(config.AuthConfig{
		ClientId:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		RedirectUrl:  REDIRECT_URL,
	}, source.GITHUB)
	if !reqRs.Ok() {
		t.Error(reqRs.Msg())
		return
	}
	rs := authRequest.Login(&model.Callback{
		Code:  CODE,
		State: STATE,
	})
	if !rs.Ok() {
		t.Error(rs.Msg())
		return
	}
	bts, _ := json.Marshal(rs.Val())
	t.Log(string(bts))
}
