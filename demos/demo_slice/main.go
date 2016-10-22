package main

import (
	"fmt"
)

func main() {
	var s []int
	fmt.Println(s)
	fmt.Println(s == nil)
	// s[0] = 1 // panic:

	s = append(s, 1) // ok

	ss := make([]int, 10)
	ss[0] = 1 // ok
}
