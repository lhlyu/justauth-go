package model

import (
	"github.com/lhlyu/justauth-go/errcode"
)

type AuthResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewAuthResponse(code int) *AuthResponse {
	msg := errcode.ResponseStatusMap[code]
	return &AuthResponse{
		Code: code,
		Msg:  msg,
	}
}

func (this *AuthResponse) Ok() bool {
	return this.Code == errcode.SUCCESS
}

func (this *AuthResponse) WithData(data interface{}) *AuthResponse {
	this.Data = data
	return this
}

var (
	Success = NewAuthResponse(errcode.SUCCESS)
	Failure = NewAuthResponse(errcode.FAILURE)
)
