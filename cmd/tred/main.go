package main

import (
	"fmt"
	"lsroa/tred/internal/path"
)

func main() {
	ls := path.GenerateTree("./examples/sample-project")

	for _, file := range ls {
		fmt.Println(file.Name(), file.IsDir())
	}
}
