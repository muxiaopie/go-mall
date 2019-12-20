package router

import (
	"github.com/gin-gonic/gin"
	"github.com/muxiaopie/go-mall/bootstrap"
	"github.com/muxiaopie/go-mall/pkg/errno"
	"github.com/muxiaopie/go-mall/router/middleware"
)


// 加载服务
func Init()  {
	var router *gin.Engine = bootstrap.Bootstrap.Router
	router.Use(middleware.NoCache)
	router.Use(middleware.Options)
	router.Use(middleware.Secure)
	router.Use(middleware.RequestId())
	router.Use(middleware.Logger())
	router.Use(middleware.ErrorHandler)

	router.NoMethod(errno.HandleNotFound)
	router.NoRoute(errno.HandleNotFound)

	adminR(router)


}


