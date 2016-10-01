package main

import (
	"fmt"
	"github.com/google/btree"
)

func main() {
	// create a 2-3-4 tree (each node contains 1-3 items and 2-4 children)
	tr := btree.New(2)

	for i := btree.Int(100); i > 0; i = i - 2 {
		tr.ReplaceOrInsert(i)
	}

	fmt.Println("len:       ", tr.Len())
	fmt.Println("get3:      ", tr.Get(btree.Int(3)))
	fmt.Println("get100:    ", tr.Get(btree.Int(100)))
	fmt.Println("del4:      ", tr.Delete(btree.Int(4)))
	fmt.Println("del100:    ", tr.Delete(btree.Int(100)))
	fmt.Println("replace5:  ", tr.ReplaceOrInsert(btree.Int(5)))
	fmt.Println("replace100:", tr.ReplaceOrInsert(btree.Int(100)))
	fmt.Println("delmin:    ", tr.DeleteMin())
	fmt.Println("delmax:    ", tr.DeleteMax())
	fmt.Println("len:       ", tr.Len())

	tr.AscendRange(btree.Int(40), btree.Int(60), func(a btree.Item) bool {
		fmt.Print(a, " -> ")
		return true
	})
	fmt.Print("\n\n")
	tr.Ascend(func(a btree.Item) bool {
		fmt.Print(a, " -> ")
		return true
	})
	fmt.Println()
}
