package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func tempDemo(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Println("parse template failed, err", err)
		return
	}

	tmpl.Execute(w, []string{"vv"})
}

func blockDemo(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/base.tmpl", "./templates/index.tmpl")

	if err != nil {
		fmt.Println("parse template failed, err", err)
		return
	}

	type User struct {
		Name string
		Age  int
	}

	err = tmpl.ExecuteTemplate(w, "index.tmpl", User{
		Name: "cc",
		Age:  18,
	})
	if err != nil {
		fmt.Println("render template failed, err:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", tempDemo)
	http.HandleFunc("/block", blockDemo)
	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		fmt.Println("http server failed, err:", err)
		return
	}

}
