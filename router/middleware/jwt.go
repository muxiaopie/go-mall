package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/muxiaopie/go-mall/pkg/errno"
	"github.com/muxiaopie/go-mall/pkg/jwt"
	"strings"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.DefaultQuery("token", "")
		if token == "" {
			//token = c.Request.Header.Get("Authorization")
			token = c.GetHeader("Authorization")
			if s := strings.Split(token, " "); len(s) == 2 {
				token = s[1]
			}
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			var httpErr *errno.Error
			if h, ok := (err).(*errno.Error); ok {
				httpErr = h
			}else if e,ok := err.(error); ok {
				httpErr = errno.JwtError(e.Error())
			}else {
				httpErr = errno.ServerError
			}
			c.JSON(httpErr.StatusCode,httpErr)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}
