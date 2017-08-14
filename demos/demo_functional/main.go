package main

import (
	"fmt"
)

func origin() {
	fmt.Println("func origin()")
}

// Type 1: nested mode
func wapper1(origin func()) func() {
	return func() {
		fmt.Println("wapper1 begin")
		origin()
		fmt.Println("wapper1 end")
	}
}

func wapper2(origin func()) func() {
	return func() {
		fmt.Println("wapper2 begin")
		origin()
		fmt.Println("wapper2 end")
	}
}

func wapper3(origin func()) func() {
	return func() {
		fmt.Println("wapper3 begin")
		origin()
		fmt.Println("wapper3 end")
	}
}

func demoNestedMode() {
	var origin = wapper2(wapper1(origin))
	origin()
	fmt.Println()
}

// Type 2: chan mode
type passable func()

func (p passable) wapper1() passable {
	return func() {
		fmt.Println("wapper1 begin")
		p()
		fmt.Println("wapper1 end")
	}
}

func (p passable) wapper2() passable {
	return func() {
		fmt.Println("wapper2 begin")
		p()
		fmt.Println("wapper2 end")
	}
}

func demoChanMode() {
	var f passable = origin
	f.wapper1().wapper2()()
}

// Type 3: Convert nested mode to chan mode
func (p passable) chanrise(w func(func()) func()) passable {
	return w(p)
}

func demoChanrise() {
	var f passable = origin
	f.chanrise(wapper1).chanrise(wapper2)()
}

// Type 4: Convert nested mode to conbined mode
type wapper func(func()) func()
type conbined []wapper

func (c conbined) run(origin func()) {
	f := origin
	for _, w := range c {
		f = w(f)
	}
	f()
}

func demoConbined() {
	var c = conbined{wapper1, wapper2}
	c = append(c, wapper3)
	c.run(origin)
}

// Mapping
func mapping(f func(int) int, li []int) {
	for i, v := range li {
		li[i] = f(v)
	}
}

func demoMapping() {
	li := []int{1, 2, 3, 4}
	mapping(func(c int) int { return c * 2 }, li)
	fmt.Println(li)
}

// Type assertion
type simplefun passable

func (s simplefun) run() {
	s()
}

func demoType() {
	var f simplefun
	f = func() { fmt.Println("OK") }
	f.run()
}

// Data Interface
type Data interface {
}

// D1D ...
type D1D struct {
	values []float64
}

// MapTo ...
func (d D1D) MapTo(f func(float64) float64) {
	for i, v := range d.values {
		d.values[i] = f(v)
	}
}

// Reduce ...
func (d D1D) Reduce(f func(float64, float64) float64, initv float64) (ret float64) {
	ret = initv
	for _, v := range d.values {
		ret = f(ret, v)
	}
	return
}

func (d D1D) String() (ret string) {
	ret += "["
	for _, v := range d.values {
		ret += fmt.Sprint(v, " ")
	}
	ret += "]"
	return
}

// NewD1D ...
func NewD1D(values ...float64) (ret *D1D) {
	ret = new(D1D)
	ret.values = values
	return
}

func demoInterface() {
	d := NewD1D(1, 2, 3, 4, 5)
	d.MapTo(func(x float64) float64 { return x * 2 })
	fmt.Println(d)
	retAdd := d.Reduce(func(x, y float64) float64 { return x + y }, 0)
	fmt.Println(retAdd)
	retMul := d.Reduce(func(x, y float64) float64 { return x * y }, 1)
	fmt.Println(retMul)
}

// demoSlice ...
func demoSlice() {
	s1 := []int{1, 2, 3}
	fmt.Println(len(s1), cap(s1), s1) //prints 3 3 [1 2 3]

	s2 := s1[1:]
	fmt.Println(len(s2), cap(s2), s2) //prints 2 2 [2 3]

	for i := range s2 {
		s2[i] += 20
	}

	//still referencing the same array
	fmt.Println(s1) //prints [1 22 23]
	fmt.Println(s2) //prints [22 23]

	s2 = append(s2, 4)

	for i := range s2 {
		s2[i] += 10
	}

	//s1 is now "stale"
	fmt.Println(s1) //prints [1 22 23]
	fmt.Println(s2) //prints [32 33 14]
}

// Try fields
func main() {
	// demoNestedMode()
	// demoChanMode()
	// demoChanrise()
	// demoConbined()
	// demoMapping()
	// demoType()
	// demoInterface()
	demoSlice()
}
