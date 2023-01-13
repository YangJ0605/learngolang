package main

import "fmt"

// 类型与接口的关系
/*
	个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。
	例如狗不仅可以叫，还可以动。
	我们完全可以分别定义Sayer接口和Mover接口
*/

// 一种类型实现多个接口
type Sayer interface {
	Say()
}

type Mover interface {
	Move()
}

type Dog struct {
	Name string
}

// 实现Sayer接口
func (d Dog) Say() {
	fmt.Printf("%v 会汪汪汪~\n", d.Name)
}

// 实现Mover接口
func (d Dog) Move() {
	fmt.Printf("%v is Moving\n", d.Name)
}

// 多种类型实现同一接口
type Car struct {
	Brand string
}

func (c Car) Move() {
	fmt.Printf("%v is 速度 70迈\n", c.Brand)
}

func Test() {
	var d = Dog{
		Name: "小邹",
	}

	var s Sayer = d
	var m Mover = d

	s.Say()
	m.Move()

	var obj Mover

	obj = Dog{
		Name: "邹",
	}
	obj.Move()

	obj = Car{
		Brand: "宝马",
	}
	obj.Move()
}

func main() {
	Test()
}
