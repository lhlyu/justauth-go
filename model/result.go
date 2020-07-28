package model

type baseResult struct {
	err  error
	code int
	msg  string
}

func (this *baseResult) Ok() bool {
	if this.err != nil {
		return false
	}
	return true
}

func (this *baseResult) Err() error {
	return this.err
}

func (this *baseResult) Code() int {
	return this.code
}

func (this *baseResult) Msg() string {
	return this.msg
}

// -------------------------------------------

type Result struct {
	baseResult

	val interface{}
}

func NewResult(code int, msg string) *Result {
	return &Result{
		baseResult: baseResult{
			code: code,
			msg:  msg,
		},
	}
}

func (this *Result) copy() *Result {
	return NewResult(this.code, this.msg)
}

func (this *Result) Val() interface{} {
	return this.val
}

func (this *Result) WithVal(val interface{}) *Result {
	r := this.copy()
	r.val = val
	return r
}

func (this *Result) WithErr(err error) *Result {
	r := this.copy()
	r.err = err
	r.WithMsg(err.Error())
	return r
}

func (this *Result) WithMsg(msg string) *Result {
	r := this.copy()
	if r.msg == "" {
		r.msg = msg
	} else {
		r.msg += ":" + msg
	}
	return r
}

func (this *Result) ToUrlResult() *UrlResult {
	val, _ := this.val.(string)
	return &UrlResult{
		baseResult: this.baseResult,
		val:        val,
	}
}

func (this *Result) ToUserResult() *UserResult {
	val, _ := this.val.(*User)
	return &UserResult{
		baseResult: this.baseResult,
		val:        val,
	}
}

func (this *Result) ToTokenResult() *TokenResult {
	val, _ := this.val.(*Token)
	return &TokenResult{
		baseResult: this.baseResult,
		val:        val,
	}
}

func (this *Result) ToStatusResult() *StatusResult {
	val, _ := this.val.(bool)
	return &StatusResult{
		baseResult: this.baseResult,
		val:        val,
	}
}

// -----------------------------------------

type UrlResult struct {
	baseResult

	val string
}

func (this *UrlResult) Val() string {
	return this.val
}

// ------------------------------------------

type UserResult struct {
	baseResult

	val *User
}

func (this *UserResult) Val() *User {
	return this.val
}

// ------------------------------------------

type TokenResult struct {
	baseResult

	val *Token
}

func (this *TokenResult) Val() *Token {
	return this.val
}

// ------------------------------------------

type StatusResult struct {
	baseResult

	val bool
}

func (this *StatusResult) Val() bool {
	return this.val
}
