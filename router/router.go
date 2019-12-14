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

	// 用户
	user := handler.User{
		Srv:service.NewUserService(),
	}

	// 类目
	category := handler.Category{
		Sev:service.NewCategoryService(),
	}

	// brand
	brand := handler.Brand{
		Sev:service.NewBrandService(),
	}

	router.POST("/login",wrapper(user.Login))
	router.POST("/users",wrapper(user.Users))
	router.POST("/register",wrapper(user.Register))

	userV1 := router.Group("/user",middleware.JWTAuth())
	{
		userV1.POST("",wrapper(user.User))
	}

	//
	categoryV1 := router.Group("/category")
	{
		categoryV1.POST("update/:id",wrapper(category.Update))
		categoryV1.POST("create",wrapper(category.Create))
		categoryV1.POST("list",wrapper(category.List))
		categoryV1.GET("delete/:id",wrapper(category.Delete))
	}


	brandV1 := router.Group("/brand")
	{
		brandV1.POST("update/:id",wrapper(brand.Update))
		brandV1.POST("create",wrapper(brand.Create))
		brandV1.POST("list",wrapper(brand.List))
		brandV1.GET("delete/:id",wrapper(brand.Delete))
	}

	v1 := router.Group("/sd")
	{
		v1.GET("/health", sd.HealthCheck)
		v1.GET("/disk", sd.DiskCheck)
		v1.GET("/cpu", sd.CPUCheck)
		v1.GET("/ram", sd.RAMCheck)
	}
}



