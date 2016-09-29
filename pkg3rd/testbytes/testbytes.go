package main

import (
	"fmt"
	"github.com/labstack/gommon/bytes"
)

func main() {
	// B
	b := bytes.Format(515)
	fmt.Println(b)

	// MB
	b = bytes.Format(13231323)
	fmt.Println(b)

	// Exact
	b = bytes.Format(1000 * 1000 * 1000)
	fmt.Println(b)

	b = bytes.Format(1323)
	fmt.Println(b)
}
