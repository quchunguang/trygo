// 2.1 Search()
// 2.2 Search2()
package main

import (
	"fmt"
)

var a = []int{0, 1, 2, 3, 4, 7, 9}

func Search(a []int, v, l, r int) int {
	for i := l; i < r; i++ {
		if a[i] == v {
			return i
		}
	}
	return -1
}

func Search2(a []int, v, l, r int) int {
	for l <= r {
		m := (l + r) / 2
		if v == a[m] {
			return m
		}
		if v < a[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}

func Searching() {
	i := Search2(a, 7, 0, N)
	fmt.Println(i)
}
