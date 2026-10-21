package service

import (
	"net/http"

	"github.com/Jie-cv/go-hmdp/model"
	"github.com/gin-gonic/gin"
)

func GetShopList(ctx *gin.Context) {
	var shopList []model.TbShop
	err := model.DB.Model(new(model.TbShop)).Find(&shopList).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": shopList})
}
