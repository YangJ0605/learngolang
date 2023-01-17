package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func sayHello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "<h1>hello word</h1>")
// }

// func main() {
// 	http.HandleFunc("/hello", sayHello)
// 	err := http.ListenAndServe(":9090", nil)
// 	if err != nil {
// 		fmt.Println("http server failed, err: ", err)
// 		return
// 	}

// }

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world!",
	})
}

func main() {
	r := gin.Default() // 返回默认的路由引擎
	r.GET("/hello", sayHello)
	r.GET("/book", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"data": []string{"react", "vue", "gin"},
		})
	})
	r.Run()
}
