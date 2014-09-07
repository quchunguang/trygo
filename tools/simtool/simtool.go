package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

func dnslookup() {
	dnssec := flag.Bool("dnssec", false, "Request DNSSEC records")
	port := flag.String("port", "53", "Set the query port")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [name ...]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if *dnssec {
	}
	if *port == "53" {

	}
}
func psgrp() {
	ps := exec.Command("ps", "-e", "-opid,ppid,comm")
	output, _ := ps.Output()
	child := make(map[int][]int)
	for i, s := range strings.Split(string(output), "\n") {
		if i == 0 || len(s) == 0 {
			continue
		}
		f := strings.Fields(s)
		fp, _ := strconv.Atoi(f[0])
		fpp, _ := strconv.Atoi(f[1])
		child[fpp] = append(child[fpp], fp)
	}
	schild := make([]int, len(child))
	i := 0
	for k, _ := range child {
		schild[i] = k
		i++
	}
	sort.Ints(schild)
	for _, ppid := range schild {
		fmt.Printf("Pid %d has %d child", ppid, len(child[ppid]))
		if len(child[ppid]) == 1 {
			fmt.Printf(": %v\n", child[ppid])
			continue
		}
		fmt.Printf("ren: %v\n", child[ppid])
	}
}

func wc() {
	var chars, words, lines int
	r := bufio.NewReader(os.Stdin)
	for {
		switch s, ok := r.ReadString('\n'); true {
		case ok != nil:
			fmt.Printf("%d %d %d\n", chars, words, lines)
			return
		default:
			chars += len(s)
			words += len(strings.Fields(s))
			lines++
		}
	}
}
func uniq() {
	list := []string{"a", "b", "a", "a", "c", "d", "e", "f"}
	first := list[0]
	fmt.Printf("%s ", first)
	for _, v := range list[1:] {
		if first != v {
			fmt.Printf("%s ", v)
			first = v
		}
	}
}

// Usage
//     $ ./test		    # server side
//     $ nc 127.0.0.1 8053  # client side
//     abc
//     abc
func echoserver() {
	l, err := net.Listen("tcp", "127.0.0.1:8053")
	if err != nil {
		fmt.Printf("Failure to listen: %s\n", err.Error())
		return
	}
	for {
		if c, err := l.Accept(); err == nil {
			go Echo(c)
		}
	}
}
func Echo(c net.Conn) {
	defer c.Close()
	line, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		fmt.Printf("Failure to read: %s\n", err.Error())
		return
	}
	_, err = c.Write([]byte(line))
	if err != nil {
		fmt.Printf("Failure to write: %s\n", err.Error())
		return
	}
}

func main() {
	// dnslookup()
	// psgrp()
	// wc()
	// uniq()
	// echoserver() // true loop by default!
}
