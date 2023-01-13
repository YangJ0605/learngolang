package main

import (
	"fmt"
	"time"
)

// 标准库 time包的用法

func main() {
	now := time.Now()
	fmt.Printf("now: %v\n", now)
	fmt.Printf("now.Year(): %v\n", now.Year())
	fmt.Printf("now.Month(): %v\n", now.Month())
	fmt.Println(now.Date())
	fmt.Printf("now.Day(): %v\n", now.Day())

	fmt.Printf("now.Hour(): %v\n", now.Hour())
	fmt.Printf("now.Minute(): %v\n", now.Minute())
	fmt.Printf("now.Second(): %v\n", now.Second())

	fmt.Printf("now.Weekday(): %d\n", now.Weekday())

	// 时间戳
	fmt.Printf("now.Unix(): %v\n", now.Unix())
	// 将时间转换成具体时间
	fmt.Printf("time.Unix(1673600061+3600, 0): %v\n", time.Unix(1673600061+3600, 0))

	// 时间间隔
	time.Sleep(1 * time.Second)
	n := 1
	// 传入int64
	time.Sleep(time.Duration(n) * time.Second)

	// 加时间
	fmt.Printf("now.Add(time.Hour): %v\n", now.Add(time.Hour))

	// 定时器
	// for k := range time.Tick(time.Second) {
	// 	fmt.Printf("k: %v\n", k)
	// }

	// 时间格式化
	// fmt.Printf("now.Format(\"Y-m-d\"): %v\n", now.Format("Y-m-d"))
	fmt.Printf("now.Format(\"2006-01-02\"): %v\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("now.Format(\"2006.01.02 03:04:05\"): %v\n", now.Format("2006.01.02 03:04:05.000 PM cc"))

	// 解析字符串类型时间
	timeObj, err := time.Parse("2006-01-02 15:04:05", "2023-01-13 17:10:43")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("timeObj: %v\n", timeObj)

	// 获取时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	timeObj2, err := time.ParseInLocation("2006-01-02 15:04:05", "2023-01-13 17:10:43", loc)
	if err != nil {
		fmt.Printf("err2: %v\n", err)
		return
	}
	fmt.Printf("timeObj2: %v\n", timeObj2)
}
