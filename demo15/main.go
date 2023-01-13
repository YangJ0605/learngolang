package main

import "fmt"

// interface 接口的基本内容

/*
	type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}
*/

// 接口不管你是什么类型 只管你实现了什么方法

type Writer interface {
	Write([]byte, []byte) error
}

// 接口Singer 里面有个Sing方法
type Singer interface {
	Sing()
}

// 结构体
type Bird struct{}

func (b Bird) Sing() {
	fmt.Println("叽叽喳喳~")
}

/*
	因为Singer接口只包含一个Sing方法，
	所以只需要给Bird结构体添加一个Sing方法就可以满足Singer接口的要求。
	这样就称为Bird实现了Singer接口。
*/

type Cat struct{}

func (c Cat) Say() {
	fmt.Println("喵喵喵~")
}

type Dog struct{}

func (d Dog) Say() {
	fmt.Println("汪汪汪~")
}

type Sheep struct{}

func (s Sheep) Say() {
	fmt.Println("咩咩咩~")
}

// MakeCatHungry 猫饿了会喵喵喵~
func MakeCatHungry(c Cat) {
	c.Say()
}

// MakeSheepHungry 羊饿了会咩咩咩~
func MakeSheepHungry(s Sheep) {
	s.Say()
}

// 假设动物越来越多 该怎么扩展
// 可以定义一个接口 实现一个会叫的类型 只要是实现了此类的接口 都有Say方法

type Sayer interface {
	Say()
}

// 所以可以定义一个通用的MakeHungry,接受参数为Sayer类型从参数
func MakeHungry(s Sayer) {
	s.Say()
}

// go语言中 只要一个类型实现了接口中规定的所有方法，那么它就实现了这个接口。

// 值接收者和指针接收者
type Mover interface {
	Move()
}

// 值接收者
type Dog2 struct{}

func (d Dog2) Move() {
	fmt.Println("dog2 is moving")
}

// 使用值接收者实现接口之后，不管是结构体类型还是对应的结构体指针类型的变量都可以赋值给该接口变量。

// 指针接收者
type Cat2 struct{}

func (c *Cat2) Move() {
	fmt.Println("cat2 is Moving")
}

// 对于值接收者实现的接口，无论使用值类型还是指针类型都没有问题。

func main() {
	c := Cat{}
	// c.Say()

	d := Dog{}
	// d.Say()

	s := Sheep{}
	// s.Say()
	MakeHungry(c)
	MakeHungry(d)
	MakeHungry(s)

	var m Mover

	var d1 = Dog2{}
	m = d1
	m.Move()

	var d2 = &Dog2{}
	m = d2
	m.Move()

	var c1 = &Cat2{}
	m = c1
	m.Move()

	// var c2 = Cat2{}
	// m = c2 // 此时c2不能赋值给m 不能将c2当作Mover类型

}
