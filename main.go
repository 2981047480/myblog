package main

import (
	"os"
	"vblog/app/token/api"
	"vblog/common/config"
	"vblog/ioc"

	_ "vblog/app"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1、读配置
	config_path := os.Getenv("CONFIG_PATH")
	if config_path == "" {
		config_path = "etc/db.yaml"
	}
	// if err := config.ReadDBConf(config_path); err != nil {
	// 	panic(err)
	// }
	config.ReadDBConf(config_path)

	// 2、初始化Ioc
	if err := ioc.ControllerImpl.Init(); err != nil {
		panic(err)
	}

	// token模块子路由
	tokenApiHandler := api.NewTokenApiHandler()

	// 设置对外暴露api
	r := gin.Default()
	root := r.Group("vblog/api/v1")
	tokenApiHandler.Register(root.Group("token"))
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
