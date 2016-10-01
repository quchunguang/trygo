// 3.9 Linked ring without head
// Each node arrowed to its left node,
//   Ex., when N=9, we get 1->2->3->4->5->6->7->8->9->(1)
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

type RNode struct {
	item interface{}
	next *RNode
}

func SingleRing() {
	var last, tmp *RNode

	// first one
	last = &RNode{1, nil}
	last.next = last

	// build ring
	N, _ := strconv.Atoi(os.Args[1])
	for i := 2; i <= N; i++ {
		tmp = &RNode{i, last.next}
		last.next = tmp
		last = tmp
	}

	// output ring
	for tmp = last.next; tmp != last; tmp = tmp.next {
		fmt.Printf("%d->", tmp.item)
	}
	fmt.Println(tmp.item) // output last one

	// Josephus reduce
	M, _ := strconv.Atoi(os.Args[2])
	for tmp = last; tmp != tmp.next; {
		for i := 1; i < M; i++ {
			tmp = tmp.next
		}
		fmt.Printf("%d->", tmp.next.item)
		tmp.next = tmp.next.next
	}
	fmt.Printf("(%d)\n", tmp.item) // found the leader
}
