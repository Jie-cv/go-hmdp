package router

import "github.com/gin-gonic/gin"

func InitShopRouter(r *gin.Engine) {
	shopRouter := r.Group("/shop")
	{
		shopRouter.GET("/list", func(ctx *gin.Context) {})
		shopRouter.GET("/list", func(ctx *gin.Context) {})
		shopRouter.GET("/list", func(ctx *gin.Context) {})
		shopRouter.GET("/list", func(ctx *gin.Context) {})
		shopRouter.GET("/list", func(ctx *gin.Context) {})
	}
}
