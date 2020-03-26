package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {

	root := "."
	culldate := time.Date(2017, time.January, 1, 0, 0, 0, 0, time.Local)
	fmt.Printf("Delete all files in current directory last modified before %v? (n/Y) ", culldate.Format("January 2, 2006"))
	confirmed := askconfirm()
	if confirmed {
		err := cull(root, culldate)
		if err != nil {
			panic(err)
		}
	}

}

func cull(root string, culldate time.Time) (e error) {
	// TODO define cull date

	e = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// TODO if the file is a directory, recursively call cull(file)
		if !info.IsDir() {
			// path = "\t" + path
			if !info.ModTime().After(culldate) {
				fmt.Println(path + " " + info.ModTime().Format(time.RFC1123))
				return os.Remove(path)

			}

		}

		return err

	})

	return e
}

func askconfirm() bool {

	var s string
	_, e := fmt.Scan(&s)
	if e != nil {
		panic(e)
	}
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "y" || s == "yes" {
		return true
	}

	return false
}
