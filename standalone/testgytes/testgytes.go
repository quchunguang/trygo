package main

import (
	"fmt"
	"github.com/labstack/gommon/gytes"
)

func main() {
	// B
	b := gytes.Format(515)
	fmt.Println(b)

	// MB
	b = gytes.Format(13231323)
	fmt.Println(b)

	// Exact
	b = gytes.Format(1000 * 1000 * 1000)
	fmt.Println(b)

	// Binary prefix
	gytes.BinaryPrefix(true)
	b = gytes.Format(1323)
	fmt.Println(b)
}
