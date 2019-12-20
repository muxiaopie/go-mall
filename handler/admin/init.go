package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/muxiaopie/go-mall/pkg/errno"
	"github.com/muxiaopie/go-mall/pkg/jwt"
	"net/http"
)

var (
	statusOk = http.StatusOK
)

type Page struct {
	Page,Limit int
}

func userId(c *gin.Context) (uint, error) {
	if claims, ok := c.Get("claims"); ok {
		if claims, ok := claims.(*jwt.Claims); ok {
			return claims.Id, nil
		}
	}
	return 0, errno.JwtError("你还没有登陆")
}
