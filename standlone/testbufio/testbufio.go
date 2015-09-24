package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	const s = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	scanner := bufio.NewScanner(strings.NewReader(s))
	count := 0
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}
	fmt.Println(count)
}
