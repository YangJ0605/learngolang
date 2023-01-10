package main

import "fmt"

// 数组相关内容

func main() {
	var a [3]int
	var b [4]int

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	// 数组的初始化
	cityArr := [4]string{"北京", "上海", "广州", "深圳"}
	fmt.Printf("cityArr: %v\n", cityArr)
	fmt.Printf("cityArr[0]: %v\n", cityArr[0])
	fmt.Printf("cityArr[3]: %v\n", cityArr[3])

	// 编译器推到数组的长度
	boolArr := [...]bool{true, false, false, true}
	fmt.Printf("boolArr: %v\n", boolArr)

	// 使用索引值方式初始化
	langArr := [...]string{1: "golang", 2: "javascript", 3: "typescript", 5: "java"}
	fmt.Printf("langArr: %v\n", langArr)
	fmt.Printf("langArr的类型: %T\n", langArr)

	// 数组遍历
	for i := 0; i < len(cityArr); i++ {
		fmt.Printf("城市: %v\n", cityArr[i])
	}
	for index, city := range cityArr {
		fmt.Printf("index: %v-%v\n", index, city)
	}

	// 二维数组
	a2 := [...][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Printf("a2: %v\n", a2)
	for _, v := range a2 {
		fmt.Printf("v: %v\n", v)
		for _, v2 := range v {
			fmt.Printf("v2: %v\n", v2)
		}
	}

	// 数组是值类型 赋值开启一个新的内存空间
	arr1 := [...]int{1, 2, 3}
	changeArr(arr1)
	fmt.Printf("arr1: %v\n", arr1)
	arr2 := arr1
	arr2[0] = 999
	fmt.Printf("arr2: %v\n", arr2)
	fmt.Printf("arr1: %v\n", arr1)

	// 求和
	sum()

	findTarget()
}

func changeArr(arr [3]int) {
	arr[0] = 100
	fmt.Printf("函数中的数组: %v\n", arr)
}

func sum() {
	arr := [...]int{1, 3, 5, 7, 8}
	res := 0
	for _, v := range arr {
		res += v
	}
	fmt.Printf("res: %v\n", res)
}

func findTarget() {
	arr := [...]int{1, 3, 5, 7, 8}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == 8 {
				fmt.Printf("(%v, %v)\n", i, j)
			}
		}
	}

}
