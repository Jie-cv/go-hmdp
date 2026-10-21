package router

import (
	"github.com/Jie-cv/go-hmdp/handler"
	"github.com/gin-gonic/gin"
)

func InitShopRouter(r *gin.Engine) {
	shopRouter := r.Group("/shop")
	{
		shopRouter.GET("/list", handler.GetShopHandlerShopList)
	}
}
