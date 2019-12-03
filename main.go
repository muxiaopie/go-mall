package main

import (
	"github.com/muxiaopie/go-mall/bootstrap"
	"github.com/muxiaopie/go-mall/config"
	"github.com/muxiaopie/go-mall/router"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {

	// init config
	if err := config.Init(); err != nil {
		panic(err)
	}

	// 启动服务
	bootstrapper := bootstrap.Bootstrap
	// 加载gin
	bootstrapper.Use(bootstrap.LoadGin())

	// 加载数据库连接
	bootstrapper.Use(bootstrap.LoadDatabase())

	// 加载路由配置
	router.Init()

	// 启动服务
	bootstrapper.Run(viper.GetString("addr"))

}
