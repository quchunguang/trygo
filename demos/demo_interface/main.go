package main

import (
	"fmt"
)

// TA with a func
type TA struct{}

// HelloA blone to TA
func (a TA) HelloA() {
	fmt.Println("HelloA()")
}

// TB with a func
type TB struct{}

// HelloB belong to TB
func (b TB) HelloB() {
	fmt.Println("HelloB()")
}

// TC combine TA and TB
type TC struct {
	TA
	TB
}

func main() {
	c := &TC{}
	c.HelloA()
	c.HelloB()
}
