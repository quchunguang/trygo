package main

import (
	"fmt"
	"sort"
)

// Insert v into a sorted set, that is, if already exist, do nothing
func Insert(sortedset *[]int, v int) {
	index := sort.SearchInts(*sortedset, v)
	if index == len(*sortedset) {
		*sortedset = append(*sortedset, v)
		return
	}
	if (*sortedset)[index] != v {
		*sortedset = append((*sortedset)[:index], v)
		*sortedset = append(*sortedset, (*sortedset)[index+1:]...)
	}
}
func main() {
	var data []int
	Insert(&data, 5)
	Insert(&data, 1)
	Insert(&data, 3)
	Insert(&data, 2)
	fmt.Println(data)
}
