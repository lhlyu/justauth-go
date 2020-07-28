package model

type AuthConfig struct {
	ClientId     string
	ClientSecret string
	RedirectUrl  string

	// alipay
	AlipayPublicKey string

	// qq
	UnionId bool

	// stackover
	StackOverflowKey string

	// wechat
	AgentId string
}
