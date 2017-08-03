package main

import "fmt"

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

// Try fields
func main() {
	demoNestedMode()
	demoChanMode()
}
