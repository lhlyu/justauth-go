package model

import "github.com/lhlyu/justauth-go/enums"

type AuthResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewAuthResponse(code int) *AuthResponse {
	msg := enums.ResponseStatusMap[code]
	return &AuthResponse{
		Code: code,
		Msg:  msg,
	}
}

func (this *AuthResponse) Ok() bool {
	return this.Code == enums.SUCCESS
}

func (this *AuthResponse) WithData(data interface{}) *AuthResponse {
	this.Data = data
	return this
}

var (
	Success = NewAuthResponse(enums.SUCCESS)
	Failure = NewAuthResponse(enums.FAILURE)
)
