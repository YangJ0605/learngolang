package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 获取绝对路径
	ex, _ := os.Executable()
	fmt.Printf("filepath.Dir(ex): %v\n", filepath.Dir(ex))

	// 自定义模板函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	// 静态文件处理
	r.Static("/static", "./static")

	r.LoadHTMLGlob("templates/*")
	r.GET("/ccc", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello cc",
		})
	})
	// html 渲染
	r.GET("/user", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "user.html", gin.H{
			"title": "<a href='https://liwenzhou.com'>李文周的博客</a>",
		})
	})

	r.GET("/posts", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "xxx.html", gin.H{
			"title": "posts html",
		})
	})

	r.GET("/json", func(ctx *gin.Context) {
		ctx.XML(http.StatusOK, gin.H{
			"message": "hello",
		})
	})

	r.GET("/json2", func(ctx *gin.Context) {
		var user struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		user.Age = 18
		user.Name = "cc"
		ctx.JSON(http.StatusOK, user)
	})

	r.GET("/yaml", func(ctx *gin.Context) {
		ctx.YAML(http.StatusOK, gin.H{
			"message": "hello",
		})
	})

	// query string
	r.GET("/user/search", func(ctx *gin.Context) {
		// username := ctx.Query("username")
		// password := ctx.Query("password")

		username := ctx.DefaultQuery("username", "ff")
		password := ctx.DefaultQuery("password", "***")

		ctx.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"password": password,
		})
	})

	// 获取form参数
	r.POST("/user/login", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		ctx.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"password": password,
		})
	})

	// 获取json参数
	r.POST("/user/json", func(ctx *gin.Context) {
		// fmt.Printf("ctx.Request.Body: %v\n", ctx.Request.Body)

		// 解析body
		// body, _ := ioutil.ReadAll(ctx.Request.Body)
		// fmt.Printf("body: %v\n", string(body))

		body, _ := ctx.GetRawData()

		var m map[string]interface{}

		json.Unmarshal(body, &m)

		// 获取header
		// fmt.Printf("\"header\": %v\n", "header")
		// for k, v := range ctx.Request.Header {
		// 	fmt.Printf("k: %v v: %v\n", k, v)
		// }
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    m,
		})
	})

	// 动态参数
	r.GET("/user/:id/:name", func(ctx *gin.Context) {
		params := ctx.Params
		userId := ctx.Param("id")
		fmt.Printf("userId: %v\n", userId)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    params,
			"id":      userId,
		})
	})

	r.Run(":9000")
}
