package student

type Student struct {
	Name, Class string
	Id          int
}

func NewStudent(name, class string, id int) *Student {
	return &Student{
		name,
		class,
		id,
	}
}
