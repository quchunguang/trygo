package main

import (
	"fmt"
	"github.com/carlescere/goback"
	"time"
)

func main() {
	b := &goback.SimpleBackoff{
		Min:    100 * time.Millisecond,
		Max:    60 * time.Second,
		Factor: 2,
	}
	goback.Wait(b)               // sleeps 100ms
	goback.Wait(b)               // sleeps 200ms
	goback.Wait(b)               // sleeps 400ms
	fmt.Println(b.NextAttempt()) // prints 800ms
	b.Reset()                    // resets the backoff
	goback.Wait(b)               // sleeps 100ms
}
