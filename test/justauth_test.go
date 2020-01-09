package test

import (
	"encoding/json"
	"fmt"
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

	CODE  = "5fb22e2057e11a93a4b6"
	STATE = "test"
)

func TestLogin(t *testing.T) {
	authRequest, err := request.NewAuthRequest(config.AuthConfig{
		ClientId:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		RedirectUrl:  REDIRECT_URL,
	}, source.GITHUB)
	if err != nil {
		t.Error(err)
		return
	}
	resp, err := authRequest.Login(&model.Callback{
		Code:  CODE,
		State: STATE,
	})
	if err != nil {
		t.Error(err)
		return
	}
	bts, _ := json.Marshal(resp)
	t.Log(string(bts))
}

func TestAuthorize(t *testing.T) {
	authRequest, err := request.NewAuthRequest(config.AuthConfig{
		ClientId:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		RedirectUrl:  REDIRECT_URL,
	}, source.GITHUB)
	if err != nil {
		t.Error(err)
		return
	}
	// 生成授权页面
	url, err := authRequest.AuthorizeWithState(STATE)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(url)
}
