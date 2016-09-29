package main

import (
	"fmt"
	"github.com/willf/bitset"
)

func main() {
	var b bitset.BitSet
	b.Set(10).Set(11)
	if b.Test(1000) {
		b.Clear(1000)
	}
	for i, e := b.NextSet(0); e; i, e = b.NextSet(i + 1) {
		fmt.Println("The following bit is set:", i)
	}
	if b.Intersection(bitset.New(100).Set(10).Set(11)).Count() > 1 {
		fmt.Println("Intersection works.")
	}
}
