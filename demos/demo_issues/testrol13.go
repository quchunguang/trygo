package main

import (
	"bytes"
	"fmt"
)

func testrol13() {
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	msg := "'Twas brillig and the slithy gopher..."
	sec := bytes.Map(rot13, []byte(msg))
	clr := bytes.Map(rot13, []byte(sec))
	fmt.Printf("%s", sec)
	fmt.Printf("%s", clr)
}
