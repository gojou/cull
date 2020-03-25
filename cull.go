package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {

	root := "."
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
			// path = "\t" + path
			if !info.ModTime().After(time.Date(2017, time.January, 1, 0, 0, 0, 0, time.Local)) {
				fmt.Println(path + " " + info.ModTime().Format(time.RFC1123))
				return os.Remove(path)

			}

		}

		return err

	})

	return e
}
