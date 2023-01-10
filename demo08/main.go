package main

import (
	"fmt"
	"sort"
)

// 切片的基本内容

func main() {
	// 定义切片
	// var a []string
	// var b []int
	// var c = []bool{true, false}

	// fmt.Printf("a: %v\n", a)
	// fmt.Printf("b: %v\n", b)
	// fmt.Printf("c: %v\n", c)

	// 基于数组得到切片
	a := [5]int{1, 2, 3, 4, 5}
	b := a[:4]
	fmt.Printf("b: %v, 类型：%T\n", b, b)
	// 切片是引用底层数组
	a[3] = 555
	fmt.Printf("b: %v\n", b)
	c := b[1:]
	fmt.Printf("c: %v\n", c)

	// make函数构造切片
	d := make([]int, 5, 10)
	fmt.Printf("d: %v, 类型：%T, 容量：%v, 长度：%v\n", d, d, cap(d), len(d))

	// nil 此时a2不占内存
	var a2 []int
	var b2 = []int{}
	fmt.Printf("a2: %v, len: %v, cap: %v\n", a2, len(a2), cap(a2))
	if a2 == nil {
		fmt.Println("a2 is nil")
	}

	fmt.Printf("b2: %v, len: %v, cap: %v\n", b2, len(b2), cap(b2))
	if b2 != nil {
		fmt.Println("b2 is not nil")
	}

	c2 := make([]int, 0)
	fmt.Printf("c2: %v, len: %v, cap: %v\n", c2, len(c2), cap(c2))
	if c2 != nil {
		fmt.Println("c2 is not nil")
	}
	// 判断一个切片为空用长度是否为0来判断

	// 切片的赋值拷贝
	a3 := make([]int, 3)
	b3 := a3
	b3[0] = 100
	fmt.Printf("a3: %v\n", a3)
	fmt.Printf("b3: %v\n", b3)

	// 切片的遍历与数组一样
	// for 循环 和 for range

	// append方法 切片的扩容
	var s []int // 此时没有申请内存，不能通过索引赋值
	// 切片在声明之后一定要初始化才能使用
	s = []int{1, 2, 3}
	fmt.Printf("s: %v, len: %v, cap: %v\n", s, len(s), cap(s))

	s = append(s, 10)
	fmt.Printf("s: %v, len: %v, cap: %v\n", s, len(s), cap(s))
	fmt.Printf("s: %v, ptr: %p\n", s, s)

	for i := 0; i < 10; i++ {
		s = append(s, i+10)
		// 每当扩容 地址就会变化
		fmt.Printf("s: %v, len: %v, cap: %v, 地址: %p\n", s, len(s), cap(s), s)
	}

	// append 函数增加多个数
	s2 := []int{1, 2, 3}
	// 三个点
	s2 = append(s2, []int{4, 5, 6}...)
	s2 = append(s2, 7, 8, 9)
	fmt.Printf("s2: %v\n", s2)

	// 使用copy函数复制切片 指向的不是同一个底层数组了
	s3 := []int{1, 2, 3, 4, 5}
	s4 := make([]int, 5, 5)

	copy(s4, s3)
	s4[0] = 1000
	s3[0] = 8888
	fmt.Printf("s3: %v\n", s3)
	fmt.Printf("s4: %v\n", s4)

	// 切片的删除
	s5 := []string{"shanghai", "beijing", "guangzhou", "shenzhen"}
	s5 = append(s5[:1], s5[2:]...)
	fmt.Printf("s5: %v\n", s5)
	//删除切片的公式 append(s[:index], s[index+1:]...)

	// test
	test1()
	arrSort()
}

func test1() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Printf("a: %v\n", a)
}

func arrSort() {
	arr := [...]int{1, 5, 3, 6, 2}

	sort.Ints(arr[:])
	fmt.Printf("arr: %v\n", arr)
}
