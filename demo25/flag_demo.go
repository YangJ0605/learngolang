package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "张三", "姓名")
	age := flag.Int("age", 18, "年龄")
	married := flag.Bool("married", false, "婚否")
	delay := flag.Duration("d", 0, "时间间隔")

	flag.Parse()

	fmt.Printf("name: %v\n", *name)
	fmt.Printf("age: %v\n", *age)
	fmt.Printf("married: %v\n", *married)
	fmt.Printf("delay: %v\n", *delay)

	//定义命令行参数方式1
	// var name string
	// var age int
	// var married bool
	// var delay time.Duration
	// flag.StringVar(&name, "name", "张三", "姓名")
	// flag.IntVar(&age, "age", 18, "年龄")
	// flag.BoolVar(&married, "married", false, "婚否")
	// flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	// //解析命令行参数
	// flag.Parse()
	// fmt.Println(name, age, married, delay)
	// //返回命令行参数后的其他参数
	// fmt.Println(flag.Args())
	// //返回命令行参数后的其他参数个数
	// fmt.Println(flag.NArg())
	// //返回使用的命令行参数个数
	// fmt.Println(flag.NFlag())
}
