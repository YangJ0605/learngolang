package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
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
	if len(s.students) == 0 {
		color.Red("学生列表为空")
		return
	}
	for _, stu := range s.students {
		str := fmt.Sprintf("姓名: %v, 学号: %v, 班级: %v\n", stu.name, stu.id, stu.class)
		color.Cyan(str)
	}
}

func (s *StuMan) addStudent(stu *Student) {
	s.students = append(s.students, stu)
	success("添加成功")
}

func (s *StuMan) modifyStudent(newStu *Student) {
	for index, stu := range s.students {
		if stu.id == newStu.id {
			s.students[index] = newStu
			success("修改成功")
			return
		}
	}
}

func (s *StuMan) deleteStudent(id int) {
	var idx int
	for index, stu := range s.students {
		if stu.id == id {
			idx = index
			break
		}
	}
	s.students = append(s.students[:idx], s.students[idx+1:]...)
	success("删除成功")
}

func NewStuMan() *StuMan {
	return &StuMan{
		students: make([]*Student, 0, 100),
	}
}

func showMenu() {
	info("欢迎使用学员管理系统")
	info("1. 展示学生列表")
	info("2. 添加学生")
	info("3. 编辑学生")
	info("4. 删除学生")
	info("5. 退出系统")
}

func getStudentInput() *Student {
	var (
		name, class string
		id          int
	)
	fmt.Println("请输入学生信息")
	fmt.Println("学号：")
	fmt.Scanln(&id)
	fmt.Println("姓名：")
	fmt.Scanln(&name)
	fmt.Println("班级")
	fmt.Scanln(&class)

	stu := NewStudent(name, class, id)
	return stu
}

func success(str string) {
	color.New(color.Bold, color.FgGreen).PrintlnFunc()(str)
}

func warning(str string) {
	color.New(color.Bold, color.FgYellow).PrintlnFunc()(str)
}

func error(str string) {
	color.New(color.Bold, color.FgRed).PrintlnFunc()(str)
}

func info(str string) {
	color.New(color.BlinkSlow, color.FgHiWhite).PrintlnFunc()(str)
}

func main() {
	stuMan := NewStuMan()
	showMenu()
	for {
		var input int
		fmt.Println("请输入你的操作")
		fmt.Scanln(&input)
		switch input {
		case 1:
			info("学生信息如下")
			stuMan.showStudents()
		case 2:
			info("添加学生操作")
			stu := getStudentInput()
			isExist := stuMan.checkHasStu(stu.id)
			if isExist {
				// fmt.Println("该学号已经存在，请重新添加")
				error("该学号已经存在，请重新添加")
			} else {
				stuMan.addStudent(stu)
			}

		case 3:
			info("编辑学生操作")
			stu := getStudentInput()
			isExist := stuMan.checkHasStu(stu.id)
			if isExist {
				stuMan.modifyStudent(stu)
			} else {
				// fmt.Println("该学号不存在，请重试")
				warning("该学号不存在，请重试")
			}
		case 4:
			info("删除学生操作")
			var id int
			// fmt.Println("请输入学生的学号：")
			info("请输入学生的学号：")
			fmt.Scanln(&id)
			isExist := stuMan.checkHasStu(id)
			if isExist {
				stuMan.deleteStudent(id)
			} else {
				// fmt.Println("该学号不存在，请重试")
				warning("该学号不存在，请重试")
			}
		case 5:
			os.Exit(0)
		default:
			// fmt.Println("您输入有误，【1-5】")
			error("您输入有误，【1-5】")
		}

	}
}
