package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/muxiaopie/go-mall/pkg/errno"
)

/**
 * 错误处理
 */
func ErrorHandler(c *gin.Context)  {
	defer func() {
		err := recover()
		if err != nil {
			var Err *errno.Error
			if e,ok := err.(*errno.Error); ok {
				Err = e
			}else if e, ok := err.(error); ok {
				Err = errno.OtherError(e.Error())
			}else{
				Err = errno.ServerError
			}
			// 记录一个错误的日志
			c.JSON(Err.StatusCode,Err)
		}
	}()
	c.Next()
}
