package main

import (
	"fmt"
	"github.com/cznic/b"
)

func main() {
	t := b.TreeNew(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	t.Set(3, "qcg")
	t.Set(1, "qcg3")
	t.Set(2, "qcg2")

	en, _ := t.SeekFirst()
	for k, v, e := en.Next(); e == nil; k, v, e = en.Next() {
		fmt.Println(k, " : ", v)
	}
}
