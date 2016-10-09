package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("Start...")
	cc := make(chan bool)
	go func() {
		time.Sleep(2 * time.Second)
		cc <- true
	}()

	select {
	case <-cc:
		fmt.Println("Done.")
	case <-time.After(1 * time.Second):
		fmt.Println("Time out.")
	}
}
