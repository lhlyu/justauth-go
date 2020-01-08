package test

import (
	"encoding/json"
	"fmt"
	"github.com/lhlyu/justauth-go/config"
	"github.com/lhlyu/justauth-go/model"
	"github.com/lhlyu/justauth-go/request"
	"testing"
)

const (
	CLIENT_ID     = "Iv1.094eec991d1d290dx"
	CLIENT_SECRET = "26074c03ea0167590039f3fb175078a14d864ce1x"
	REDIRECT_URL  = "http://localhost:8080/login"

	CODE  = "6f0a63b8a8cd963772c4"
	STATE = "test"
)

func TestLogin(t *testing.T) {
	authRequest := request.NewGithubRequest(config.AuthConfig{
		ClientId:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		RedirectUrl:  REDIRECT_URL,
	})
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
	authRequest := request.NewGithubRequest(config.AuthConfig{
		ClientId:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		RedirectUrl:  REDIRECT_URL,
	})
	// 生成授权页面
	url, err := authRequest.AuthorizeWithState(STATE)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(url)
}
