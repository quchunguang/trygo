package trygo

import (
	"fmt"
	"sync"
)

// create the pipeline.
func genpipe(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// filter pipeline
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// endpoint of pipeline
func outpipe(in <-chan int) int {
	s := 0
	for n := range in {
		s += n
	}
	return s
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// A simpler buffered pipe generator put all data and return immediately.
func genpipe2(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
	return out
}

func DemoPipe() {
	pfrom := genpipe2(3, 4, 5, 6)

	// 1. filter pipe with sq()
	pto := sq(pfrom)
	ret := outpipe(pto)

	// 2. filter can use many times
	// pto := sq(sq(pfrom))
	// ret := outpipe(pto)

	// 3. Distribute the sq work across two goroutines that both read from in.
	// pto1 := sq(pfrom)
	// pto2 := sq(pfrom)
	// pmerge := merge(pto1, pto2)
	// ret := outpipe(pmerge)

	fmt.Println(ret)
}
