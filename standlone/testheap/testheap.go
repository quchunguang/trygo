package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	lenh := len(*h)
	x := (*h)[lenh-1]
	*h = (*h)[0 : lenh-1]
	return x
}
func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("Smallest item is %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("Pop item is %d\n", heap.Pop(h))
	}
}
