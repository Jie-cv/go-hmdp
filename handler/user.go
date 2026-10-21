package handler

import (
	"github.com/Jie-cv/go-hmdp/service"
	"github.com/gin-gonic/gin"
)

func LoginHandler(ctx *gin.Context) {
	service.LoginService(ctx)
}
