package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	root := "emptydir"
	err := cull(root)
	if err != nil {
		panic(err)
	}

}

func cull(root string) (e error) {
	// TODO define cull date

	e = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// TODO if the file is a directory, recursively call cull(file)
		if !info.IsDir() {
			path = "\t" + path
		}
		fmt.Println(path)
		return err

	})

	return e
}
