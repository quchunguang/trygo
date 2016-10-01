package main

import (
	"fmt"
)

func RecursiveSubsets(li []int, b []bool, start int, end int) {
	if start <= end {
		b[start] = true
		RecursiveSubsets(li, b, start+1, end)
		b[start] = false
		RecursiveSubsets(li, b, start+1, end)
	} else {
		for i := 0; i <= end; i++ {
			if b[i] {
				fmt.Print(li[i])
			}
		}
		fmt.Println("*")
	}
}

// Recursive method
func PrintAllSubsets1(li []int) {
	b := make([]bool, len(li))
	RecursiveSubsets(li, b, 0, len(li)-1)
}

// Binary add method
func PrintAllSubsets2(li []int) {
	b := make([]bool, len(li))

	for {
		// Output one
		for i := 0; i < len(li); i++ {
			if b[i] {
				fmt.Print(li[i])
			}
		}
		fmt.Println("*")

		k := len(li) - 1
		for k >= 0 {
			if !b[k] {
				b[k] = true
				break
			} else {
				b[k] = false
				k--
			}
		}
		if k < 0 {
			break
		}
	}
}

func Permutation() {
	li := []int{1, 2, 3, 4}
	// PrintAllSubsets1(li)
	PrintAllSubsets2(li)
}
