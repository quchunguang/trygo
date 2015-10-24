package main

import (
	"bytes"
	"fmt"
	"github.com/steveyen/gtreap"
	"math/rand"
)

func stringCompare(a, b interface{}) int {
	return bytes.Compare([]byte(a.(string)), []byte(b.(string)))
}

func main() {
	t := gtreap.NewTreap(stringCompare)
	t = t.Upsert("hi", rand.Int())
	t = t.Upsert("hola", rand.Int())
	t = t.Upsert("bye", rand.Int())
	t = t.Upsert("adios", rand.Int())

	fmt.Println(t.Get("hi"))
	fmt.Println(t.Get("bye"))

	// Some example Delete()'s...
	t = t.Delete("bye")
	fmt.Println(t.Get("bye"))
	t2 := t.Delete("hi")
	fmt.Println(t2.Get("hi"))

	// Since we still hold onto treap t, we can still access "hi".
	fmt.Println(t.Get("hi"))

	t.VisitAscend("a", func(i gtreap.Item) bool {
		// This visitor callback will be invoked with every item
		// from "cya" onwards.  So: "hi", "hola".
		// If we want to stop visiting, return false;
		// otherwise a true return result means keep visiting items.
		fmt.Print(i, " -> ")
		return true
	})
	fmt.Println()
}
