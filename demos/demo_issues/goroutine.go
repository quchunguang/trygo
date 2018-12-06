package main

import (
	"fmt"
	"time"
)

// WRONG!!! For val variable in the above loops is actually a single variable.
// func wrong1(values []int) {
// 	for val := range values {
// 		// go val.my_method()
// 		go func() {
// 			fmt.Print(val)
// 		}()
// 	}
// }

// RIGHT!!! The val is evaluated at each iteration and placed on the stack for the goroutine.
func right1(values []int) {
	for val := range values {
		go func(val interface{}) {
			fmt.Print(val)
		}(val)
	}
}

// RIGHT!!! The variables declared within the body of a loop are not shared between iterations, and thus can be used separately in a closure.
func right2(values []int) {
	for i := range values {
		val := values[i]
		go func() {
			fmt.Print(val)
		}()
	}
}

// RIGHT!!! Without executing this closure as a goroutine, the code runs as expected
func right3(values []int) {
	for i := 0; i < 10; i++ {
		func() {
			fmt.Print(i)
		}()
	}
}
func goroutine() {
	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// wrong1(values)
	// right1(values)
	// right2(values)
	right3(values)
	time.Sleep(time.Second)
}
