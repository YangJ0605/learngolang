package management

import (
	"fmt"
	stu "stu_management/student"
	"stu_management/utils"

	"github.com/fatih/color"
)

type StuMan struct {
	students []*stu.Student
}

func (s *StuMan) CheckHasStu(id int) bool {
	for _, v := range s.students {
		if v.Id == id {
			return true
		}
	}
	return false
}

func (s *StuMan) ShowStudents() {
	if len(s.students) == 0 {
		color.Red("学生列表为空")
		return
	}
	for _, stu := range s.students {
		str := fmt.Sprintf("姓名: %v, 学号: %v, 班级: %v\n", stu.Name, stu.Id, stu.Class)
		color.Cyan(str)
	}
}

func (s *StuMan) AddStudent(stu *stu.Student) {
	s.students = append(s.students, stu)
	utils.Success("添加成功")
}

func (s *StuMan) ModifyStudent(newStu *stu.Student) {
	for index, stu := range s.students {
		if stu.Id == newStu.Id {
			s.students[index] = newStu
			utils.Success("修改成功")
			return
		}
	}
}

func (s *StuMan) DeleteStudent(id int) {
	var idx int
	for index, stu := range s.students {
		if stu.Id == id {
			idx = index
			break
		}
	}
	s.students = append(s.students[:idx], s.students[idx+1:]...)
	utils.Success("删除成功")
}

func NewStuMan() *StuMan {
	return &StuMan{
		students: make([]*stu.Student, 0, 100),
	}
}
