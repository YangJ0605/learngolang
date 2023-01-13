package utils

import (
	"fmt"
	stu "stu_management/student"
)

func GetStudentInput() *stu.Student {
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

	stu := stu.NewStudent(name, class, id)
	return stu
}
