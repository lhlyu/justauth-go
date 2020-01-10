package result

import "github.com/lhlyu/justauth-go/model"

type baseResult struct {
	err  error
	code int
	msg  string
}

func (this *baseResult) Ok() bool {
	return this.code == SUCCESS
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

func (this *Result) Val() interface{} {
	return this.val
}

func (this *Result) WithVal(val interface{}) *Result {
	this.val = val
	return this
}

func (this *Result) WithErr(err error) *Result {
	this.err = err
	this.WithMsg(err.Error())
	return this
}

func (this *Result) WithMsg(msg string) *Result {
	if this.msg == "" {
		this.msg = msg
	} else {
		this.msg += ":" + msg
	}
	return this
}

func (this *Result) ToUrlResult() *UrlResult {
	val, _ := this.val.(string)
	return &UrlResult{
		baseResult: this.baseResult,
		val:        val,
	}
}

func (this *Result) ToUserResult() *UserResult {
	val, _ := this.val.(*model.AuthUser)
	return &UserResult{
		baseResult: this.baseResult,
		val:        val,
	}
}

func (this *Result) ToTokenResult() *TokenResult {
	val, _ := this.val.(*model.AuthToken)
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

	val *model.AuthUser
}

func (this *UserResult) Val() *model.AuthUser {
	return this.val
}

// ------------------------------------------

type TokenResult struct {
	baseResult

	val *model.AuthToken
}

func (this *TokenResult) Val() *model.AuthToken {
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
