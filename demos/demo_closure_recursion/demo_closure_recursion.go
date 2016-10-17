// Closure recursion
package main

func Problem(n int) (ret int) {
	size := n        // works like static variable inside inner()
	var inner func() // must defines clearly with var

	inner = func() { // closure function
		if size == 0 {
			return // return condition
		}
		ret += size // change result
		size--      // static
		inner()     // recursion
	}
	inner() // start inner()
	return  // return ret
}

func main() {
	Problem(4) // Output: 10
}
