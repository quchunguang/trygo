package main

import (
	"fmt"
	"github.com/Workiva/go-datastructures/queue"
)

func main() {
	q := queue.New(10)
	q.Put(1)
	q.Put("hello")
	g1, _ := q.Get(1)
	g2, _ := q.Get(1)
	fmt.Println(g1[0].(int))
	fmt.Println(g2[0].(string))
}
