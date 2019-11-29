package router

import (
	"github.com/gin-gonic/gin"
	"github.com/muxiaopie/go-mall/bootstrap"
	"github.com/muxiaopie/go-mall/handler/sd"
	"github.com/muxiaopie/go-mall/router/middleware"
)

// 加载服务
func Router()  {
	var router *gin.Engine = bootstrap.Bootstrap.Router
	router.Use(middleware.NoCache)
	router.Use(middleware.Options)
	router.Use(middleware.Secure)
	router.Use(middleware.RequestId())

	v1 := router.Group("/sd")
	{
		v1.GET("/health", sd.HealthCheck)
		v1.GET("/disk", sd.DiskCheck)
		v1.GET("/cpu", sd.CPUCheck)
		v1.GET("/ram", sd.RAMCheck)
	}
}

