package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("HIc")
	cc := make(chan bool)
	go func() {
		time.Sleep(1 * time.Second)
		cc <- true
	}()

	select {
	case <-cc:
		fmt.Println(".")
	case <-time.After(3 * time.Second):
		fmt.Println("Done")
	}
}
