package exception

import "github.com/lhlyu/justauth-go/enums"

type AuthException struct {
	ErrorCode int    `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
}

func NewAuthException(code int) *AuthException {
	return &AuthException{code, enums.ResponseStatusMap[code]}
}
