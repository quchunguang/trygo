package main

import (
	"fmt"
	"github.com/petar/GoLLRB/llrb"
)

type lessInt int

func (a lessInt) Less(b llrb.Item) bool { return a < b.(lessInt) }

func printItem(i llrb.Item) bool {
	fmt.Println(i.(lessInt))
	return true
}

func main() {
	tree := llrb.New()
	data := []int{2, 1, 3, 4, 7, 3}
	for _, i := range data {
		tree.ReplaceOrInsert(lessInt(i))
	}
	tree.Delete(lessInt(1))

	tree.AscendLessThan(lessInt(3), printItem)
}
