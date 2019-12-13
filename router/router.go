package router

import (
	"github.com/gin-gonic/gin"
	"github.com/muxiaopie/go-mall/bootstrap"
	"github.com/muxiaopie/go-mall/handler"
	"github.com/muxiaopie/go-mall/handler/sd"
	"github.com/muxiaopie/go-mall/pkg/errno"
	"github.com/muxiaopie/go-mall/router/middleware"
	"github.com/muxiaopie/go-mall/service"
)

// 加载服务
func Init()  {
	var router *gin.Engine = bootstrap.Bootstrap.Router
	router.Use(middleware.NoCache)
	router.Use(middleware.Options)
	router.Use(middleware.Secure)
	router.Use(middleware.RequestId())
	router.Use(middleware.Logger())

	router.NoMethod(errno.HandleNotFound)
	router.NoRoute(errno.HandleNotFound)

	user := handler.User{
		Srv:service.NewUserService(),
	}

	category := handler.Category{
		Sev:service.NewCategoryService(),
	}





	router.POST("/login",wrapper(user.Login))
	router.Any("/users",wrapper(user.Users))
	router.POST("/register",wrapper(user.Register))

	u := router.Group("/user",middleware.JWTAuth())
	{
		u.Any("",wrapper(user.User))
	}

	categoryRouter := router.Group("/category")
	{
		categoryRouter.POST("update",wrapper(category.Update))
		categoryRouter.POST("create",wrapper(category.Create))
		categoryRouter.POST("list",wrapper(category.List))
	}

	v1 := router.Group("/sd")
	{
		v1.GET("/health", sd.HealthCheck)
		v1.GET("/disk", sd.DiskCheck)
		v1.GET("/cpu", sd.CPUCheck)
		v1.GET("/ram", sd.RAMCheck)
	}
}



