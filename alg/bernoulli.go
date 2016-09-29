// 3.7 Bernoulli test
// A test throw corn N times, see how many times gets the front side.
// M round tests at total.
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func Bernoulli() {
	rand.Seed(time.Now().UnixNano())
	N, _ := strconv.Atoi(os.Args[1])
	M, _ := strconv.Atoi(os.Args[2])
	f := make([]int, N+1)

	for i := 0; i < M; i++ {
		cnt := 0
		for j := 0; j < N; j++ {
			if rand.Float32() < 0.5 {
				cnt++
			}
		}
		f[cnt]++
	}

	// draw graph
	for j := 0; j <= N; j++ {
		if f[j] == 0 {
			fmt.Println(".")
			continue
		}
		for i := 0; i < f[j]; i += 10 {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
