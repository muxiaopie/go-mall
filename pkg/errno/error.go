package errno

import "github.com/gin-gonic/gin"

type Error struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
}

func (e *Error) Error() string {
	return e.Msg
}

func NewError(statusCode, Code int, msg string) *Error {
	return &Error{
		StatusCode: statusCode,
		Code:       Code,
		Msg:        msg,
	}
}

func HandleNotFound(c *gin.Context) {
	err := NotFound
	c.JSON(err.StatusCode,err)
	return
}