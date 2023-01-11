package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// map的基本知识 引用类型 nil

// map[KeyType]ValueType

func main() {
	var s []int // var声明的切片无需初始化就可以使用append
	s = append(s, 1, 2, 3)
	fmt.Printf("s: %v\n", s)

	// 仅仅var 声明没初始化
	var map1 map[string]string
	fmt.Printf("map1: %v\n", map1 == nil) // true

	// 声明方式
	scoreMap := make(map[string]int, 8)
	scoreMap["cc"] = 99
	scoreMap["dd"] = 95
	fmt.Printf("scoreMap: %#v 类型是: %T\n", scoreMap, scoreMap)
	fmt.Printf("scoreMap[\"cc\"]: %v\n", scoreMap["cc"])

	userInfo := map[string]string{
		"username": "CC",
		"age":      "18",
	}
	fmt.Printf("userInfo: %v\n", userInfo)

	// 判断map中是否存在某个键
	value, ok := userInfo["age"]
	if ok {
		fmt.Printf("value: %v\n", value)
	} else {
		fmt.Printf("\"cc\": %v不存在\n", "cc")
	}

	// map的遍历
	for k, v := range userInfo {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}

	// 删除某个值
	delete(userInfo, "age")
	fmt.Printf("userInfo: %v\n", userInfo)

	// 元素为map的切片
	mapSlice := make([]map[string]string, 3)
	// 必须先初始化map
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "cc"
	mapSlice[0]["password"] = "147258"
	// fmt.Printf("mapSlice: %v\n", mapSlice)
	for index, v := range mapSlice {
		fmt.Printf("index: %v, value: %v\n", index, v)
	}

	// 值为切片的map
	sliceMap := make(map[string][]string, 3)
	sliceMap["h"] = []string{"唱", "跳", "rapper"}
	for key, value := range sliceMap {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}

	// 指定顺序遍历map
	test1()
	//
	test2()
}

func test1() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}

	keys := make([]string, 0, 100)

	for key := range scoreMap {
		keys = append(keys, key)
	}

	fmt.Printf("keys: %v\n", keys)
	sort.Strings(keys)
	fmt.Printf("keys: %v\n", keys)

	for _, key := range keys {
		fmt.Printf("stu: %v, score: %v\n", key, scoreMap[key])
	}

	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	fmt.Printf("m: %v\n", m["q1mi"])
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"]) // 133
}

func test2() {
	str := "how do you do"
	var countMap = make(map[string]int, 10)
	values := strings.Split(str, " ")
	for _, v := range values {
		count, ok := countMap[v]
		if !ok {
			countMap[v] = 1
		} else {
			countMap[v] = count + 1
		}
	}
	fmt.Printf("countMap: %v\n", countMap)
}
