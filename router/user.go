package router

import "github.com/gin-gonic/gin"

func InitUserRouter(r *gin.Engine) {
	userRouter := r.Group("/user")
	{
		userRouter.POST("/login", func(ctx *gin.Context) {})
		userRouter.POST("/code", func(ctx *gin.Context) {})
		userRouter.POST("/me", func(ctx *gin.Context) {})
		userRouter.POST("/logout", func(ctx *gin.Context) {})
		userRouter.POST("/:id", func(ctx *gin.Context) {})
		userRouter.POST("/info/:id", func(ctx *gin.Context) {})
	}
}
