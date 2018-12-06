package main

import "testing"
import "fmt"

var tests = []struct {
	in, out string
}{
	{"", ""},
	{"{}", "{}"},
	{"{1,2}", "{1,2}"},
	{"{1,{2,3}}", "{1,{2,3}}"},
}

func TestNew(t *testing.T) {
	for _, tt := range tests {
		var tree = New(tt.in)
		fmt.Println(tree)
	}
}
