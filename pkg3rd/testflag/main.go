package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Usage = func() {
		fmt.Println("............................")
		flag.PrintDefaults()
		fmt.Println("............................")
	}

	var times int
	flag.IntVar(&times, "n", 0, "Repeat times")
	flag.Parse()
	cmd := flag.Args()

	for i := 0; i < times; i++ {
		fmt.Println(cmd[1])
	}
}
