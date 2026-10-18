package main

import (
	"github.com/Jie-cv/go-hmdp/config"
	"github.com/Jie-cv/go-hmdp/router"
)

func main() {
	port, err := config.Load()
	if err != nil {

	}
	r := router.AppRouter()
	r.Run(":" + port)
}
