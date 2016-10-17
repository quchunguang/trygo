package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		for i := 1; i < 5; i++ {
			c <- i
			time.Sleep(time.Second)
		}
	}()

	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-time.After(3 * time.Second):
			return
		}
		fmt.Println("round...")
	}
}
