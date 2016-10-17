package main

import "fmt"

type Mutatable struct {
	a int
	b int
}

func (m Mutatable) StayTheSame() {
	m.a = 5
	m.b = 7
}

func (m *Mutatable) Mutate() {
	m.a = 5
	m.b = 7
}

func main() {
	m := Mutatable{0, 0}
	n := &Mutatable{0, 0}
	fmt.Println(m, n)
	m.StayTheSame()
	n.StayTheSame()
	fmt.Println(m, n)
	m.Mutate()
	n.Mutate()
	fmt.Println(m, n)
}
