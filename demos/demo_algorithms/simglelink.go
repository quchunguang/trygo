package main

import (
	"fmt"
)

type Node struct {
	item interface{}
	next *Node
}

func NewNode(item interface{}, next *Node) *Node {
	n := new(Node)
	n.item = item
	n.next = next
	return n
}

func SingleLink() {
	n2 := NewNode(2, nil)
	n1 := NewNode(1, n2)
	n0 := NewNode(0, n1)

	n0.next = n0.next.next // delete node after n0

	for n := n0; n != nil; n = n.next {
		fmt.Println(n.item)
	}
}
