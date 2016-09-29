package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dir_list, err := ioutil.ReadDir("d:/repos/winpi")
	if err != nil {
		panic(err)
	}
	for i, v := range dir_list {
		fmt.Println(i, "=", v.Name(), v.ModTime())
	}
}
