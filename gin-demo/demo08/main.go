package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// 定义一个中间件
func StatCost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		// 上下文设置值
		ctx.Set("name", "cc")
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "middleware",
		})
		// ctx.Abort() // 阻止后续的处理函数
		ctx.Next()

		cost := time.Since(start)
		log.Println(ctx.Request.Method+" ", ctx.Request.URL.RequestURI()+" ", cost)
	}
}

func main() {

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file. 写log到本地
	t := time.Now()
	f, _ := os.Create(fmt.Sprintf("gin_%d_%d_%d_%d_%d.log", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute()))
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()

	// r.Use(StatCost()) // 注册全局中间件

	r.GET("/test", StatCost(), func(ctx *gin.Context) {
		time.Sleep(2 * time.Second)
		name, _ := ctx.Get("name")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"name":    name,
		})
	})

	r.GET("/panic", func(ctx *gin.Context) {
		panic("error name")
	})

	r.Run(":9000")
}
