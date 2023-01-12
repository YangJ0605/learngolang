package main

import (
	"fmt"
	"math/rand"
	"time"
)

// type Person struct {
// 	Name string
// 	age uint8
// }

type Student struct {
	ID    int
	Score int
	Name  string
	Age   int
}

type List struct {
	Students []Student
}

func (l *List) showList() {
	for _, v := range l.Students {
		fmt.Printf("学生: %v\n", v)
	}
}

func (l *List) addStudents(students []Student) {
	if l.Students == nil {
		l.Students = make([]Student, 0, len(students)*2)
	}
	stus := make([]Student, len(students))
	copy(stus, students)
	l.Students = append(l.Students, stus...)

}

func (l *List) editStudent(index int, student Student) {
	l.Students[index] = student
}

func (l *List) deleteStudent(index int) {
	l.Students = append(l.Students[:index], l.Students[index+1:]...)
}

func main() {
	// var s Student
	// fmt.Printf("s: %v\n", s == nil)
	var input int
	fmt.Println("请选择操作：")
	// fmt.Scanf("%d\n", &input)
	fmt.Scanln(&input)
	fmt.Printf("input: %v\n", input)
	rand.Seed(time.Now().Unix()) //生成随机数种子
	// fmt.Printf("rand.Intn(100): %v\n", rand.Intn(100))
	list := List{}
	stus := make([]Student, 0, 10)
	for i := 0; i < 10; i++ {
		stu := Student{
			ID:    i,
			Score: rand.Intn(100),
			Name:  fmt.Sprintf("stu%02d", i),
			Age:   20 - rand.Intn(10),
		}
		stus = append(stus, stu)
	}
	list.addStudents(stus)
	stus[2].Age = 99999
	list.showList()
	fmt.Println("==========")
	list.editStudent(2, Student{
		ID: 2222,
	})
	list.showList()
	fmt.Println("==========")
	list.deleteStudent(3)
	list.showList()
}
