package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	//	root := "/home/mark/coding/go/src"
	root := "."
	err := cull(root)
	// err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
	// 	files = append(files, path)
	// 	return nil
	// })
	if err != nil {
		panic(err)
	}

}

func cull(root string) (e error) {
	var files []string
	// TODO define cull date

	e = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// TODO if the file is a directory, recursively call cull(file)
		files = append(files, path)
		return err

	})

	if e != nil {
		return e
	}

	for _, file := range files {
		// TODO use cull date to delete files
		fmt.Println(file)
	}

	return e
}
