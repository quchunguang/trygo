package main

import (
	"fmt"
	"github.com/rakyll/coop"
	"time"
)

func main() {
	fmt.Println("Start ...")
	done := coop.At(time.Now().Add(5*time.Second), func() {
		fmt.Println("At()")
	})
	<-done // wait for fn to be done

	coop.Every(time.Second, func() {
		fmt.Println("Every()")
	})

	done2 := coop.Timeout(5*time.Second, func() {
		time.Sleep(time.Hour)
		fmt.Println("Wont happen")
	})
	<-done2 // will return false, because timeout occurred
	fmt.Println("Timeout()")

	printFn := func() {
		fmt.Println("All()")
	}
	<-coop.All(printFn, printFn, printFn, printFn)

	printFn2 := func() {
		fmt.Println("AllWithThrottle()")
	}
	<-coop.AllWithThrottle(2, printFn2, printFn2, printFn2, printFn2)

	<-coop.Replicate(5, func() {
		time.Sleep(time.Second)
		fmt.Println("Replicate()")
	})
}
