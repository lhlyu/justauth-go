package model

type AuthToken struct {
	GoogleAttr
	MiAttr
	WechatAttr
	Twitter

	AccessToken  string
	ExpireIn     int
	RefreshToken string
	Uid          string
	OpenId       string
	AccessCode   string
	UnionId      string
}

type GoogleAttr struct {
	Scope     string
	TokenType string
	IdToken   string
}

type MiAttr struct {
	MacAlgorithm string
	MacKey       string
}

type WechatAttr struct {
	Code string
}

type Twitter struct {
	OauthToken             string
	OauthTokenSecret       string
	UserId                 string
	ScreenName             string
	OauthCallbackConfirmed bool
}
