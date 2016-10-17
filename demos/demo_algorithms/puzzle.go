// 5.2
package main

func Puzzle(n int) int {
	if n == 1 {
		return 1
	}
	if n%2 == 0 {
		return Puzzle(n / 2)
	} else {
		return Puzzle(3*n + 1)
	}
}
