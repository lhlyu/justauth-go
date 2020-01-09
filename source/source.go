package source

type AuthSource interface {
	// 授权的api
	Authorize() string
	// 获取accessToken的api
	AccessToken() string
	// 获取用户信息的api
	UserInfo() string
	// 取消授权的api
	Revoke() string
	// 刷新授权的api
	Refresh() string
	// 获取当前授权来源
	ToString() string
}

// 32
var (
	GITHUB            = &GithubSource{}           // Github
	WEIBO             = &WeiboSource{}            // 新浪微博
	GITEE             = &GiteeSource{}            // gitee
	DINGTALK          = &DingtalkSource{}         // 钉钉
	BAIDU             = &BaiduSource{}            // 百度
	CSDN              = &CsdnSource{}             // csdn
	CODING            = &CodingSource{}           // Coding
	TENCENT_CLOUD     = &TencentCloudSource{}     // 腾讯云开发者平台（coding升级后就变成腾讯云开发者平台了）
	OSCHINA           = &OschinaSource{}          // oschina 开源中国
	ALIPAY            = &AlipaySource{}           // 支付宝
	QQ                = &QQSource{}               // QQ
	WECHAT_OPEN       = &WechatOpenSource{}       // 微信开放平台
	WECHAT_MP         = &WechatMpSource{}         // 微信公众平台
	TAOBAO            = &TaobaoSource{}           // 淘宝
	GOOGLE            = &GoogleSource{}           // Google
	FACEBOOK          = &FacebookSource{}         // Facebook
	DOUYIN            = &DouyinSource{}           // 抖音
	LINKEDIN          = &LinkedinSource{}         // 领英
	MICROSOFT         = &MicrosoftSource{}        // 微软
	MI                = &MiSource{}               // 小米
	TOUTIAO           = &ToutiaoSource{}          // 今日头条
	TEAMBITION        = &TeambitionSource{}       // Teambition
	RENREN            = &RenrenSource{}           // 人人网
	PINTEREST         = &PinterestSource{}        // Pinterest
	STACK_OVERFLOW    = &StackOverflowSource{}    // Stack Overflow
	HUAWEI            = &HuaweiSource{}           // 华为
	WECHAT_ENTERPRISE = &WechatEnterpriseSource{} // 企业微信
	KUJIALE           = &KujialeSource{}          // 酷家乐
	GITLAB            = &GitlabSource{}           // Gitlab
	MEITUAN           = &MeituanSource{}          // 美团
	ELEME             = &ElemeSource{}            // 饿了么
	TWITTER           = &TwitterSource{}          // Twitter
)

// Github
type GithubSource struct{}

func (*GithubSource) Authorize() string   { return "https://github.com/login/oauth/authorize" }
func (*GithubSource) AccessToken() string { return "https://github.com/login/oauth/access_token" }
func (*GithubSource) UserInfo() string    { return "https://api.github.com/user" }
func (*GithubSource) Revoke() string      { return "" }
func (*GithubSource) Refresh() string     { return "" }
func (*GithubSource) ToString() string    { return "Github" }

// 新浪微博
type WeiboSource struct{}

func (*WeiboSource) Authorize() string   { return "https://api.weibo.com/oauth2/authorize" }
func (*WeiboSource) AccessToken() string { return "https://api.weibo.com/oauth2/access_token" }
func (*WeiboSource) UserInfo() string    { return "https://api.weibo.com/2/users/show.json" }
func (*WeiboSource) Revoke() string      { return "https://api.weibo.com/oauth2/revokeoauth2" }
func (*WeiboSource) Refresh() string     { return "" }
func (*WeiboSource) ToString() string    { return "Weibo" }

// gitee
type GiteeSource struct{}

func (*GiteeSource) Authorize() string   { return "https://gitee.com/oauth/authorize" }
func (*GiteeSource) AccessToken() string { return "https://gitee.com/oauth/token" }
func (*GiteeSource) UserInfo() string    { return "https://gitee.com/api/v5/user" }
func (*GiteeSource) Revoke() string      { return "" }
func (*GiteeSource) Refresh() string     { return "" }
func (*GiteeSource) ToString() string    { return "Gitee" }

// 钉钉
type DingtalkSource struct{}

func (*DingtalkSource) Authorize() string   { return "https://oapi.dingtalk.com/connect/qrconnect" }
func (*DingtalkSource) AccessToken() string { return "w AuthException(AuthResponseStatus.UNSUPPORTED" }
func (*DingtalkSource) UserInfo() string    { return "https://oapi.dingtalk.com/sns/getuserinfo_bycode" }
func (*DingtalkSource) Revoke() string      { return "" }
func (*DingtalkSource) Refresh() string     { return "" }
func (*DingtalkSource) ToString() string    { return "Dingtalk" }

// 百度
type BaiduSource struct{}

func (*BaiduSource) Authorize() string   { return "https://openapi.baidu.com/oauth/2.0/authorize" }
func (*BaiduSource) AccessToken() string { return "https://openapi.baidu.com/oauth/2.0/token" }
func (*BaiduSource) UserInfo() string {
	return "https://openapi.baidu.com/rest/2.0/passport/users/getInfo"
}
func (*BaiduSource) Revoke() string {
	return "https://openapi.baidu.com/rest/2.0/passport/auth/revokeAuthorization"
}
func (*BaiduSource) Refresh() string  { return "https://openapi.baidu.com/oauth/2.0/token" }
func (*BaiduSource) ToString() string { return "Baidu" }

// csdn
type CsdnSource struct{}

func (*CsdnSource) Authorize() string   { return "https://api.csdn.net/oauth2/authorize" }
func (*CsdnSource) AccessToken() string { return "https://api.csdn.net/oauth2/access_token" }
func (*CsdnSource) UserInfo() string    { return "https://api.csdn.net/user/getinfo" }
func (*CsdnSource) Revoke() string      { return "" }
func (*CsdnSource) Refresh() string     { return "" }
func (*CsdnSource) ToString() string    { return "Csdn" }

// Coding
type CodingSource struct{}

func (*CodingSource) Authorize() string   { return "https://coding.net/oauth_authorize.html" }
func (*CodingSource) AccessToken() string { return "https://coding.net/api/oauth/access_token" }
func (*CodingSource) UserInfo() string    { return "https://coding.net/api/account/current_user" }
func (*CodingSource) Revoke() string      { return "" }
func (*CodingSource) Refresh() string     { return "" }
func (*CodingSource) ToString() string    { return "Coding" }

// 腾讯云开发者平台（coding升级后就变成腾讯云开发者平台了）
type TencentCloudSource struct{}

func (*TencentCloudSource) Authorize() string { return "https://dev.tencent.com/oauth_authorize.html" }
func (*TencentCloudSource) AccessToken() string {
	return "https://dev.tencent.com/api/oauth/access_token"
}
func (*TencentCloudSource) UserInfo() string {
	return "https://dev.tencent.com/api/account/current_user"
}
func (*TencentCloudSource) Revoke() string   { return "" }
func (*TencentCloudSource) Refresh() string  { return "" }
func (*TencentCloudSource) ToString() string { return "TencentCloud" }

// oschina 开源中国
type OschinaSource struct{}

func (*OschinaSource) Authorize() string   { return "https://www.oschina.net/action/oauth2/authorize" }
func (*OschinaSource) AccessToken() string { return "https://www.oschina.net/action/openapi/token" }
func (*OschinaSource) UserInfo() string    { return "https://www.oschina.net/action/openapi/user" }
func (*OschinaSource) Revoke() string      { return "" }
func (*OschinaSource) Refresh() string     { return "" }
func (*OschinaSource) ToString() string    { return "Oschina" }

// 支付宝
type AlipaySource struct{}

func (*AlipaySource) Authorize() string {
	return "https://openauth.alipay.com/oauth2/publicAppAuthorize.htm"
}
func (*AlipaySource) AccessToken() string { return "https://openapi.alipay.com/gateway.do" }
func (*AlipaySource) UserInfo() string    { return "https://openapi.alipay.com/gateway.do" }
func (*AlipaySource) Revoke() string      { return "" }
func (*AlipaySource) Refresh() string     { return "" }
func (*AlipaySource) ToString() string    { return "Alipay" }

// QQ
type QQSource struct{}

func (*QQSource) Authorize() string   { return "https://graph.qq.com/oauth2.0/authorize" }
func (*QQSource) AccessToken() string { return "https://graph.qq.com/oauth2.0/token" }
func (*QQSource) UserInfo() string    { return "https://graph.qq.com/user/get_user_info" }
func (*QQSource) Revoke() string      { return "" }
func (*QQSource) Refresh() string     { return "https://graph.qq.com/oauth2.0/token" }
func (*QQSource) ToString() string    { return "QQ" }

// 微信开放平台
type WechatOpenSource struct{}

func (*WechatOpenSource) Authorize() string { return "https://open.weixin.qq.com/connect/qrconnect" }
func (*WechatOpenSource) AccessToken() string {
	return "https://api.weixin.qq.com/sns/oauth2/access_token"
}
func (*WechatOpenSource) UserInfo() string { return "https://api.weixin.qq.com/sns/userinfo" }
func (*WechatOpenSource) Revoke() string   { return "" }
func (*WechatOpenSource) Refresh() string  { return "https://api.weixin.qq.com/sns/oauth2/refresh_token" }
func (*WechatOpenSource) ToString() string { return "WechatOpen" }

// 微信公众平台
type WechatMpSource struct{}

func (*WechatMpSource) Authorize() string {
	return "https://open.weixin.qq.com/connect/oauth2/authorize"
}
func (*WechatMpSource) AccessToken() string {
	return "https://api.weixin.qq.com/sns/oauth2/access_token"
}
func (*WechatMpSource) UserInfo() string { return "https://api.weixin.qq.com/sns/userinfo" }
func (*WechatMpSource) Revoke() string   { return "" }
func (*WechatMpSource) Refresh() string  { return "https://api.weixin.qq.com/sns/oauth2/refresh_token" }
func (*WechatMpSource) ToString() string { return "WechatMp" }

// 淘宝
type TaobaoSource struct{}

func (*TaobaoSource) Authorize() string   { return "https://oauth.taobao.com/authorize" }
func (*TaobaoSource) AccessToken() string { return "https://oauth.taobao.com/token" }
func (*TaobaoSource) UserInfo() string    { return "w AuthException(AuthResponseStatus.UNSUPPORTED" }
func (*TaobaoSource) Revoke() string      { return "" }
func (*TaobaoSource) Refresh() string     { return "" }
func (*TaobaoSource) ToString() string    { return "Taobao" }

// Google
type GoogleSource struct{}

func (*GoogleSource) Authorize() string   { return "https://accounts.google.com/o/oauth2/v2/auth" }
func (*GoogleSource) AccessToken() string { return "https://www.googleapis.com/oauth2/v4/token" }
func (*GoogleSource) UserInfo() string    { return "https://www.googleapis.com/oauth2/v3/userinfo" }
func (*GoogleSource) Revoke() string      { return "" }
func (*GoogleSource) Refresh() string     { return "" }
func (*GoogleSource) ToString() string    { return "Google" }

// Facebook
type FacebookSource struct{}

func (*FacebookSource) Authorize() string { return "https://www.facebook.com/v3.3/dialog/oauth" }
func (*FacebookSource) AccessToken() string {
	return "https://graph.facebook.com/v3.3/oauth/access_token"
}
func (*FacebookSource) UserInfo() string { return "https://graph.facebook.com/v3.3/me" }
func (*FacebookSource) Revoke() string   { return "" }
func (*FacebookSource) Refresh() string  { return "" }
func (*FacebookSource) ToString() string { return "Facebook" }

// 抖音
type DouyinSource struct{}

func (*DouyinSource) Authorize() string   { return "https://open.douyin.com/platform/oauth/connect" }
func (*DouyinSource) AccessToken() string { return "https://open.douyin.com/oauth/access_token/" }
func (*DouyinSource) UserInfo() string    { return "https://open.douyin.com/oauth/userinfo/" }
func (*DouyinSource) Revoke() string      { return "" }
func (*DouyinSource) Refresh() string     { return "https://open.douyin.com/oauth/refresh_token/" }
func (*DouyinSource) ToString() string    { return "Douyin" }

// 领英
type LinkedinSource struct{}

func (*LinkedinSource) Authorize() string   { return "https://www.linkedin.com/oauth/v2/authorization" }
func (*LinkedinSource) AccessToken() string { return "https://www.linkedin.com/oauth/v2/accessToken" }
func (*LinkedinSource) UserInfo() string    { return "https://api.linkedin.com/v2/me" }
func (*LinkedinSource) Revoke() string      { return "" }
func (*LinkedinSource) Refresh() string     { return "https://www.linkedin.com/oauth/v2/accessToken" }
func (*LinkedinSource) ToString() string    { return "Linkedin" }

// 微软
type MicrosoftSource struct{}

func (*MicrosoftSource) Authorize() string {
	return "https://login.microsoftonline.com/common/oauth2/v2.0/authorize"
}
func (*MicrosoftSource) AccessToken() string {
	return "https://login.microsoftonline.com/common/oauth2/v2.0/token"
}
func (*MicrosoftSource) UserInfo() string { return "https://graph.microsoft.com/v1.0/me" }
func (*MicrosoftSource) Revoke() string   { return "" }
func (*MicrosoftSource) Refresh() string {
	return "https://login.microsoftonline.com/common/oauth2/v2.0/token"
}
func (*MicrosoftSource) ToString() string { return "Microsoft" }

// 小米
type MiSource struct{}

func (*MiSource) Authorize() string   { return "https://account.xiaomi.com/oauth2/authorize" }
func (*MiSource) AccessToken() string { return "https://account.xiaomi.com/oauth2/token" }
func (*MiSource) UserInfo() string    { return "https://open.account.xiaomi.com/user/profile" }
func (*MiSource) Revoke() string      { return "" }
func (*MiSource) Refresh() string     { return "https://account.xiaomi.com/oauth2/token" }
func (*MiSource) ToString() string    { return "Mi" }

// 今日头条
type ToutiaoSource struct{}

func (*ToutiaoSource) Authorize() string   { return "https://open.snssdk.com/auth/authorize" }
func (*ToutiaoSource) AccessToken() string { return "https://open.snssdk.com/auth/token" }
func (*ToutiaoSource) UserInfo() string    { return "https://open.snssdk.com/data/user_profile" }
func (*ToutiaoSource) Revoke() string      { return "" }
func (*ToutiaoSource) Refresh() string     { return "" }
func (*ToutiaoSource) ToString() string    { return "Toutiao" }

// Teambition
type TeambitionSource struct{}

func (*TeambitionSource) Authorize() string { return "https://account.teambition.com/oauth2/authorize" }
func (*TeambitionSource) AccessToken() string {
	return "https://account.teambition.com/oauth2/access_token"
}
func (*TeambitionSource) UserInfo() string { return "https://api.teambition.com/users/me" }
func (*TeambitionSource) Revoke() string   { return "" }
func (*TeambitionSource) Refresh() string {
	return "https://account.teambition.com/oauth2/refresh_token"
}
func (*TeambitionSource) ToString() string { return "Teambition" }

// 人人网
type RenrenSource struct{}

func (*RenrenSource) Authorize() string   { return "https://graph.renren.com/oauth/authorize" }
func (*RenrenSource) AccessToken() string { return "https://graph.renren.com/oauth/token" }
func (*RenrenSource) UserInfo() string    { return "https://api.renren.com/v2/user/get" }
func (*RenrenSource) Revoke() string      { return "" }
func (*RenrenSource) Refresh() string     { return "https://graph.renren.com/oauth/token" }
func (*RenrenSource) ToString() string    { return "Renren" }

// Pinterest
type PinterestSource struct{}

func (*PinterestSource) Authorize() string   { return "https://api.pinterest.com/oauth" }
func (*PinterestSource) AccessToken() string { return "https://api.pinterest.com/v1/oauth/token" }
func (*PinterestSource) UserInfo() string    { return "https://api.pinterest.com/v1/me" }
func (*PinterestSource) Revoke() string      { return "" }
func (*PinterestSource) Refresh() string     { return "" }
func (*PinterestSource) ToString() string    { return "Pinterest" }

// Stack Overflow
type StackOverflowSource struct{}

func (*StackOverflowSource) Authorize() string { return "https://stackoverflow.com/oauth" }
func (*StackOverflowSource) AccessToken() string {
	return "https://stackoverflow.com/oauth/access_token/json"
}
func (*StackOverflowSource) UserInfo() string { return "https://api.stackexchange.com/2.2/me" }
func (*StackOverflowSource) Revoke() string   { return "" }
func (*StackOverflowSource) Refresh() string  { return "" }
func (*StackOverflowSource) ToString() string { return "StackOverflow" }

// 华为
type HuaweiSource struct{}

func (*HuaweiSource) Authorize() string {
	return "https://oauth-login.cloud.huawei.com/oauth2/v2/authorize"
}
func (*HuaweiSource) AccessToken() string {
	return "https://oauth-login.cloud.huawei.com/oauth2/v2/token"
}
func (*HuaweiSource) UserInfo() string { return "https://api.vmall.com/rest.php" }
func (*HuaweiSource) Revoke() string   { return "" }
func (*HuaweiSource) Refresh() string  { return "https://oauth-login.cloud.huawei.com/oauth2/v2/token" }
func (*HuaweiSource) ToString() string { return "Huawei" }

// 企业微信
type WechatEnterpriseSource struct{}

func (*WechatEnterpriseSource) Authorize() string {
	return "https://open.work.weixin.qq.com/wwopen/sso/qrConnect"
}
func (*WechatEnterpriseSource) AccessToken() string {
	return "https://qyapi.weixin.qq.com/cgi-bin/gettoken"
}
func (*WechatEnterpriseSource) UserInfo() string {
	return "https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo"
}
func (*WechatEnterpriseSource) Revoke() string   { return "" }
func (*WechatEnterpriseSource) Refresh() string  { return "" }
func (*WechatEnterpriseSource) ToString() string { return "WechatEnterprise" }

// 酷家乐
type KujialeSource struct{}

func (*KujialeSource) Authorize() string   { return "https://oauth.kujiale.com/oauth2/show" }
func (*KujialeSource) AccessToken() string { return "https://oauth.kujiale.com/oauth2/auth/token" }
func (*KujialeSource) UserInfo() string    { return "https://oauth.kujiale.com/oauth2/openapi/user" }
func (*KujialeSource) Revoke() string      { return "" }
func (*KujialeSource) Refresh() string     { return "https://oauth.kujiale.com/oauth2/auth/token/refresh" }
func (*KujialeSource) ToString() string    { return "Kujiale" }

// Gitlab
type GitlabSource struct{}

func (*GitlabSource) Authorize() string   { return "https://gitlab.com/oauth/authorize" }
func (*GitlabSource) AccessToken() string { return "https://gitlab.com/oauth/token" }
func (*GitlabSource) UserInfo() string    { return "https://gitlab.com/api/v4/user" }
func (*GitlabSource) Revoke() string      { return "" }
func (*GitlabSource) Refresh() string     { return "" }
func (*GitlabSource) ToString() string    { return "Gitlab" }

// 美团
type MeituanSource struct{}

func (*MeituanSource) Authorize() string { return "https://openapi.waimai.meituan.com/oauth/authorize" }
func (*MeituanSource) AccessToken() string {
	return "https://openapi.waimai.meituan.com/oauth/access_token"
}
func (*MeituanSource) UserInfo() string { return "https://openapi.waimai.meituan.com/oauth/userinfo" }
func (*MeituanSource) Revoke() string   { return "" }
func (*MeituanSource) Refresh() string {
	return "https://openapi.waimai.meituan.com/oauth/refresh_token"
}
func (*MeituanSource) ToString() string { return "Meituan" }

// 饿了么
type ElemeSource struct{}

func (*ElemeSource) Authorize() string   { return "https://open-api.shop.ele.me/authorize" }
func (*ElemeSource) AccessToken() string { return "https://open-api.shop.ele.me/token" }
func (*ElemeSource) UserInfo() string    { return "https://open-api.shop.ele.me/api/v1/" }
func (*ElemeSource) Revoke() string      { return "" }
func (*ElemeSource) Refresh() string     { return "https://open-api.shop.ele.me/token" }
func (*ElemeSource) ToString() string    { return "Eleme" }

// Twitter
type TwitterSource struct{}

func (*TwitterSource) Authorize() string   { return "https://api.twitter.com/oauth/authenticate" }
func (*TwitterSource) AccessToken() string { return "https://api.twitter.com/oauth/access_token" }
func (*TwitterSource) UserInfo() string    { return "https://api.twitter.com/1.1/users/show.json" }
func (*TwitterSource) Revoke() string      { return "" }
func (*TwitterSource) Refresh() string     { return "" }
func (*TwitterSource) ToString() string    { return "Twitter" }
