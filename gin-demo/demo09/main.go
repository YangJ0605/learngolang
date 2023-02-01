package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CatchPanic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic: %v\n", r)

				// debug.PrintStack()

				ctx.JSON(http.StatusOK, gin.H{
					"code": 1,
					"msg":  ErrorToString(r),
					"data": nil,
				})

				ctx.Abort()
			}

		}()

		ctx.Next()
	}
}

// 把一个错误转换成string
func ErrorToString(err interface{}) string {
	switch v := err.(type) {
	case error:
		return v.Error()
	default:
		return err.(string)
	}
}

func main() {
	r := gin.Default()

	r.Use(CatchPanic())

	r.GET("/test", func(ctx *gin.Context) {
		s := []int{1, 2, 4}
		s[100] = 500

	})
	r.Run(":9000")
}
