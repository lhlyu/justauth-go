package errcode

import "github.com/lhlyu/justauth-go/enums"

type ErrCode struct {
	Code int
	Msg  string
}

func NewErrCode(code int) *ErrCode {
	if v, ok := enums.ResponseStatusMap[code]; ok {
		return &ErrCode{
			Code: code,
			Msg:  v,
		}
	}
	return &ErrCode{
		Code: code,
	}
}

func (this *ErrCode) SetMsg(msg string) *ErrCode {
	this.Msg = msg
	return this
}

func (this *ErrCode) WithMsg(msg string) *ErrCode {
	this.Msg += ":" + msg
	return this
}
