package model

type AuthToken struct {
	GoogleAttr
	MiAttr
	WechatAttr
	Twitter

	AccessToken  string `json:"accessToken"`
	ExpireIn     int    `json:"expireIn"`
	RefreshToken string `json:"refreshToken"`
	Uid          string `json:"uid"`
	OpenId       string `json:"openId"`
	AccessCode   string `json:"accessCode"`
	UnionId      string `json:"unionId"`
}

type GoogleAttr struct {
	Scope     string `json:"scope"`
	TokenType string `json:"tokenType"`
	IdToken   string `json:"idToken"`
}

type MiAttr struct {
	MacAlgorithm string `json:"macAlgorithm"`
	MacKey       string `json:"macKey"`
}

type WechatAttr struct {
	Code string `json:"code"`
}

type Twitter struct {
	OauthToken             string `json:"oauthToken"`
	OauthTokenSecret       string `json:"oauthTokenSecret"`
	UserId                 string `json:"userId"`
	ScreenName             string `json:"screenName"`
	OauthCallbackConfirmed bool   `json:"oauthCallbackConfirmed"`
}
