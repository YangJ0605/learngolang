package main

import (
	"fmt"
	"os"
)

// 程序提供展示学生列表、添加学生、编辑学生信息、删除学生

type Student struct {
	name, class string
	id          int
}

func NewStudent(name, class string, id int) *Student {
	return &Student{
		name,
		class,
		id,
	}
}

type StuMan struct {
	students []*Student
}

func (s *StuMan) checkHasStu(id int) bool {
	for _, v := range s.students {
		if v.id == id {
			return true
		}
	}
	return false
}

func (s *StuMan) showStudents() {
	for _, stu := range s.students {
		fmt.Printf("姓名: %v, 学号: %v, 班级: %v\n", stu.name, stu.id, stu.class)
	}
}

func (s *StuMan) addStudent(stu *Student) {
	s.students = append(s.students, stu)
}

func (s *StuMan) modifyStudent(newStu *Student) {
	for index, stu := range s.students {
		if stu.id == newStu.id {
			s.students[index] = newStu
			return
		}
	}
}

func (s *StuMan) deleteStudent(id)

func NewStuMan() *StuMan {
	return &StuMan{
		students: make([]*Student, 0, 100),
	}
}

func showMenu() {
	fmt.Println("欢迎使用学员管理系统")
	fmt.Println("1. 展示学生列表")
	fmt.Println("2. 添加学生")
	fmt.Println("3. 编辑学生")
	fmt.Println("4. 删除学生")
	fmt.Println("5. 退出系统")
}

func main() {
	stuMan := NewStuMan()
	for {
		showMenu()
		var input int
		fmt.Println("请输入你的操作")
		fmt.Scanln(&input)
		switch input {
		case 1:
			stuMan.showStudents()
		case 2:
			stuMan.addStudent()
		case 3:
		case 4:
		case 5:
			os.Exit(0)
		default:
			fmt.Println("您输入有误，【1-5】")
		}

	}
}
