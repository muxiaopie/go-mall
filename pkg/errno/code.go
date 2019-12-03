package errno

import "net/http"

const (
	SUCCESS_CODE    = 0
	SERVER_ERROR    = 999
	PARAMETER_ERROR = 1000
	NOT_FOUND       = 1001
	OTHER_ERROR     = 1002
	JWT_ERROR       = 1004
	UNKNOWN_ERROR   = 1005
)

var (
	Success     = NewError(http.StatusOK, SUCCESS_CODE, "success")
	ServerError = NewError(http.StatusInternalServerError, SERVER_ERROR, "系统异常，请稍后重试!")
	NotFound    = NewError(http.StatusNotFound, NOT_FOUND, http.StatusText(http.StatusNotFound))
)

// 其他错误
func OtherError(message string) *Error {
	return NewError(http.StatusForbidden, OTHER_ERROR, message)
}

// 参数错误
func ParameterError(message string) *Error {
	return NewError(http.StatusBadRequest, PARAMETER_ERROR, message)
}

// token
func JwtError(message string) *Error {
	return NewError(http.StatusUnauthorized, JWT_ERROR, message)
}
