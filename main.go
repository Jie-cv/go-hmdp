package main

import (
	"log"

	"github.com/Jie-cv/go-hmdp/config"
	"github.com/Jie-cv/go-hmdp/model"
	"github.com/Jie-cv/go-hmdp/router"
)

func main() {
	config, err := config.LoadEnv()
	if err != nil {
		log.Println("加载环境变量错误!", err)
	}
	model.InitDB(config.DSN)
	r := router.AppRouter()
	r.Run(":" + config.ServerPort)
}
