// 5.7
package main

import (
	"fmt"
)

//
// d==1: move one step to the right
// d==-1: move two step to the right (equal to move one step to the left)
func Hanoi(n int, d int) {
	if n == 0 {
		return
	}
	Hanoi(n-1, -d)
	fmt.Printf("shift(%d, %d)\n", n, d)
	Hanoi(n-1, -d)
}
