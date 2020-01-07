package model

type Callback struct {
	AlipayCallback
	HuaweiCallback
	TwitterCallback

	Code  string `json:"code"`
	State string `json:"state"`
}

type AlipayCallback struct {
	AuthCode string `json:"auth_code"`
}

type HuaweiCallback struct {
	AuthorizationCode string `json:"authorization_code"`
}

type TwitterCallback struct {
	OauthToken    string `json:"oauthToken"`
	OauthVerifier string `json:"oauthVerifier"`
}
