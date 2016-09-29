// 3.6 Sieve of Eratosthenes
package main

import (
	"fmt"
	"os"
	"strconv"
)

func GenPrime() {
	N, _ := strconv.Atoi(os.Args[1])
	a := make([]bool, N) // false by default
	fmt.Println(a[10])
	for i := 2; i < N; i++ {
		if !a[i] {
			for j := i; j*i < N; j++ {
				a[j*i] = true
			}
		}
	}
	for i := 2; i < N; i++ {
		if !a[i] {
			fmt.Println(i)
		}
	}
}
