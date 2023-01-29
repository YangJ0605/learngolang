package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"pwd" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/user", func(ctx *gin.Context) {
		var u UserInfo
		err := ctx.ShouldBind(&u)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("u: %#v\n", u)
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Ok",
			})
		}
	})

	r.POST("/form", func(ctx *gin.Context) {
		var u UserInfo
		err := ctx.ShouldBind(&u)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("u: %#v\n", u)
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Ok",
			})
		}
	})

	r.POST("/json", func(ctx *gin.Context) {
		var u UserInfo
		err := ctx.ShouldBind(&u)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("u: %#v\n", u)
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Ok",
			})
		}
	})

	r.Run(":9000")
}
