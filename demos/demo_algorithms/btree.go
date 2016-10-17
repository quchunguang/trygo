package main

import (
	"fmt"
	"github.com/quchunguang/queue"
	"github.com/quchunguang/stack"
)

type BTNode struct {
	Item interface{}
	L, R *BTNode
}

// 5.19
func GenMax(s string) *BTNode {
	a := []byte(s)
	var max func(int, int) *BTNode
	max = func(l, r int) *BTNode {
		m := (l + r) / 2
		x := new(BTNode)
		x.Item = a[m]
		if l == r {
			return x
		}
		x.L = max(l, m)
		x.R = max(m+1, r)
		if x.L.Item.(byte) > x.R.Item.(byte) {
			x.Item = x.L.Item
		} else {
			x.Item = x.R.Item
		}
		return x
	}
	return max(0, len(a)-1)
}

// 5.14
func BTTraverseFront(root *BTNode) {
	if root == nil {
		return
	}
	fmt.Printf("%c", root.Item.(byte)) // visit
	BTTraverseFront(root.L)
	BTTraverseFront(root.R)
}

func BTTraverseMiddle(root *BTNode) {
	if root == nil {
		return
	}
	BTTraverseMiddle(root.L)
	fmt.Printf("%c", root.Item.(byte)) // visit
	BTTraverseMiddle(root.R)
}

func BTTraverseLast(root *BTNode) {
	if root == nil {
		return
	}
	BTTraverseLast(root.L)
	BTTraverseLast(root.R)
	fmt.Printf("%c", root.Item.(byte)) // visit
}

// 5.15
func BTTraverseFrontNR(root *BTNode) {
	s := stack.New()
	s.Push(root)
	for !s.Empty() {
		x := s.Pop().(*BTNode)
		fmt.Printf("%c", x.Item.(byte)) // visit
		if x.R != nil {
			s.Push(x.R)
		}
		if x.L != nil {
			s.Push(x.L)
		}
	}
}

// 5.16
func BTTraverseLevel(root *BTNode) {
	q := queue.New()
	q.Put(root)

	for !q.Empty() {
		x := q.Get().(*BTNode)
		fmt.Printf("%c", x.Item.(byte)) // visit
		if x.L != nil {
			q.Put(x.L)
		}
		if x.R != nil {
			q.Put(x.R)
		}
	}
}
