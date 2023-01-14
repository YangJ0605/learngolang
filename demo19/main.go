package main

import (
	"fmt"
	"reflect"
)

// 反射的基本内容

func reflectType(x interface{}) {
	// go语言中获取类型 要么借助类型断言 要么借助反射
	v := reflect.TypeOf(x)
	// fmt.Printf("v: %v, type: %T\n", v, v)
	fmt.Printf("v.Name(): %v, v.Kind(): %v\n", v.Name(), v.Kind())

}

type Person struct {
}

type Animal struct{}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("v: %v, type: %T\n", v, v)

	k := v.Kind()
	fmt.Printf("k: %v\n", k)
	switch k {
	case reflect.Float32:
		ret := float32(v.Float())
		fmt.Printf("ret: %v, type: %T\n", ret, ret)
	case reflect.Int32:
		ret := int32(v.Int())
		fmt.Printf("ret: %v, type: %T\n", ret, ret)
	}
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	// k := v.Kind()
	// 根据指针取值
	k := v.Elem().Kind()

	switch k {
	case reflect.Int32:
		v.Elem().SetInt(1000)

	}
}

func main() {
	// reflectType(1)
	// reflectType("ccc")
	// reflectType(true)
	// reflectType(1.234)
	// // 切片 map 数组获取到的Name为空
	// reflectType([]string{})
	// reflectType([3]string{})
	// reflectType(make(map[string]string, 0))

	// var p Person
	// reflectType(p)

	// var a Animal
	// reflectType(a)

	// reflectValue(1)
	// reflectValue("ccc")
	// reflectValue(1.234)
	// var ff float32 = 1.124
	// reflectValue(ff)

	var i int32 = 88
	var c int = 10000
	reflectSetValue(&i)
	reflectSetValue(&c)
	fmt.Printf("i: %v, type: %T\n", i, i)
	fmt.Printf("c: %v, type: %T\n", c, c)
}
