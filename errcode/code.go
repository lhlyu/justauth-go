package errcode

const (
	SUCCESS = 2000
	FAILURE = 5000 + iota
	NOT_IMPLEMENTED
	PARAMETER_INCOMPLETE
	UNSUPPORTED
	NO_AUTH_SOURCE
	UNIDENTIFIED_PLATFORM
	ILLEGAL_REDIRECT_URI
	ILLEGAL_REQUEST
	ILLEGAL_CODE
	ILLEGAL_STATUS
	REQUIRED_REFRESH_TOKEN
)

var ResponseStatusMap = map[int]string{
	SUCCESS:                "Success",
	FAILURE:                "Failure",
	NOT_IMPLEMENTED:        "Not Implemented",
	PARAMETER_INCOMPLETE:   "Parameter incomplete",
	UNSUPPORTED:            "Unsupported operation",
	NO_AUTH_SOURCE:         "AuthDefaultSource cannot be null",
	UNIDENTIFIED_PLATFORM:  "Unidentified platform",
	ILLEGAL_REDIRECT_URI:   "Illegal redirect uri",
	ILLEGAL_REQUEST:        "Illegal request",
	ILLEGAL_CODE:           "Illegal code",
	ILLEGAL_STATUS:         "Illegal state",
	REQUIRED_REFRESH_TOKEN: "The refresh token is required; it must not be null",
}
