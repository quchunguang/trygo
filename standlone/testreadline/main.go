package main

import (
	"bufio"
	"fmt"
	"os"
)

func try() {
	f, _ := os.Open("test.py")
	defer f.Close()
	r := bufio.NewScanner(f)

	var lines int
	for r.Scan() {
		lines++
		fmt.Printf("%-3d%s\n", lines, r.Text())
	}
	fmt.Println("-----------")
	fmt.Println("Lines: ", lines)
}

func main() {
	var m = make(map[string]int)
	m["qcg"] = 23
	m["ynn"] = 54
	delete(m, "qcg")
	fmt.Println(m)

	var p = []int{2, 3, 4, 5}
	fmt.Println(p)
}
