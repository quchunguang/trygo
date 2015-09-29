package main

import (
	"fmt"
	"github.com/ryszard/goskiplist/skiplist"
)

func main() {
	s := skiplist.NewIntMap()
	s.Set(7, "seven")
	s.Set(1, "one")
	s.Set(0, "zero")
	s.Set(5, "five")
	s.Set(9, "nine")
	s.Set(10, "ten")
	s.Set(3, "three")

	firstValue, ok := s.Get(0)
	if ok {
		fmt.Println(firstValue)
	}
	// prints:
	//  zero

	s.Delete(7)

	secondValue, ok := s.Get(7)
	if ok {
		fmt.Println(secondValue)
	}
	// prints: nothing.

	s.Set(9, "niner")

	// Iterate through all the elements, in order.
	unboundIterator := s.Iterator()
	for unboundIterator.Next() {
		fmt.Printf("%d: %s\n", unboundIterator.Key(), unboundIterator.Value())
	}
	// prints:
	//  0: zero
	//  1: one
	//  3: three
	//  5: five
	//  9: niner
	//  10: ten

	for unboundIterator.Previous() {
		fmt.Printf("%d: %s\n", unboundIterator.Key(), unboundIterator.Value())
	}
	//  9: niner
	//  5: five
	//  3: three
	//  1: one
	//  0: zero

	boundIterator := s.Range(3, 10)
	// Iterate only through elements in some range.
	for boundIterator.Next() {
		fmt.Printf("%d: %s\n", boundIterator.Key(), boundIterator.Value())
	}
	// prints:
	//  3: three
	//  5: five
	//  9: niner

	for boundIterator.Previous() {
		fmt.Printf("%d: %s\n", boundIterator.Key(), boundIterator.Value())
	}
	// prints:
	//  5: five
	//  3: three

	var iterator skiplist.Iterator

	iterator = s.Seek(3)
	fmt.Printf("%d: %s\n", iterator.Key(), iterator.Value())
	// prints:
	//  3: three

	iterator = s.Seek(2)
	fmt.Printf("%d: %s\n", iterator.Key(), iterator.Value())
	// prints:
	//  3: three

	iterator = s.SeekToFirst()
	fmt.Printf("%d: %s\n", iterator.Key(), iterator.Value())
	// prints:
	//  0: zero

	iterator = s.SeekToLast()
	fmt.Printf("%d: %s\n", iterator.Key(), iterator.Value())
	// prints:
	//  10: ten

	// SkipList can also reduce subsequent forward seeking costs by reusing the
	// same iterator:

	iterator = s.Seek(3)
	fmt.Printf("%d: %s\n", iterator.Key(), iterator.Value())
	// prints:
	//  3: three

	iterator.Seek(5)
	fmt.Printf("%d: %s\n", iterator.Key(), iterator.Value())
	// prints:
	//  5: five
}
