package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	a := []int64{2, 3, 4, 5, 7}
	b := a[1:3]
	b[0] = -1
	fmt.Println(a)

	c := [...]int64{2, 3, 4, 5, 7}
	// d := make([]int64, len(c))
	// fmt.Println(copy(d, c[:]))
	d := c[:]
	d[0] = -1

	fmt.Println(c)

	var e []int
	fmt.Println("init", cap(e))
	for i := 0; i < 3; i++ {
		e = append(e, i)
		fmt.Println(i, cap(e))
	}

	fmt.Println(string(FindDigits3("testdata/find_digits.txt")))
}

var digitRegexp = regexp.MustCompile("[0-9]+")

// This will cause golang holding full file data not to release
func FindDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	return digitRegexp.Find(b)
}

func FindDigits2(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

func FindDigits3(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	return append([]byte{}, digitRegexp.Find(b)...)
}
