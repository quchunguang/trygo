package main

import (
	"fmt"
	"github.com/endeveit/guesslanguage"
)

func main() {
	lang, _ := guesslanguage.Guess("这是一段中文。")
	fmt.Println(lang)
	// Output:
	// en
}
