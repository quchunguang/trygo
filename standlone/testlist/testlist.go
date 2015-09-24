package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack("cc")
	l.PushBack("dd")
	l.PushBack("bb")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("-> %s\n", e.Value.(string))
	}
}
