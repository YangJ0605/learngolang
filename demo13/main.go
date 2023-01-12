package main

import (
	"fmt"
	"unsafe"
)

// 结构体 struct 值类型
// go语言中通过struct来实现面向对象

/*
type 类型名 struct {
    字段名 字段类型
    字段名 字段类型
    …
}

*/

// type person struct {
// 	name string
// 	age int8
// 	city string
// }

type person struct {
	name, city string
	age        int8
}

func main() {
	var p1 person
	p1.age = 18
	p1.name = "CC"
	p1.city = "shenzhen"
	fmt.Printf("p1: %v\n", p1)
	fmt.Printf("p1: %#v\n", p1)

	// 匿名结构体 只使用一次
	var user struct {
		name string
		age  int8
	}
	user.age = 18
	user.name = "dd"
	fmt.Printf("user: %#v\n", user)

	// 创建指针类型结构体
	var p2 = new(person)
	fmt.Printf("p2的type: %T\n", p2)
	fmt.Printf("p2: %p %#v\n", p2, p2)
	fmt.Printf("p2.age: %v\n", p2.age)

	// 取结构体的地址实例化
	// 使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作
	p3 := &person{}
	fmt.Printf("p3: %T\n", p3)
	fmt.Printf("p3: %#v\n", p3)
	p3.age = 20
	p3.name = "ff"
	p3.city = "changsha"
	// p3.name = "ff" 底层相当于(*p3).name = "ff" go语言的语法糖罢了
	fmt.Printf("p3: %#v\n", p3)

	// 结构体初始化
	var p4 person
	// 没有初始化的结构体，其成员变量都是对应其类型的零值。
	fmt.Printf("p4: %#v\n", p4)

	// 使用键值对初始化
	p5 := person{
		name: "xiaojiu",
		city: "shanxi",
		age:  24,
	}
	fmt.Printf("p5: %#v\n", p5)
	// 对结构体指针进行键值对初始化
	p6 := &person{
		name: "hh",
		city: "lishui",
		age:  24,
	}
	fmt.Printf("p6: %#v\n", p6)
	// 某些字段可以省略不写 没写的默认都是零值
	p7 := &person{
		city: "shanghai",
	}
	fmt.Printf("p7: %#v\n", p7)
	// 初始化的时候不写键，直接写值
	// 这种方式必须全部写所有字段，并且填充顺序保持一致与结构体的声明顺序
	p8 := &person{
		"hq",
		"shenzhen",
		24,
	}
	fmt.Printf("p8: %#v\n", p8)

	// 结构体内存布局 结构体占用一块连续的内存。内存对齐
	type test struct {
		a int8
		b int8
		c int8
		d int8
	}

	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a: %p\n", &n.a)
	fmt.Printf("n.b: %p\n", &n.b)
	fmt.Printf("n.c: %p\n", &n.c)
	fmt.Printf("n.d: %p\n", &n.d)

	// 空结构体 空结构体是不占用空间的。
	var v struct{}
	fmt.Printf("unsafe.Sizeof(v): %v\n", unsafe.Sizeof(v)) //0

	// 面试题
	test1()

	// 构造函数 go语言中没有构造函数的
	p9 := newPerson("张三", "沙河", 90)
	fmt.Printf("p9: %#v\n", p9)
	p9.Dream()
	p9.SetAge(88)
	fmt.Printf("p9: %#v\n", p9)
	p9.SetAge2(80)
	fmt.Printf("p9: %#v\n", p9)

	// 方法和接收者
	// 定义格式如下
	/*
				func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
		    函数体
		}
	*/

	var m1 MyInt
	m1.SayHello()

	// 结构体的匿名字段 选用类型作为字段名字
	type Person2 struct {
		int
		string
	}

	p10 := Person2{
		20,
		"cc",
	}
	fmt.Printf("p10: %#v\n", p10)

	// 嵌套结构体

	type Address struct {
		Province   string
		City       string
		CreateTime string
	}

	type User struct {
		Name    string
		Gender  string
		Address Address
	}

	user1 := User{
		Name:   "cc",
		Gender: "男",
		Address: Address{
			Province: "广东",
			City:     "深圳",
		},
	}
	fmt.Printf("user1: %#v\n", user1)
	fmt.Printf("user1.Address.City: %v\n", user1.Address.City)

	// 嵌套匿名字段
	type Email struct {
		Account    string
		CreateTime string
	}

	type User2 struct {
		Name   string
		Gender string
		Address
		Email
	}

	user2 := User2{
		Name:   "dd",
		Gender: "女",
		Address: Address{
			Province: "广东",
			City:     "广州",
		},
	}
	// 嵌套匿名字段不需要去嵌套的匿名字段中查找
	fmt.Printf("user2.Address.City: %v\n", user2.Address.City)
	fmt.Printf("user2.City: %v\n", user2.City)

	// 嵌套匿名字段冲突
	// user2.CreateTime = "2000" // 报错 无法直接读取
	user2.Email.CreateTime = "1999"
	user2.Address.CreateTime = "1996"

	fmt.Printf("user2: %#v\n", user2)

}

/*
 因为struct是值类型，如果结构体比较复杂的话，
 值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。
*/
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

// 非指针类型的接收者
func (p person) Dream() {
	fmt.Printf("p.name 做梦 %v\n", p.name)
}

// 指针类型的接收者 指针类型可以修改这个指针的其他任意成员
func (p *person) SetAge(newAge int8) {
	p.age = newAge
}

// 值类型的接收者 无法修改 只能获取
func (p person) SetAge2(newAge int8) {
	p.age = newAge
}

// 任意类型添加方法
type MyInt int

func (i MyInt) SayHello() {
	fmt.Println("int 自定义类型加一个方法")
}

/*
	什么时候应该使用指针类型接收者
	需要修改接收者中的值
	接收者是拷贝代价比较大的大对象
	保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

*/

func test1() {
	type student struct {
		name string
		age  uint8
	}

	m := make(map[string]*student, 10)
	stus := []student{
		{name: "xiaozou", age: 18},
		{name: "xiaoyang", age: 19},
		{name: "xiaoliao", age: 20},
	}
	/*
		for-range其实是语法糖，内部调用还是for循环，初始化会拷贝带遍历的列表（如array，slice，map），然后每次遍历的v都是对同一个元素的遍历赋值。 也就是说如果直接对v取地址，最终只会拿到一个地址，而对应的值就是最后遍历的那个元素所附给v的值。
		https://zhuanlan.zhihu.com/p/105435646
		采用临时变量
		或者直接索引
	*/
	for index, stu := range stus {
		// m[stu.name] = &stu

		/*
				temp := stu
			m[stu.name] = &temp
		*/
		m[stu.name] = &stus[index]
	}

	for k, v := range m {
		fmt.Printf("k: %v, value: %v\n", k, v.name)
	}

	// var sliceA = []student{
	// 	{name: "xiaozou", age: 18},
	// 	{name: "xiaoyang", age: 19},
	// 	{name: "xiaoliao", age: 20},
	// }
	// for i := 0; i < len(sliceA); i++ {
	// 	fmt.Printf("sliceA: %p\n", sliceA)
	// 	fmt.Printf("sliceA[i]: %p\n", &sliceA[i])
	// }

	// for index, stu := range sliceA {
	// 	fmt.Printf("sliceA: %p\n", sliceA)
	// 	fmt.Printf("sliceItem: %p\n", &stu)
	// 	fmt.Printf("slice[index]: %p\n", &sliceA[index])
	// }

	/*
		遍历前对v做了拷贝，所以期间对原来v的修改不会反映到遍历中
	*/
	v := []int{1, 2, 3}
	for i := range v {
		fmt.Println("rang~")
		v = append(v, i)
	}
}
