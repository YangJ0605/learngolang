package main

import (
	"encoding/json"
	"fmt"
)

// 结构体的继承

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("a.name %v 会移动\n", a.name)
}

type Dog struct {
	age     int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("d.name %v 会汪汪汪~\n", d.name)
}

// 结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

// 结构体与json序列化
type Student struct {
	ID     int
	Gender string
	Name   string
}

type Class struct {
	Title    string
	Students []*Student
}

func testJson() {
	c := &Class{
		Title:    "116",
		Students: make([]*Student, 0, 200),
	}

	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}

		c.Students = append(c.Students, stu)
	}

	// 序列化 结构体转化成字符切片
	data, err := json.Marshal(c)

	if err != nil {
		fmt.Println("json marshal failed")
		return
	}

	fmt.Printf("json: %s, 类型：%T\n", data, data)
	str := `
	{"Title":"116","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}
	`
	c1 := &Class{}
	// 反序列化
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("c1: %#v\n", c1)
}

// 结构体标签
/*
	Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来。 Tag在结构体字段的后方定义，由一对反引号包裹起来，具体的格式如下：
	`key1:"value1" key2:"value2"`
*/

func testTag() {
	type Student2 struct {
		ID     int    `json:"id" db:"ID"` //通过指定tag实现json序列化该字段时的key
		Gender string //json序列化是默认使用字段名作为key
		name   string //私有不能被json包访问
	}
	s1 := Student2{
		ID:     1,
		Gender: "男",
		name:   "cc",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("marshal 失败")
		return
	}
	fmt.Printf("tag data: %s\n", data)
	s2 := Student2{}
	str := `{"id":1,"Gender":"男"}`
	json.Unmarshal([]byte(str), &s2)
	fmt.Printf("s2: %#v\n", s2)
}

// 结构体和方法补充知识点
/*
	因为slice和map这两种数据类型都包含了指向底层数据的指针，因此我们在需要复制它们时要特别注意。
	我们来看下面的例子：
*/
type Person struct {
	name   string
	age    uint8
	dreams []string
}

func (p *Person) SetDreams(dreams []string) {
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}

func testStruct() {
	p1 := Person{name: "cc", age: 18}
	fmt.Printf("p1: %v\n", p1)
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data)
	fmt.Printf("p1: %v\n", p1)
	data[1] = "不睡觉"

	fmt.Printf("p1: %v\n", p1)
}

func main() {
	d := &Dog{
		age: 4,
		Animal: &Animal{
			name: "乐乐",
		},
	}

	d.wang()
	d.move()

	testJson()

	testTag()

	testStruct()
}
