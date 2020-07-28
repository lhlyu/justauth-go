package model

type Callback struct {
	Code  string `json:"code"`
	State string `json:"state"`

	// alipay
	AuthCode string `json:"auth_code"`

	// huawei
	AuthorizationCode string `json:"authorization_code"`

	// twitter
	OauthToken    string `json:"oauthToken"`
	OauthVerifier string `json:"oauthVerifier"`
}
