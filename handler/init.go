package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muxiaopie/go-mall/pkg/errno"
	"github.com/muxiaopie/go-mall/pkg/jwt"
)

func userId(c *gin.Context) (uint, error) {
	if claims, ok := c.Get("claims"); ok {
		if claims, ok := claims.(*jwt.Claims); ok {
			return claims.Id, nil
		}
	}
	return 0, errno.JwtError("你还没有登陆")
}
