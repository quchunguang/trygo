package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	p := fmt.Println
	p(strings.Compare("a", "ab"))
	p(strings.Contains("foobar", "foo"))
	p(strings.ContainsAny("foobar", "aei"))
	p(strings.EqualFold("Go", "go"))
	p(strings.Fields("  foo bar  baz   "))
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	p(strings.FieldsFunc("  foo1;bar2,baz3...", f))

}
