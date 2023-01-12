package student

func Add(a, b int) int {
	return a + b
}

func init() {
	Make() // 同一个包中的不同文件可以直接调用
}
