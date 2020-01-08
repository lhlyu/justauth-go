package config

type AuthConfig struct {
	AlipayConfig
	QQConfig
	StackOverflowKeyConfig
	WechatConfig

	ClientId     string
	ClientSecret string
	RedirectUrl  string
}

type AlipayConfig struct {
	AlipayPublicKey string
}

type QQConfig struct {
	UnionId bool
}

type StackOverflowKeyConfig struct {
	StackOverflowKey string
}

type WechatConfig struct {
	AgentId string
}
