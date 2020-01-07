package request

import "github.com/lhlyu/justauth-go/exception"

type AuthRequest interface {
	Authorize() (string, *exception.AuthException)
}
