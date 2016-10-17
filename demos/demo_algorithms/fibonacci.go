// 5.11 Dynamic programming algorithm - Fibonacci
package main

// int64 range: math.MinInt64 through math.MaxInt64.
// F(92) = 7540113804746346429 < 9223372036854775807 = 0x7FFFFFFFFFFFFFFF
// F(93) will overflow
var knownF [93]int64

func F(n int64) int64 {
	if knownF[n] != 0 {
		return knownF[n]
	}
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	knownF[n] = F(n-1) + F(n-2)
	return knownF[n]
}
