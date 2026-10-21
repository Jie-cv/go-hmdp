package handler

import (
	"github.com/Jie-cv/go-hmdp/service"
	"github.com/gin-gonic/gin"
)

func GetShopHandlerShopList(ctx *gin.Context) {
	service.GetShopList(ctx)
}
