package main

import (
	"fmt"
	"github.com/quchunguang/trygo"
)

// 4.5 Inverse Poland Notation
func PostfixCalc(s string) {
	st := trygo.NewStack()
	for c := 0; c < len(s); c++ {
		if s[c] == '+' {
			st.Push(st.Pop().(int) + st.Pop().(int))
		}
		if s[c] == '*' {
			st.Push(st.Pop().(int) * st.Pop().(int))
		}
		if s[c] >= '0' && s[c] <= '9' {
			st.Push(0)
		}
		for ; s[c] >= '0' && s[c] <= '9'; c++ {
			st.Push(st.Pop().(int)*10 + (int(s[c]) - '0'))
		}
	}
	fmt.Println(st.Pop())
}

// 4.6 Infix Notation
// Notice: Not like calculates, here the "+" and "*" has same priority with
// left->right order!
func Infix2Postfix(s string) (ret string) {
	ops := trygo.NewStack()
	for c := 0; c < len(s); c++ {
		if s[c] == '+' || s[c] == '*' {
			ret += " "
			ops.Push(s[c])
		}
		if s[c] == ')' {
			ret += " " + string(ops.Pop().(byte))
		}
		if s[c] >= '0' && s[c] <= '9' {
			ret += string(s[c])
		}
	}
	for !ops.Empty() {
		ret += " " + string(ops.Pop().(byte))
	}
	fmt.Println(ret)
	return
}

// 5.4 Prefix Notation
func PrefixCalc(s string) {
	i := 0
	var eval func() int
	eval = func() int {
		x := 0
		for s[i] == ' ' {
			i++
		}
		if s[i] == '+' {
			i++
			return eval() + eval()
		}
		if s[i] == '*' {
			i++
			return eval() * eval()
		}
		for s[i] >= '0' && s[i] <= '9' {
			x = 10*x + (int(s[i]) - '0')
			i++
		}
		return x
	}
	fmt.Println(eval())
}
