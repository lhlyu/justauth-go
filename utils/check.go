package utils

import (
	"github.com/lhlyu/justauth-go/config"
	"github.com/lhlyu/justauth-go/result"
	"github.com/lhlyu/justauth-go/source"
)

func CheckAuth(cfg config.AuthConfig, src source.AuthSource) *result.Result {
	isSupported := true
	if cfg.ClientId == "" || cfg.ClientSecret == "" || cfg.RedirectUrl == "" {
		isSupported = false
	}
	if isSupported && src == source.ALIPAY {
		isSupported = cfg.AlipayPublicKey != ""
	}
	if isSupported && src == source.STACK_OVERFLOW {
		isSupported = cfg.StackOverflowKey != ""
	}
	if isSupported && src == source.WECHAT_ENTERPRISE {
		isSupported = cfg.AgentId != ""
	}
	if !isSupported {
		return result.ParameterIncomplete
	}
	return nil
}
