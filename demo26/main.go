package main

import "fmt"

// https://segmentfault.com/a/1190000041634906

// 泛型函数
func PrintIntTypeSlice[T int | int64 | string](slice []T) {
	fmt.Printf("slice: %T\n", slice)
}

// 泛型切片
type mySlice[T int | string] []T

// 泛型map
type myMap[K string | int, V any] map[K]V

// ~代表底层数据的意思
type Num interface {
	~int | ~int32 | ~int64 | uint8
}

type Str interface {
	~string
}

type AB[T int | string] struct{}

// 方法不支持泛型 但是支持从结构体本身获取
func (ab AB[T]) Add(a, b T) T {
	return a + b
}

type NumStr interface {
	Num | Str
}

type A = string // 这仅仅是类型别名 本质还是string类型
type B string   // 这是一个新的类型

type myString string
type myUint8 uint8

// 定义指针类型的泛型 最好加上interface
type NewType[T interface{ *int | *string }] []T

func Test[T NumStr](str T) {
	fmt.Printf("str: %T\n", str)
}

type Int interface {
	~int | int8 | int16 | int32 | int64
}

type Unit interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Float interface {
	float32 | float64
}

type X interface {
	Int | Unit | Float
}

type Slice[T X] []T

func main() {
	var a A
	var b B
	fmt.Printf("a: %T\n", a)
	fmt.Printf("b: %T\n", b)
	PrintIntTypeSlice([]int{122})
	PrintIntTypeSlice([]int64{144})
	// 可以省略不写 会自己推断类型
	// PrintIntTypeSlice[int64]([]int64{144})
	PrintIntTypeSlice([]string{"cc"})

	slice := mySlice[string]{"cc"}

	PrintIntTypeSlice(slice)

	m1 := myMap[string, string]{}
	fmt.Printf("m1: %T\n", m1)

	m2 := myMap[int, string]{
		1: "vv",
		2: "ff",
	}

	fmt.Printf("m2: %v\n", m2)

	n := NewType[*int]{
		new(int),
	}
	fmt.Printf("n: %v\n", n)

	var ss myString = "ccc"
	Test(ss)

	// ~代表底层类型
	var s1 Slice[int]
	type MyInt int
	var s2 Slice[MyInt] // 不能传入此类型 除非加入波浪线

}
