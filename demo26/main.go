package main

import "fmt"

// 泛型函数
func PrinIntTypeSlice[T int | int64 | string](slice []T) {
	fmt.Printf("slice: %T\n", slice)
}

// 泛型切片
type mySlice[T int | string] []T

// 泛型map
type myMap[K string, V any] map[K]V

// ~代表底层数据的意思
type Num interface {
	~int | ~int32 | ~int64 | uint8
}

type Str interface {
	~string
}

type NumStr interface {
	Num | Str
}

type myString string
type myUint8 uint8

func Test[T NumStr](str T) {
	fmt.Printf("str: %T\n", str)
}

func main() {
	PrinIntTypeSlice[int]([]int{122})
	PrinIntTypeSlice[int64]([]int64{144})
	PrinIntTypeSlice[string]([]string{"cc"})

	slice := mySlice[string]{"cc"}

	PrinIntTypeSlice(slice)

	m1 := myMap[string, string]{}
	fmt.Printf("m1: %T\n", m1)

	var ss myString = "ccc"
	Test(ss)

}
