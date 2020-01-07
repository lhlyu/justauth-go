package enums

const (
	EC0 = iota
	EC1
	EC2
	EC3
	EC4
	EC5
	EC6
	EC7
	EC8
	EC9
	EC10
	EC11
	EC12
	EC13
	EC21
	EC999
)

var ECMap = map[int]string{
	EC0:   "接口调用成功",
	EC1:   "API配置错误，未传入Client Key",
	EC2:   "API配置错误，Client Key错误，请检查是否和开放平台的ClientKey一致",
	EC3:   "没有授权信息",
	EC4:   "响应类型错误",
	EC5:   "授权类型错误",
	EC6:   "client_secret错误",
	EC7:   "authorize_code过期",
	EC8:   "指定url的scheme不是https",
	EC9:   "接口内部错误，请联系头条技术",
	EC10:  "access_token过期",
	EC11:  "缺少access_token",
	EC12:  "参数缺失",
	EC13:  "url错误",
	EC21:  "域名与登记域名不匹配",
	EC999: "未知错误，请联系头条技术",
}

type AuthToutiaoErrorCode struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
}

func GetErrorCode(errorCode int) *AuthToutiaoErrorCode {
	if v, ok := ECMap[errorCode]; ok {
		return &AuthToutiaoErrorCode{errorCode, v}
	}
	return &AuthToutiaoErrorCode{EC999, ECMap[EC999]}
}
