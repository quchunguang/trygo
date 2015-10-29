package main

import (
	"flag"
	"fmt"
)

func main() {
	var times int
	flag.IntVar(&times, "n", 0, "Repeat times")
	flag.Parse()
	cmd := flag.Args()

	for i := 0; i < times; i++ {
		fmt.Println(cmd[1])
	}
}
