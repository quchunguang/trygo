package main

import (
	"fmt"
)

const N int = 10

func FastSearch() {
	var a [N]int
	var p, q int

	for i := 0; i < N; i++ {
		a[i] = i
	}

	// 0 1 2 3 4 5 6 7 8 9 init out
	// 0 2 2 3 4 5 6 7 8 9 1-2  1-2
	// 0 3 3 3 4 5 6 7 8 9 1-3  1-3
	// 0 3 3 3 4 5 6 7 8 9 2-3
	for {
		fmt.Scanf("%d %d\n", &p, &q)
		t := a[p]
		if t == a[q] {
			continue
		}
		for i := 0; i < N; i++ {
			if a[i] == t {
				a[i] = a[q]
			}
		}
		fmt.Println(p, "-", q)
	}

}

func FastMerge() {
	var a [N]int
	var p, q int

	for i := 0; i < N; i++ {
		a[i] = -1 // Not connect to anyone
	}

	for {
		fmt.Scanf("%d %d\n", &p, &q)
		t := a[p]
		for t != -1 {
			if a[t] == q {

			}
		}
	}
}
