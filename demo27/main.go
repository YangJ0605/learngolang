package main

import (
	"fmt"
	"os"
)

func create() {
	// 覆盖式创建
	// f, err := os.Create("1.txt")
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// } else {
	// 	fmt.Printf("f.Name(): %v\n", f.Name())
	// }

	// MkdirAll 创建多个层级目录
	// err = os.Mkdir("test", os.ModePerm) // 777权限
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }

}

func remove() {
	// err := os.Remove("1.txt")
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }

	// 删除所有的文件 remove只能删除单个
	err := os.RemoveAll("a")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func getWd() {
	dir, _ := os.Getwd()
	fmt.Printf("dir: %v\n", dir)

	// os.Chdir("d:/")
	fmt.Printf("os.TempDir(): %v\n", os.TempDir())
}

func rename() {
	fmt.Printf("os.Rename(\"1.txt\", \"2.txt\"): %v\n", os.Rename("1.txt", "2.txt"))
}

func readFile() {
	b, _ := os.ReadFile("2.txt")
	fmt.Printf("b: %v\n", string(b))
}

func writeFile() {
	os.WriteFile("3.txt", []byte("cccc"), os.ModePerm)
}

func main() {
	// remove()
	// getWd()

	// rename()
	// readFile()
	writeFile()
}
