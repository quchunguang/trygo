package main

type (
	Tree *Node
	Node struct {
		name        string
		left, right *Node
	}
)

func New(s string) Tree {
	if s == "" {
		return nil
	}

	var tree, cur *Node
	var first = true
	for _, c := range s {
		from:
		if first {
			// do something
			delete from:to
		}
		to:
		if c == '{' {
			cur = cur
		}
	}
}
