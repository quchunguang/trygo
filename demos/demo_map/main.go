// https://blog.golang.org/go-maps-in-action
package main

import (
	"fmt"
	"sort"
	"sync"
)

func Common() {
	var m map[string]int
	fmt.Println(m)
	fmt.Println(m == nil)
	// m["aa"] = 1  // panic:

	mm := make(map[string]int)
	mm["aa"] = 1 // ok

	fmt.Println(mm["bb"])
	fmt.Println(len(mm))

	if v, ok := mm["bb"]; ok {
		fmt.Println(v)
	}

	commits := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}
	fmt.Println(commits)
}

func ZeroBool() {
	type Node struct {
		Next  *Node
		Value interface{}
	}
	var first *Node

	visited := make(map[*Node]bool)
	for n := first; n != nil; n = n.Next {
		if visited[n] { // zero value of bool is (false, bool), so no need 2-value form
			fmt.Println("cycle detected")
			break
		}
		visited[n] = true
		fmt.Println(n.Value)
	}
}

func ZeroSlice() {
	type Person struct {
		Name  string
		Likes []string
	}
	var people []*Person

	likes := make(map[string][]*Person)
	for _, p := range people {
		for _, l := range p.Likes {
			likes[l] = append(likes[l], p) // zero value of slice is (nil, Type), so no need 2-value form
		}
	}

	for _, p := range likes["cheese"] {
		fmt.Println(p.Name, "likes cheese.")
	}
	fmt.Println(len(likes["bacon"]), "people like bacon.")
}

func MapMap() {
	// slices, maps, and functions cannot be compared, so cannot be key

	// Each key of the outer map is the path to a web page with its own inner map. Each inner map key is a two-letter country code. This expression retrieves the number of times an Australian has loaded the documentation page:
	hits := make(map[string]map[string]int)

	// Unfortunately, this approach becomes unwieldy when adding data, as for any given outer key you must check if the inner map exists, and create it if needed:
	add := func(m map[string]map[string]int, path, country string) {
		mm, ok := m[path]
		if !ok {
			mm = make(map[string]int)
			m[path] = mm
		}
		mm[country]++
	}
	add(hits, "/doc/", "au")

	n := hits["/doc/"]["au"]
	fmt.Println(n)
}

// same example as MapMap
func StructMap() {
	type Key struct {
		Path, Country string
	}
	hits := make(map[Key]int)
	hits[Key{"/ref/spec", "ch"}]++
	n := hits[Key{"/ref/spec", "ch"}]
	fmt.Println(n)
}

// ConcurrencyMap using sync.RWMutex for maps are not safe for concurrent use.
func ConcurrencyMap() {
	var wg sync.WaitGroup

	var counter = struct {
		sync.RWMutex
		m map[string]int
	}{m: make(map[string]int)}

	wg.Add(1)
	go func() {
		counter.Lock()
		counter.m["some_key"]++
		counter.Unlock()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		counter.RLock()
		n := counter.m["some_key"]
		counter.RUnlock()
		fmt.Println("some_key:", n)
		wg.Done()
	}()

	wg.Wait()
}

// OrderMap using a slice keep the order of items in map for map do not keep iterator order.
func OrderMap() {
	var m map[int]string
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}
}

func main() {
	Common()
	ZeroBool()
	ZeroBool()
	MapMap()
	StructMap()
	ConcurrencyMap()
	OrderMap()
}
