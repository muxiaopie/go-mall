package router

import (
	"github.com/gin-gonic/gin"
	"github.com/muxiaopie/go-mall/pkg/errno"
)

type HandlerFunc func(c *gin.Context) error

// 全局错误处理
func wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		err := handler(c)
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
			return
		}
	}
}
