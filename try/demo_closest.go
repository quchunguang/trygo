// 3.8 Close point pairs count with O(n^2)
package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Point struct {
	x, y float64
}

func main() {
	rand.Seed(time.Now().UnixNano())
	N, _ := strconv.Atoi(os.Args[1])
	d, _ := strconv.ParseFloat(os.Args[2], 64)
	P := make([]Point, N)

	for i := 0; i < N; i++ {
		P[i].x = rand.Float64()
		P[i].y = rand.Float64()
	}

	cnt := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			// if distance between Pi and Pj is less than d
			if math.Hypot(P[i].x-P[j].x, P[i].y-P[j].y) < d {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}
