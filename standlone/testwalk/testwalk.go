package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walkFn(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println("Walker Error: ", err)
		return nil
	}
	if info.IsDir() {
		fmt.Println("Directory: ", path)
	} else {
		fmt.Println("File: ", path)
	}
	return nil
}

func main() {
	root := "./"
	filepath.Walk(root, walkFn)
}
