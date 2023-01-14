package main

import (
	"fmt"
	"reflect"
)

// 结构体反射

type Student struct {
	Name  string `json:"name"`
	Score int    `json:"score" ini:"s"`
}

// 注意要大写
func (s Student) Study() {
	fmt.Printf(" %v 正在学习~~\n", s.Name)
}

func (s Student) Run() string {
	// fmt.Printf("a: %v\n", a)
	fmt.Printf(" %v 正在run~~\n", s.Name)
	return "cc"
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Printf("t.NumMethod(): %v\n", t.NumMethod())
	// fmt.Printf("v.NumMethod(): %v\n", v.NumMethod())
	// 因为要拿到方法去调用 所以用的是值信息
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("t.Method(i).Name: %v\n", t.Method(i).Name)
		fmt.Printf("methodType: %v\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

func main() {
	stu1 := Student{
		"cc",
		99,
	}
	// t := reflect.TypeOf(stu1)
	// fmt.Printf("t.Name(): %v, t.Kind(): %v\n", t.Name(), t.Kind())
	// fmt.Printf("t.NumField(): %v\n", t.NumField())
	// for i := 0; i < t.NumField(); i++ {
	// 	field := t.Field(i)
	// 	// fmt.Printf("field.Tag: %T\n", field.Tag)
	// 	fmt.Printf("field.Name: %v, index: %d type: %v josn tag: %v\n", field.Name, field.Index, field.Type, field.Tag)
	// 	fmt.Printf("field.Tag.Get(\"json\"): %v\n", field.Tag.Get("json"))
	// 	fmt.Printf("field.Tag.Get(\"ini\"): %v\n", field.Tag.Get("ini"))
	// }
	// filed, ok := t.FieldByName("Name")
	// if ok {
	// 	fmt.Printf("filed.Name: %v, type: %v, tag: %v\n", filed.Name, filed.Type, filed.Tag)
	// }

	printMethod(stu1)
}
