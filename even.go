/*
   The even package implements a simple library for
   even test.
*/
package testgo

// Even tests if gaven i is even. Return true for even.
func Even(i int) bool {
	return i%2 == 0
}

// private function starts with lower charactor
func odd(i int) bool {
	return i%2 == 1
}
