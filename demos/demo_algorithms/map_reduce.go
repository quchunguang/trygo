package main

import (
	"fmt"
)

type Thing struct {
	Data int
	List []int
}

func (t Thing) Times(n int) Thing {
	t.Data *= n
	return t
}

func (t Thing) Append(n ...int) Thing {
	t.List = append(t.List, n...)
	return t
}

func (t Thing) Delete(n ...int) Thing {
	var rest []int
	found := false
	for _, v := range t.List {
		for _, dv := range n {
			if v == dv {
				found = true
				break
			}
		}
		if !found {
			rest = append(rest, v)
		} else {
			found = false
		}
	}
	t.List = rest
	return t
}

type FuncMap func(int) int

func (t Thing) Map(f FuncMap) Thing {
	for i, v := range t.List {
		t.List[i] = f(v)
	}
	return t
}

type FuncReduce func(int, int) int

func (t Thing) Reduce(f FuncReduce, from int) Thing {
	for _, v := range t.List {
		from = f(from, v)
	}
	if t.List != nil {
		t.List[0] = from
		t.List = t.List[:1]
	}
	return t
}

type FuncFilter func(int) bool

func (t Thing) Filter(f FuncFilter) Thing {
	// The best way to remove items from slice in for loop is just rebuild one!
	var rest []int
	for _, v := range t.List {
		if f(v) {
			rest = append(rest, v)
		}
	}
	t.List = rest
	return t
}

func MapReduce() {
	t := Thing{2, nil}
	t = t.Times(2).Times(3)
	t = t.Append(1).Append(2).Append(3, 4)
	t = t.Delete(4, 5)
	fmt.Println(t.List)
	t = t.Filter(func(i int) bool { return i%2 == 0 })
	t = t.Map(func(i int) int { return 2 * i })
	t = t.Reduce(func(i, j int) int { return i + j }, 0)
	fmt.Println(t)
}
