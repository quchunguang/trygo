package main

import (
	"fmt"
)

func FastSort(li []int, start, end int) {
	p := round(li, start, end)
	if start < p-1 {
		FastSort(li, start, p-1)
	}
	if p+1 < end {
		FastSort(li, p+1, end)
	}
}

func round(li []int, start, end int) int {
	i := start
	j := end - 1
	for {
		for li[i] < li[end] {
			i++
		}
		for li[j] >= li[end] {
			j--
		}
		if i >= j {
			// Inn any condition, switch li[i] with li[end]
			li[i], li[end] = li[end], li[i]
			fmt.Println(i)
			break
		}
		li[i], li[j] = li[j], li[i]
		i++
		j--
	}
	return i
}

func main() {
	li := []int{5, 0, 2, 6, 2, 1, 3, 3}
	FastSort(li, 0, len(li)-1)
	fmt.Println(li)
}
