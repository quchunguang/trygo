// 5.6
package main

func Max(a []int, l int, r int) int {
	if l == r {
		return a[l]
	}
	m := (l + r) / 2
	u := Max(a, l, m)
	v := Max(a, m+1, r)
	if u > v {
		return u
	} else {
		return v
	}
}
