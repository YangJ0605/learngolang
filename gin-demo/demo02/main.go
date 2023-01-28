package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("http listen failed, err:", err)
		return
	}

}

func sayHello(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("./hello.tpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}

	// 利用给定数据渲染模板，并将结果写入w
	tpl.Execute(w, User{
		Name: "cc",
		Age:  18,
	})

}
