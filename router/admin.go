package router

import (
	"github.com/gin-gonic/gin"
	"github.com/muxiaopie/go-mall/handler/admin"
	"github.com/muxiaopie/go-mall/handler/sd"
	"github.com/muxiaopie/go-mall/router/middleware"
	"github.com/muxiaopie/go-mall/service"
)

func adminR(router *gin.Engine)  {
	adminV1 := router.Group("/admin")

	// 用户
	user := admin.User{
		Srv:service.NewUserService(),
	}
	// 类目
	category := admin.Category{
		Srv:service.NewCategoryService(),
	}

	// 品牌
	brand := admin.Brand{
		Srv:service.NewBrandService(),
	}

	adminV1.POST("/login",user.Login)
	adminV1.POST("/users",user.Users)
	adminV1.POST("/register",user.Register)

	userV1 := adminV1.Group("/user",middleware.JWTAuth())
	{
		userV1.POST("",user.User)
	}

	//
	categoryV1 := adminV1.Group("/category",middleware.JWTAuth())
	{
		categoryV1.POST("update/:id",category.Update)
		categoryV1.POST("create",category.Create)
		categoryV1.POST("list",category.List)
		categoryV1.GET("delete/:id",category.Delete)
	}


	brandV1 := adminV1.Group("/brand",middleware.JWTAuth())
	{
		brandV1.POST("update/:id",brand.Update)
		brandV1.POST("create",brand.Create)
		brandV1.POST("list",brand.List)
		brandV1.GET("delete/:id",brand.Delete)
	}

	v1 := adminV1.Group("/sd")
	{
		v1.GET("/health", sd.HealthCheck)
		v1.GET("/disk", sd.DiskCheck)
		v1.GET("/cpu", sd.CPUCheck)
		v1.GET("/ram", sd.RAMCheck)
	}

}
