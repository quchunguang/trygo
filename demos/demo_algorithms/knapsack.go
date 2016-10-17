// 5.13 Dynamic programming algorithm - Knapsack problem
package main

import (
	"fmt"
)

type item struct {
	size int
	val  int
}

// Maximize sum of value with limit of caps
func Knapsack(caps int) int {
	const unknown = -1
	var knap func(int) int

	// size-value pairs of given objects
	items := []item{{3, 4}, {4, 5}, {7, 10}, {8, 11}, {9, 13}}

	// itemKnown[M] is the best plan just at cap M
	itemKnown := make([]*item, caps+1, caps+1)

	// maxKnown[M] is the max value of cap M
	maxKnown := make([]int, caps+1, caps+1)
	for i := 0; i < caps+1; i++ {
		maxKnown[i] = unknown
	}

	knap = func(M int) int {
		var max, maxi int
		if maxKnown[M] != unknown {
			return maxKnown[M]
		}
		for i := 0; i < len(items); i++ {
			if space := M - items[i].size; space >= 0 {
				if t := knap(space) + items[i].val; t > max {
					max = t
					maxi = i
				}
			}
		}
		maxKnown[M] = max
		itemKnown[M] = &items[maxi]
		return max
	}
	ret := knap(caps) // solve

	// get result
	for M := caps; M > 0; M -= itemKnown[M].size {
		fmt.Println(*itemKnown[M])
	}
	return ret
}
