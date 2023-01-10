package main

import "fmt"

//流程控制

func gotoDemo() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto breakTag
			}
			fmt.Printf("i - j: %v - %v\n", i, j)
		}
	}
breakTag:
	fmt.Println("breaktag")
}

func switchDemo() {
	// fallthrough
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}

func print99() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", j, i, i*j)
		}
		fmt.Println("")
	}
}

func main() {
	// score := 65

	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score >= 75 {
		fmt.Println("B")

	} else {
		fmt.Println("C")

	}
	// for循环的基本写法
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("i: %v\n", i)
	// }

	// for循环 省略初始语句，但是必须保留分号
	// i := 0
	// for ; i < 10; i++ {
	// 	fmt.Printf("i: %v\n", i)
	// }

	// for循环 省略初始语句和结束语句
	// i := 0
	// for i < 10 {
	// 	fmt.Printf("i: %v\n", i)
	// 	i++
	// }

	// 无限循环
	// for {
	// 	fmt.Println("无限循环")
	// }

	// break 跳出for循环
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("i: %v\n", i)
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// continue 跳出本次继续下一次
	// for i := 0; i < 10; i++ {

	// 	if i == 5 {
	// 		continue
	// 	}

	// 	fmt.Printf("i: %v\n", i)
	// }

	// goto 跳转到指定标签
	gotoDemo()

	// for range 遍历字符串 切片 数组 map等
	// for _, v := range "cc" {
	// 	fmt.Printf("v: %c\n", v)
	// }

	// switch case
	// finger := 3
	// switch finger {
	// case 1:
	// 	fmt.Println("大拇指")
	// case 2:
	// 	fmt.Println("食指")
	// case 3:
	// 	fmt.Println("中指")
	// case 4:
	// 	fmt.Println("无名指")
	// case 5:
	// 	fmt.Println("小拇指")
	// default:
	// 	fmt.Println("无效的输入")
	// }

	// case 一次判断多个
	// num := 5
	// switch num {
	// case 1, 3, 5:
	// 	fmt.Println("奇数")
	// case 2, 4, 6, 8:
	// 	fmt.Println("偶数")
	// }

	// case中使用表达式
	// age := 18

	// switch {
	// case age >= 18:
	// 	fmt.Println("成年了")
	// case age < 18:
	// 	fmt.Println("未成年")
	// default:
	// 	fmt.Println("cccc")
	// }

	// fallthrough 语法可以执行满足条件的case的下一个case
	switchDemo()

	// 打印99乘法表
	print99()
}
