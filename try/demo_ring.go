// 3.9 Linked ring without head
// Each node arrowed to its left node,
//   1->2->3->4->5->6->7->8->9->(1)
//
// Josephus problem
// Reduce one every count to M, till leave one to be the leader
//   Ex., when M=5, we get 5->1->7->4->3->6->9->2->(8), 8 is the leader.
package main

import (
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	item interface{}
	next *Node
}

func main() {
	var last, tmp *Node

	// first one
	last = &Node{1, nil}
	last.next = last

	// build ring
	for i := 2; i < 10; i++ {
		tmp = &Node{i, last.next}
		last.next = tmp
		last = tmp
	}

	// output ring
	for tmp = last.next; tmp != last; tmp = tmp.next {
		fmt.Printf("%d->", tmp.item)
	}
	fmt.Println(tmp.item) // output last one

	// Josephus reduce
	M, _ := strconv.Atoi(os.Args[1])
	for tmp = last; tmp != tmp.next; {
		for i := 1; i < M; i++ {
			tmp = tmp.next
		}
		fmt.Printf("%d->", tmp.next.item)
		tmp.next = tmp.next.next
	}
	fmt.Printf("(%d)\n", tmp.item) // found the leader
}
