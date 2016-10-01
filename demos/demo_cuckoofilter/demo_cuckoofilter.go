package main

import (
	"fmt"
	"github.com/seiflotfy/cuckoofilter"
)

func main() {
	cf := cuckoofilter.NewDefaultCuckooFilter()
	cf.InsertUnique([]byte("geeky ogre"))

	// Lookup a string (and it a miss) if it exists in the cuckoofilter
	cf.Lookup([]byte("hello"))

	fmt.Println("count == ", cf.Count())
	// count == 1

	// Delete a string (and it a miss)
	cf.Delete([]byte("hello"))

	fmt.Println("count == ", cf.Count())
	// count == 1

	// Delete a string (a hit)
	cf.Delete([]byte("geeky ogre"))

	fmt.Println("count == ", cf.Count())
	// count == 0
}
