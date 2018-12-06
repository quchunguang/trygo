package main

import (
	"fmt"
	"testing"
	"time"
)

var calcTests = []struct {
	a, b int
	ret  int
}{
	{1, 2, 3},
	{0, 0, 0},
	{0, 1, 1},
}

func TestCalc(t *testing.T) {
	for _, test := range calcTests {
		ret := Calc(test.a, test.b)
		if ret != test.ret {
			t.Errorf("")
		}
	}
}

func BenchmarkCalc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range calcTests {
			Calc(test.a, test.b)
		}
	}
}

func BenchmarkPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("OK")
	}
}

func BenchmarkSleep(b *testing.B) {
	time.Sleep(time.Second)
}
