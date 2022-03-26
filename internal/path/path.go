package path

import (
	"fmt"
	"os"
)

type File struct {
	name     string
	children []File
}

func GenerateTree(dir string) (tree []File) {
	f, err := os.Open(dir)

	if err != nil {
		fmt.Println(err)
		return
	}

	files, err := f.Readdir(0)

	if err != nil {
		fmt.Println(err)
		return
	}

	tree = files

	return
}
