package main

import (
	"fmt"
	"io"
)

func main() {
	r, w := io.Pipe()
	go w.Write([]byte("Hello World!"))
	buf := make([]byte, 20)
	n, err := r.Read(buf)
	fmt.Println(string(buf), n, err)
}
