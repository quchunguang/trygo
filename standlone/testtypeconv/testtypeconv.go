package main

/*
#include <stdio.h>

*/
import "C"
import (
	"fmt"
	"github.com/quchunguang/trygo"
	// "time"
)

type empty interface{}

func double(d empty) empty {
	switch d.(type) {
	case int:
		return d.(int) * 2
	case string:
		return d.(string) + d.(string)
	}
	return d
}
func map3(e []empty, f func(empty) empty) empty {
	m := make([]empty, len(e))
	for k, v := range e {
		m[k] = f(v)
	}
	return m
}

func call(i int64) {
	fmt.Println(i)
}
func main() {
	C.puts(C.CString("Hello, 世界\n"))
	fmt.Println("hi")
	trygo.DemoChannel5()
}
