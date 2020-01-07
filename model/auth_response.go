package model

import "github.com/lhlyu/justauth-go/enums"

type AuthResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (this AuthResponse) Ok() bool {
	return this.Code == enums.SUCCESS
}
