package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	mapenv := make(map[string]string)
	p := fmt.Println
	p(os.Getpid())
	p(os.Getpagesize())
	envs := os.Environ()
	for _, k := range envs {
		slice := strings.SplitN(k, "=", 2)
		mapenv[slice[0]] = slice[1]
	}
	p(os.LookupEnv("PATH"))
}
