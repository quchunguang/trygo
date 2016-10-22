package main

import (
	"fmt"
	"github.com/quchunguang/trygo"
)

func init() {
	fmt.Println("init(): package main")
}

func init() {
	fmt.Println("init(): package main")
}

func main() {
	fmt.Println("main(): package main")
	fmt.Println(trygo.Pi)
}
