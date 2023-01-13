package main

import (
	"fmt"
	"os"

	utils "stu_management/utils"

	stuM "stu_management/management"
)

// 程序提供展示学生列表、添加学生、编辑学生信息、删除学生

func showMenu() {
	info := utils.Info
	info("欢迎使用学员管理系统")
	info("1. 展示学生列表")
	info("2. 添加学生")
	info("3. 编辑学生")
	info("4. 删除学生")
	info("5. 退出系统")
}

func main() {
	stuMan := stuM.NewStuMan()
	showMenu()
	for {
		var input int
		fmt.Println("请输入你的操作")
		fmt.Scanln(&input)
		switch input {
		case 1:
			utils.Info("学生信息如下")
			stuMan.ShowStudents()
		case 2:
			utils.Info("添加学生操作")
			stu := utils.GetStudentInput()
			isExist := stuMan.CheckHasStu(stu.Id)
			if isExist {
				// fmt.Println("该学号已经存在，请重新添加")
				utils.Error("该学号已经存在，请重新添加")
			} else {
				stuMan.AddStudent(stu)
			}

		case 3:
			utils.Info("编辑学生操作")
			stu := utils.GetStudentInput()
			isExist := stuMan.CheckHasStu(stu.Id)
			if isExist {
				stuMan.ModifyStudent(stu)
			} else {
				// fmt.Println("该学号不存在，请重试")
				utils.Warning("该学号不存在，请重试")
			}
		case 4:
			utils.Info("删除学生操作")
			var id int
			// fmt.Println("请输入学生的学号：")
			utils.Info("请输入学生的学号：")
			fmt.Scanln(&id)
			isExist := stuMan.CheckHasStu(id)
			if isExist {
				stuMan.DeleteStudent(id)
			} else {
				// fmt.Println("该学号不存在，请重试")
				utils.Warning("该学号不存在，请重试")
			}
		case 5:
			os.Exit(0)
		default:
			// fmt.Println("您输入有误，【1-5】")
			utils.Error("您输入有误，【1-5】")
		}

	}
}
