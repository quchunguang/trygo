package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(10)
	for i := 1; i <= r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}

	fmt.Println("Move(5): ", r.Move(5).Value)
	r = r.Unlink(9)
	fmt.Print("Unlink(9): ")
	r.Do(func(p interface{}) { fmt.Print(" -> ", p.(int)) })
	fmt.Println(" (before -> 1)")
	fmt.Println(r.Len())

	rtmp := ring.New(1)
	rtmp.Value = 100
	r.Link(rtmp)
	fmt.Print("Link(rtmp): ")
	r.Do(func(p interface{}) { fmt.Print(" -> ", p.(int)) })
	fmt.Println(" (before -> 2, rtmp->100)")

}
