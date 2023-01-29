package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Any("/any", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"method": ctx.Request.Method,
		})
	})

	userRouter := r.Group("/user")

	{
		userRouter.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": ctx.Request.Method + ctx.Request.RequestURI,
			})
		})

		userRouter.GET("/login", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": ctx.Request.Method + ctx.Request.RequestURI,
			})
		})
	}

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": ctx.Request.Method + ctx.Request.RequestURI + " is not found",
		})
	})
	r.Run(":9000")
}
