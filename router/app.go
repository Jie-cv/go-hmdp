package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AppRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "黑马点评加载中....")
	})
	InitUserRouter(r)
	InitShopRouter(r)
	return r
}
