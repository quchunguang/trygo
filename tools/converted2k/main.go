package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
)

func process(in, out string) {
	inFile, _ := os.Open(in)
	defer inFile.Close()
	outFile, _ := os.Create(out)
	defer outFile.Close()

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		uu, _ := url.QueryUnescape(scanner.Text())
		fmt.Println(uu)
		outFile.WriteString(uu + "\n")
	}
}

func main() {
	process("in.txt", "out.txt")
}
