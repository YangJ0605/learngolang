package main

import (
	"fmt"

	stu "demo14/student"
)

func main() {
	fmt.Println("hello, module")
	stu.Make()
	// fmt.Printf("stu.a: %v\n", stu.a)
	fmt.Printf("stu.CommonConfig: %v\n", stu.CommonConfig)
	var student stu.Student

	student.Age = 1
	fmt.Printf("student: %v\n", student)
	// stu.Add()
}
