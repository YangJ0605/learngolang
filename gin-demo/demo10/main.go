package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name string `json:"name" binding:"required,max=6,min=2" uri:"name"`
	Age  int8   `json:"age" uri:"age"`
	Sex  string `json:"sex" uri:"sex"`
}

var requestUser = map[string]string{
	"Name.min":      "用户名最少为两位",
	"Name.required": "用户名是必填的",
	"Name.max":      "用户名最长为两位",
}

type Validator interface {
	GetMessage() map[string]any
}

func GetErrorMessage(err error) string {

	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, err := range errs {
			str := err.Field() + "." + err.Tag()
			if msg, ok := requestUser[str]; ok {
				return msg
			} else {
				return err.Error()
			}
		}
	}
	return "params error"
}

func main() {
	r := gin.Default()

	r.POST("/user", func(ctx *gin.Context) {
		var user User
		err := ctx.ShouldBind(&user)
		fmt.Printf("err: %v\n", err)
		if err != nil {
			ctx.JSON(200, gin.H{
				"msg": GetErrorMessage(err),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"msg": user,
		})
	})

	r.POST("/user/:name/:age/:sex", func(ctx *gin.Context) {
		var user User
		err := ctx.ShouldBindUri(&user)
		fmt.Printf("err: %v\n", err)
		if err != nil {
			ctx.JSON(200, gin.H{
				"msg": GetErrorMessage(err),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"msg": user,
		})
	})

	r.Run(":9001")
}
