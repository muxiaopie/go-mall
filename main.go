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

	bootstrapper.Use(bootstrap.LoadDatabase())

	router.Router()

	// 启动服务
	bootstrapper.Run(viper.GetString("addr"))

}