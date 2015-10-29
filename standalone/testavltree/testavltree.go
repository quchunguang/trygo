package main

import (
	"fmt"
	"github.com/ancientlore/go-avltree"
)

type MyObject struct {
	Key   string
	Value string
}

func (o MyObject) Compare(b avltree.Interface) int {
	if o.Key < b.(MyObject).Key {
		return -1
	}
	if o.Key > b.(MyObject).Key {
		return 1
	}
	return 0
}

func main() {
	// Base AVL tree
	t := avltree.New(func(v1, v2 interface{}) int {
		return v1.(int) - v2.(int)
	}, 0)
	for i := 100; i > 0; i -= 2 {
		t.Add(i)
	}
	fmt.Println("Find 10(exist):     ", t.Find(10))
	fmt.Println("Find 11(not exist): ", t.Find(11))
	fmt.Println("Height:             ", t.Height())

	t.Do(func(v interface{}) {
		fmt.Print(v.(int), " -> ")
	})
	fmt.Print("\n\n")

	// String AVL tree
	st := avltree.NewStringTree(0)
	st.Add("add 1")
	st.Add("add 3")
	st.Add("add 2")
	st.Do(func(v string) {
		fmt.Println(v)
	})
	fmt.Println("Height: ", st.Height())
	fmt.Println()

	// Key(string)-value(interface{}) AVL tree
	pt := avltree.NewPairTree(0)
	pt.Add(avltree.Pair{"b", 5})
	pt.Add(avltree.Pair{"a", 9})
	pt.Add(avltree.Pair{"c", 3})
	pt.Do(func(v avltree.Pair) {
		fmt.Println(v.Key, " -> ", v.Value)
	})

	// Object AVL tree
	ot := avltree.NewObjectTree(0)
	ov, dupe := ot.Add(MyObject{"foo", "bar"})
	if dupe {
		fmt.Println("Key exist:", ov)
	}
	ov, dupe = ot.Add(MyObject{"foo", "tree"})
	if dupe {
		fmt.Println("Key exist:", ov)
	}
	ov, dupe = ot.Add(MyObject{"bar", "bar"})
	if dupe {
		fmt.Println("Key exist:", ov)
	}
	ot.Do(func(v interface{}) {
		fmt.Println(v.(MyObject))
	})
}
