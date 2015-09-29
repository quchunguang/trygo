package main

import (
	"fmt"
	"github.com/Workiva/go-datastructures/trie/ctrie"
)

func main() {
	t := ctrie.New(nil)
	t.Insert([]byte("A"), 15)
	t.Insert([]byte("to"), 7)
	t.Insert([]byte("tea"), 3)
	t.Insert([]byte("ted"), 4)
	t.Insert([]byte("ten"), 12)
	t.Insert([]byte("i"), 11)
	t.Insert([]byte("in"), 5)
	t.Insert([]byte("inn"), 9)
	val, ok := t.Lookup([]byte("ted"))
	if ok {
		fmt.Println(val.(int))
	}
}
