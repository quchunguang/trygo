package main

import (
	"fmt"
	"gopkg.in/fatih/set.v0"
	"strconv"
	"sync"
)

var p = fmt.Println

func main() {
	// create a set with zero items
	s := set.New()
	// s := set.NewNonTS() // non thread-safe version

	// ... or with some initial values
	// s := set.New("istanbul", "frankfurt", 30.123, "san francisco", 1234)
	// s := set.NewNonTS("kenya", "ethiopia", "sumatra")

	// add items
	s.Add("istanbul")
	s.Add("istanbul") // nothing happens if you add duplicate item

	// add multiple items
	s.Add("ankara", "san francisco", 3.14)

	// remove item
	s.Remove("frankfurt")
	s.Remove("frankfurt") // nothing happes if you remove a nonexisting item

	// remove multiple items
	s.Remove("barcelona", 3.14, "ankara")

	// removes an arbitary item and return it
	item := s.Pop()
	p(item)

	// create a new copy
	other := s.Copy()
	p(other)

	// remove all items
	other.Clear()

	// number of items in the set
	len := s.Size()
	p(len)

	// return a list of items
	items := s.List()
	p(items)

	// string representation of set
	p(s.String())

	// check for set emptiness, returns true if set is empty
	p(s.IsEmpty())

	// check for a single item exist
	p(s.Has("istanbul"))

	// ... or for multiple items. This will return true if all of the items exist.
	p(s.Has("istanbul", "san francisco", 3.14))

	// create two sets for the following checks...
	u := set.New("1", "2", "3", "4", "5")
	v := set.New("1", "2", "3")

	// check if they are the same
	if !u.IsEqual(v) {
		fmt.Println("u is not equal to v")
	}

	// if s contains all elements of t
	if u.IsSubset(v) {
		fmt.Println("u is a subset of v")
	}

	// ... or if s is a superset of t
	if v.IsSuperset(u) {
		fmt.Println("v is a superset of u")
	}

	// let us initialize two sets with some values
	a := set.New("ankara", "berlin", "san francisco")
	b := set.New("frankfurt", "berlin")

	// creates a new set with the items in a and b combined.
	// [frankfurt, berlin, ankara, san francisco]
	p(set.Union(a, b))

	// contains items which is in both a and b
	// [berlin]
	p(set.Intersection(a, b))

	// contains items which are in a but not in b
	// [ankara, san francisco]
	p(set.Difference(a, b))

	// contains items which are in one of either, but not in both.
	// [frankfurt, ankara, san francisco]
	p(set.SymmetricDifference(a, b))

	// like Union but saves the result back into a.
	a.Merge(b)

	// removes the set items which are in b from a and saves the result back into a.
	a.Separate(b)

	mulSetOpt()
	helperOpt()
	threadSafe()
}

// Multiple Set Operations
func mulSetOpt() {
	a := set.New("1", "3", "4", "5")
	b := set.New("2", "3", "4", "5")
	c := set.New("4", "5", "6", "7")

	// creates a new set with items in a, b and c
	// [1 2 3 4 5 6 7]
	p(set.Union(a, b, c))

	// creates a new set with items in a but not in b and c
	// [1]
	p(set.Difference(a, b, c))

	// creates a new set with items that are common to a, b and c
	// [5]
	p(set.Intersection(a, b, c))
}

//
func helperOpt() {
	// create a set of mixed types
	s := set.New("ankara", "5", "8", "san francisco", 13, 21)

	// convert s into a slice of strings (type is []string)
	// [ankara 5 8 san francisco]
	p(set.StringSlice(s))

	// u contains a slice of ints (type is []int)
	// [13, 21]
	p(set.IntSlice(s))
}

func threadSafe() {
	var wg sync.WaitGroup // this is just for waiting until all goroutines finish

	// Initialize our thread safe Set
	s := set.New()

	// Add items concurrently (item1, item2, and so on)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			item := "item" + strconv.Itoa(i)
			p("adding", item)
			s.Add(item)
			wg.Done()
		}(i)
	}

	// Wait until all concurrent calls finished and print our set
	wg.Wait()
	p(s)
}
