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

	e = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil

	})

	for _, file := range files {
		fmt.Println(file)
	}

	return e
}
