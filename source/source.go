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
	Name() string
}

// 32
var (
	GITHUB = &githubSource{} // Github
	//WEIBO             = &weiboSource{}            // 新浪微博
	GITEE = &giteeSource{} // gitee
	//DINGTALK          = &dingtalkSource{}         // 钉钉
	//BAIDU             = &baiduSource{}            // 百度
	//CSDN              = &csdnSource{}             // csdn
	//CODING            = &codingSource{}           // Coding
	//TENCENT_CLOUD     = &tencentCloudSource{}     // 腾讯云开发者平台（coding升级后就变成腾讯云开发者平台了）
	//OSCHINA           = &oschinaSource{}          // oschina 开源中国
	//ALIPAY            = &alipaySource{}           // 支付宝
	//QQ                = &qqSource{}               // QQ
	//WECHAT_OPEN       = &wechatOpenSource{}       // 微信开放平台
	//WECHAT_MP         = &wechatMpSource{}         // 微信公众平台
	//TAOBAO            = &taobaoSource{}           // 淘宝
	//GOOGLE            = &googleSource{}           // Google
	//FACEBOOK          = &facebookSource{}         // Facebook
	//DOUYIN            = &douyinSource{}           // 抖音
	//LINKEDIN          = &linkedinSource{}         // 领英
	//MICROSOFT         = &microsoftSource{}        // 微软
	//MI                = &miSource{}               // 小米
	//TOUTIAO           = &toutiaoSource{}          // 今日头条
	//TEAMBITION        = &teambitionSource{}       // Teambition
	//RENREN            = &renrenSource{}           // 人人网
	//PINTEREST         = &pinterestSource{}        // Pinterest
	//STACK_OVERFLOW    = &stackOverflowSource{}    // Stack Overflow
	//HUAWEI            = &huaweiSource{}           // 华为
	//WECHAT_ENTERPRISE = &wechatEnterpriseSource{} // 企业微信
	//KUJIALE           = &kujialeSource{}          // 酷家乐
	//GITLAB            = &gitlabSource{}           // Gitlab
	//MEITUAN           = &meituanSource{}          // 美团
	//ELEME             = &elemeSource{}            // 饿了么
	//TWITTER           = &twitterSource{}          // Twitter
)

// Github
type githubSource struct{}

func (*githubSource) Authorize() string   { return "https://github.com/login/oauth/authorize" }
func (*githubSource) AccessToken() string { return "https://github.com/login/oauth/access_token" }
func (*githubSource) UserInfo() string    { return "https://api.github.com/user" }
func (*githubSource) Revoke() string      { return "" }
func (*githubSource) Refresh() string     { return "" }
func (*githubSource) Name() string        { return "Github" }

// 新浪微博
type weiboSource struct{}

func (*weiboSource) Authorize() string   { return "https://api.weibo.com/oauth2/authorize" }
func (*weiboSource) AccessToken() string { return "https://api.weibo.com/oauth2/access_token" }
func (*weiboSource) UserInfo() string    { return "https://api.weibo.com/2/users/show.json" }
func (*weiboSource) Revoke() string      { return "https://api.weibo.com/oauth2/revokeoauth2" }
func (*weiboSource) Refresh() string     { return "" }
func (*weiboSource) Name() string        { return "Weibo" }

// gitee
type giteeSource struct{}

func (*giteeSource) Authorize() string   { return "https://gitee.com/oauth/authorize" }
func (*giteeSource) AccessToken() string { return "https://gitee.com/oauth/token" }
func (*giteeSource) UserInfo() string    { return "https://gitee.com/api/v5/user" }
func (*giteeSource) Revoke() string      { return "" }
func (*giteeSource) Refresh() string     { return "https://gitee.com/oauth/token" }
func (*giteeSource) Name() string        { return "Gitee" }

// 钉钉
type dingtalkSource struct{}

func (*dingtalkSource) Authorize() string   { return "https://oapi.dingtalk.com/connect/qrconnect" }
func (*dingtalkSource) AccessToken() string { return "w AuthException(AuthResponseStatus.UNSUPPORTED" }
func (*dingtalkSource) UserInfo() string    { return "https://oapi.dingtalk.com/sns/getuserinfo_bycode" }
func (*dingtalkSource) Revoke() string      { return "" }
func (*dingtalkSource) Refresh() string     { return "" }
func (*dingtalkSource) Name() string        { return "Dingtalk" }

// 百度
type baiduSource struct{}

func (*baiduSource) Authorize() string   { return "https://openapi.baidu.com/oauth/2.0/authorize" }
func (*baiduSource) AccessToken() string { return "https://openapi.baidu.com/oauth/2.0/token" }
func (*baiduSource) UserInfo() string {
	return "https://openapi.baidu.com/rest/2.0/passport/users/getInfo"
}
func (*baiduSource) Revoke() string {
	return "https://openapi.baidu.com/rest/2.0/passport/auth/revokeAuthorization"
}
func (*baiduSource) Refresh() string { return "https://openapi.baidu.com/oauth/2.0/token" }
func (*baiduSource) Name() string    { return "Baidu" }

// csdn
type csdnSource struct{}

func (*csdnSource) Authorize() string   { return "https://api.csdn.net/oauth2/authorize" }
func (*csdnSource) AccessToken() string { return "https://api.csdn.net/oauth2/access_token" }
func (*csdnSource) UserInfo() string    { return "https://api.csdn.net/user/getinfo" }
func (*csdnSource) Revoke() string      { return "" }
func (*csdnSource) Refresh() string     { return "" }
func (*csdnSource) Name() string        { return "Csdn" }

// Coding
type codingSource struct{}

func (*codingSource) Authorize() string   { return "https://coding.net/oauth_authorize.html" }
func (*codingSource) AccessToken() string { return "https://coding.net/api/oauth/access_token" }
func (*codingSource) UserInfo() string    { return "https://coding.net/api/account/current_user" }
func (*codingSource) Revoke() string      { return "" }
func (*codingSource) Refresh() string     { return "" }
func (*codingSource) Name() string        { return "Coding" }

// 腾讯云开发者平台（coding升级后就变成腾讯云开发者平台了）
type tencentCloudSource struct{}

func (*tencentCloudSource) Authorize() string { return "https://dev.tencent.com/oauth_authorize.html" }
func (*tencentCloudSource) AccessToken() string {
	return "https://dev.tencent.com/api/oauth/access_token"
}
func (*tencentCloudSource) UserInfo() string {
	return "https://dev.tencent.com/api/account/current_user"
}
func (*tencentCloudSource) Revoke() string  { return "" }
func (*tencentCloudSource) Refresh() string { return "" }
func (*tencentCloudSource) Name() string    { return "TencentCloud" }

// oschina 开源中国
type oschinaSource struct{}

func (*oschinaSource) Authorize() string   { return "https://www.oschina.net/action/oauth2/authorize" }
func (*oschinaSource) AccessToken() string { return "https://www.oschina.net/action/openapi/token" }
func (*oschinaSource) UserInfo() string    { return "https://www.oschina.net/action/openapi/user" }
func (*oschinaSource) Revoke() string      { return "" }
func (*oschinaSource) Refresh() string     { return "" }
func (*oschinaSource) Name() string        { return "Oschina" }

// 支付宝
type alipaySource struct{}

func (*alipaySource) Authorize() string {
	return "https://openauth.alipay.com/oauth2/publicAppAuthorize.htm"
}
func (*alipaySource) AccessToken() string { return "https://openapi.alipay.com/gateway.do" }
func (*alipaySource) UserInfo() string    { return "https://openapi.alipay.com/gateway.do" }
func (*alipaySource) Revoke() string      { return "" }
func (*alipaySource) Refresh() string     { return "" }
func (*alipaySource) Name() string        { return "Alipay" }

// QQ
type qqSource struct{}

func (*qqSource) Authorize() string   { return "https://graph.qq.com/oauth2.0/authorize" }
func (*qqSource) AccessToken() string { return "https://graph.qq.com/oauth2.0/token" }
func (*qqSource) UserInfo() string    { return "https://graph.qq.com/user/get_user_info" }
func (*qqSource) Revoke() string      { return "" }
func (*qqSource) Refresh() string     { return "https://graph.qq.com/oauth2.0/token" }
func (*qqSource) Name() string        { return "QQ" }

// 微信开放平台
type wechatOpenSource struct{}

func (*wechatOpenSource) Authorize() string { return "https://open.weixin.qq.com/connect/qrconnect" }
func (*wechatOpenSource) AccessToken() string {
	return "https://api.weixin.qq.com/sns/oauth2/access_token"
}
func (*wechatOpenSource) UserInfo() string { return "https://api.weixin.qq.com/sns/userinfo" }
func (*wechatOpenSource) Revoke() string   { return "" }
func (*wechatOpenSource) Refresh() string  { return "https://api.weixin.qq.com/sns/oauth2/refresh_token" }
func (*wechatOpenSource) Name() string     { return "WechatOpen" }

// 微信公众平台
type wechatMpSource struct{}

func (*wechatMpSource) Authorize() string {
	return "https://open.weixin.qq.com/connect/oauth2/authorize"
}
func (*wechatMpSource) AccessToken() string {
	return "https://api.weixin.qq.com/sns/oauth2/access_token"
}
func (*wechatMpSource) UserInfo() string { return "https://api.weixin.qq.com/sns/userinfo" }
func (*wechatMpSource) Revoke() string   { return "" }
func (*wechatMpSource) Refresh() string  { return "https://api.weixin.qq.com/sns/oauth2/refresh_token" }
func (*wechatMpSource) Name() string     { return "WechatMp" }

// 淘宝
type taobaoSource struct{}

func (*taobaoSource) Authorize() string   { return "https://oauth.taobao.com/authorize" }
func (*taobaoSource) AccessToken() string { return "https://oauth.taobao.com/token" }
func (*taobaoSource) UserInfo() string    { return "w AuthException(AuthResponseStatus.UNSUPPORTED" }
func (*taobaoSource) Revoke() string      { return "" }
func (*taobaoSource) Refresh() string     { return "" }
func (*taobaoSource) Name() string        { return "Taobao" }

// Google
type googleSource struct{}

func (*googleSource) Authorize() string   { return "https://accounts.google.com/o/oauth2/v2/auth" }
func (*googleSource) AccessToken() string { return "https://www.googleapis.com/oauth2/v4/token" }
func (*googleSource) UserInfo() string    { return "https://www.googleapis.com/oauth2/v3/userinfo" }
func (*googleSource) Revoke() string      { return "" }
func (*googleSource) Refresh() string     { return "" }
func (*googleSource) Name() string        { return "Google" }

// Facebook
type facebookSource struct{}

func (*facebookSource) Authorize() string { return "https://www.facebook.com/v3.3/dialog/oauth" }
func (*facebookSource) AccessToken() string {
	return "https://graph.facebook.com/v3.3/oauth/access_token"
}
func (*facebookSource) UserInfo() string { return "https://graph.facebook.com/v3.3/me" }
func (*facebookSource) Revoke() string   { return "" }
func (*facebookSource) Refresh() string  { return "" }
func (*facebookSource) Name() string     { return "Facebook" }

// 抖音
type douyinSource struct{}

func (*douyinSource) Authorize() string   { return "https://open.douyin.com/platform/oauth/connect" }
func (*douyinSource) AccessToken() string { return "https://open.douyin.com/oauth/access_token/" }
func (*douyinSource) UserInfo() string    { return "https://open.douyin.com/oauth/userinfo/" }
func (*douyinSource) Revoke() string      { return "" }
func (*douyinSource) Refresh() string     { return "https://open.douyin.com/oauth/refresh_token/" }
func (*douyinSource) Name() string        { return "Douyin" }

// 领英
type linkedinSource struct{}

func (*linkedinSource) Authorize() string   { return "https://www.linkedin.com/oauth/v2/authorization" }
func (*linkedinSource) AccessToken() string { return "https://www.linkedin.com/oauth/v2/accessToken" }
func (*linkedinSource) UserInfo() string    { return "https://api.linkedin.com/v2/me" }
func (*linkedinSource) Revoke() string      { return "" }
func (*linkedinSource) Refresh() string     { return "https://www.linkedin.com/oauth/v2/accessToken" }
func (*linkedinSource) Name() string        { return "Linkedin" }

// 微软
type microsoftSource struct{}

func (*microsoftSource) Authorize() string {
	return "https://login.microsoftonline.com/common/oauth2/v2.0/authorize"
}
func (*microsoftSource) AccessToken() string {
	return "https://login.microsoftonline.com/common/oauth2/v2.0/token"
}
func (*microsoftSource) UserInfo() string { return "https://graph.microsoft.com/v1.0/me" }
func (*microsoftSource) Revoke() string   { return "" }
func (*microsoftSource) Refresh() string {
	return "https://login.microsoftonline.com/common/oauth2/v2.0/token"
}
func (*microsoftSource) Name() string { return "Microsoft" }

// 小米
type miSource struct{}

func (*miSource) Authorize() string   { return "https://account.xiaomi.com/oauth2/authorize" }
func (*miSource) AccessToken() string { return "https://account.xiaomi.com/oauth2/token" }
func (*miSource) UserInfo() string    { return "https://open.account.xiaomi.com/user/profile" }
func (*miSource) Revoke() string      { return "" }
func (*miSource) Refresh() string     { return "https://account.xiaomi.com/oauth2/token" }
func (*miSource) Name() string        { return "Mi" }

// 今日头条
type toutiaoSource struct{}

func (*toutiaoSource) Authorize() string   { return "https://open.snssdk.com/auth/authorize" }
func (*toutiaoSource) AccessToken() string { return "https://open.snssdk.com/auth/token" }
func (*toutiaoSource) UserInfo() string    { return "https://open.snssdk.com/data/user_profile" }
func (*toutiaoSource) Revoke() string      { return "" }
func (*toutiaoSource) Refresh() string     { return "" }
func (*toutiaoSource) Name() string        { return "Toutiao" }

// Teambition
type teambitionSource struct{}

func (*teambitionSource) Authorize() string { return "https://account.teambition.com/oauth2/authorize" }
func (*teambitionSource) AccessToken() string {
	return "https://account.teambition.com/oauth2/access_token"
}
func (*teambitionSource) UserInfo() string { return "https://api.teambition.com/users/me" }
func (*teambitionSource) Revoke() string   { return "" }
func (*teambitionSource) Refresh() string {
	return "https://account.teambition.com/oauth2/refresh_token"
}
func (*teambitionSource) Name() string { return "Teambition" }

// 人人网
type renrenSource struct{}

func (*renrenSource) Authorize() string   { return "https://graph.renren.com/oauth/authorize" }
func (*renrenSource) AccessToken() string { return "https://graph.renren.com/oauth/token" }
func (*renrenSource) UserInfo() string    { return "https://api.renren.com/v2/user/get" }
func (*renrenSource) Revoke() string      { return "" }
func (*renrenSource) Refresh() string     { return "https://graph.renren.com/oauth/token" }
func (*renrenSource) Name() string        { return "Renren" }

// Pinterest
type pinterestSource struct{}

func (*pinterestSource) Authorize() string   { return "https://api.pinterest.com/oauth" }
func (*pinterestSource) AccessToken() string { return "https://api.pinterest.com/v1/oauth/token" }
func (*pinterestSource) UserInfo() string    { return "https://api.pinterest.com/v1/me" }
func (*pinterestSource) Revoke() string      { return "" }
func (*pinterestSource) Refresh() string     { return "" }
func (*pinterestSource) Name() string        { return "Pinterest" }

// Stack Overflow
type stackOverflowSource struct{}

func (*stackOverflowSource) Authorize() string { return "https://stackoverflow.com/oauth" }
func (*stackOverflowSource) AccessToken() string {
	return "https://stackoverflow.com/oauth/access_token/json"
}
func (*stackOverflowSource) UserInfo() string { return "https://api.stackexchange.com/2.2/me" }
func (*stackOverflowSource) Revoke() string   { return "" }
func (*stackOverflowSource) Refresh() string  { return "" }
func (*stackOverflowSource) Name() string     { return "StackOverflow" }

// 华为
type huaweiSource struct{}

func (*huaweiSource) Authorize() string {
	return "https://oauth-login.cloud.huawei.com/oauth2/v2/authorize"
}
func (*huaweiSource) AccessToken() string {
	return "https://oauth-login.cloud.huawei.com/oauth2/v2/token"
}
func (*huaweiSource) UserInfo() string { return "https://api.vmall.com/rest.php" }
func (*huaweiSource) Revoke() string   { return "" }
func (*huaweiSource) Refresh() string  { return "https://oauth-login.cloud.huawei.com/oauth2/v2/token" }
func (*huaweiSource) Name() string     { return "Huawei" }

// 企业微信
type wechatEnterpriseSource struct{}

func (*wechatEnterpriseSource) Authorize() string {
	return "https://open.work.weixin.qq.com/wwopen/sso/qrConnect"
}
func (*wechatEnterpriseSource) AccessToken() string {
	return "https://qyapi.weixin.qq.com/cgi-bin/gettoken"
}
func (*wechatEnterpriseSource) UserInfo() string {
	return "https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo"
}
func (*wechatEnterpriseSource) Revoke() string  { return "" }
func (*wechatEnterpriseSource) Refresh() string { return "" }
func (*wechatEnterpriseSource) Name() string    { return "WechatEnterprise" }

// 酷家乐
type kujialeSource struct{}

func (*kujialeSource) Authorize() string   { return "https://oauth.kujiale.com/oauth2/show" }
func (*kujialeSource) AccessToken() string { return "https://oauth.kujiale.com/oauth2/auth/token" }
func (*kujialeSource) UserInfo() string    { return "https://oauth.kujiale.com/oauth2/openapi/user" }
func (*kujialeSource) Revoke() string      { return "" }
func (*kujialeSource) Refresh() string     { return "https://oauth.kujiale.com/oauth2/auth/token/refresh" }
func (*kujialeSource) Name() string        { return "Kujiale" }

// Gitlab
type gitlabSource struct{}

func (*gitlabSource) Authorize() string   { return "https://gitlab.com/oauth/authorize" }
func (*gitlabSource) AccessToken() string { return "https://gitlab.com/oauth/token" }
func (*gitlabSource) UserInfo() string    { return "https://gitlab.com/api/v4/user" }
func (*gitlabSource) Revoke() string      { return "" }
func (*gitlabSource) Refresh() string     { return "" }
func (*gitlabSource) Name() string        { return "Gitlab" }

// 美团
type meituanSource struct{}

func (*meituanSource) Authorize() string { return "https://openapi.waimai.meituan.com/oauth/authorize" }
func (*meituanSource) AccessToken() string {
	return "https://openapi.waimai.meituan.com/oauth/access_token"
}
func (*meituanSource) UserInfo() string { return "https://openapi.waimai.meituan.com/oauth/userinfo" }
func (*meituanSource) Revoke() string   { return "" }
func (*meituanSource) Refresh() string {
	return "https://openapi.waimai.meituan.com/oauth/refresh_token"
}
func (*meituanSource) Name() string { return "Meituan" }

// 饿了么
type elemeSource struct{}

func (*elemeSource) Authorize() string   { return "https://open-api.shop.ele.me/authorize" }
func (*elemeSource) AccessToken() string { return "https://open-api.shop.ele.me/token" }
func (*elemeSource) UserInfo() string    { return "https://open-api.shop.ele.me/api/v1/" }
func (*elemeSource) Revoke() string      { return "" }
func (*elemeSource) Refresh() string     { return "https://open-api.shop.ele.me/token" }
func (*elemeSource) Name() string        { return "Eleme" }

// Twitter
type twitterSource struct{}

func (*twitterSource) Authorize() string   { return "https://api.twitter.com/oauth/authenticate" }
func (*twitterSource) AccessToken() string { return "https://api.twitter.com/oauth/access_token" }
func (*twitterSource) UserInfo() string    { return "https://api.twitter.com/1.1/users/show.json" }
func (*twitterSource) Revoke() string      { return "" }
func (*twitterSource) Refresh() string     { return "" }
func (*twitterSource) Name() string        { return "Twitter" }
