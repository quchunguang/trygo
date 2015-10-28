package main

import (
	"fmt"
	"github.com/saschpe/tribool"
)

func main() {
	a, b, c := tribool.True, tribool.Indeterminate, tribool.False
	fmt.Println(a.Not(), a.And(b), a.Or(c),
		tribool.FromBool(true), tribool.FromString("indeterminate"),
		tribool.Indeterminate.String())

	if !a.True() || a.Indeterminate() || a.False() {
		fmt.Println("err")
	}

	if a.Equal(b) {
		fmt.Println("err")
	}
}
