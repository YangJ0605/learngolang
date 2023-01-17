package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("index: %v args[index]=%v type: %T\n", index, arg, arg)
		}
	}

}
