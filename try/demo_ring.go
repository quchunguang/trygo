// Linked ring (with head node) and Josephus problem
// Each node arrowed to its left node,
//   9->8->7->6->5->4->3->2->1->head->(9)
package main

import (
	"fmt"
)

type Node struct {
	item interface{}
	next *Node
}

func main() {
	var head, tmp *Node

	// head node not use !!!
	head = new(Node)
	head.item = -1
	head.next = head

	// build ring
	for i := 1; i < 10; i++ {
		tmp = new(Node)
		tmp.item = i
		tmp.next = head.next
		head.next = tmp
	}

	// output ring
	for tmp := head.next; tmp != head; tmp = tmp.next {
		fmt.Print(tmp.item)
		fmt.Print("->")
	}
	fmt.Println()

	// Josephus reduce
	// remove head
	var first *Node
	for first = head; first.next != head; first = first.next {
	}
	first.next = head.next

	for tmp := first.next; tmp != first; tmp = tmp.next {
		fmt.Print(tmp.item)
		fmt.Print("->")
	}
	fmt.Println()
}
