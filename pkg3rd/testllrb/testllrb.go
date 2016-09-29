package main

import (
	"fmt"
	"github.com/petar/GoLLRB/llrb"
)

func main() {
	rb := llrb.New()
	for i := llrb.Int(100); i > 0; i -= 2 {
		rb.InsertNoReplace(i)
	}
	rb.AscendGreaterOrEqual(llrb.Int(0), func(i llrb.Item) bool {
		fmt.Print(i, " -> ")
		return true
	})
	fmt.Println()
	fmt.Println(rb.Root().Item)
}
