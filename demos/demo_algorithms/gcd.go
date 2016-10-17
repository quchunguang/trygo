// 5.3
package main

func Gcd(m, n int) int {
	if n == 0 {
		return m
	}
	return Gcd(n, m%n)
}
