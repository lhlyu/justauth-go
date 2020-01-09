package errcode

type ErrCode struct {
	Code int
	Msg  string
}

func NewErrCode(code int) *ErrCode {
	if v, ok := ResponseStatusMap[code]; ok {
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

var (
	Success             = NewErrCode(SUCCESS)
	Failure             = NewErrCode(FAILURE)
	NotImplemented      = NewErrCode(NOT_IMPLEMENTED)
	ParameterIncomplete = NewErrCode(PARAMETER_INCOMPLETE)
)
