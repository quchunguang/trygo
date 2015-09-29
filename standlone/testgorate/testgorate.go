package main

import (
	"fmt"
	"time"

	"github.com/beefsack/go-rate"
)

func main() {
	rl := rate.New(3, time.Second) // 3 times per second
	begin := time.Now()
	for i := 1; i <= 10; i++ {
		rl.Wait()
		fmt.Printf("%d started at %s\n", i, time.Now().Sub(begin))
	}
	// Output:
	// 1 started at 12.584us
	// 2 started at 40.13us
	// 3 started at 44.92us
	// 4 started at 1.000125362s
	// 5 started at 1.000143066s
	// 6 started at 1.000144707s
	// 7 started at 2.000224641s
	// 8 started at 2.000240751s
	// 9 started at 2.00024244s
	// 10 started at 3.000314332s
}
