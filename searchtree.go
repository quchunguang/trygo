package trygo

import (
	"fmt"
)

type SearchTreeNode struct {
	Chars    []rune
	Children []*SearchTreeNode
	End      bool
}

func NewSearchTree() *SearchTreeNode {
	return new(SearchTreeNode)
}

func (n *SearchTreeNode) hasChar(c rune) (bool, int) {
	for idx, cc := range n.Chars {
		if c == cc {
			return true, idx
		}
	}
	return false, -1
}

func (n *SearchTreeNode) addChar(c rune) *SearchTreeNode {
	n.Chars = append(n.Chars, c)
	nn := new(SearchTreeNode)
	n.Children = append(n.Children, nn)
	return nn
}

func (n *SearchTreeNode) Print() {
	fmt.Print("{")
	if n.End {
		fmt.Print(".")
	}
	for idx, c := range n.Chars {
		fmt.Print(string(c))
		n.Children[idx].Print()
	}
	fmt.Print("}")
}

func (n *SearchTreeNode) Contain(word string) (ret bool) {
	cur := n
	for _, c := range word {
		if ok, idx := cur.hasChar(c); ok {
			cur = cur.Children[idx]
		} else {
			return
		}
	}
	if cur.End {
		ret = true
	}
	return
}

func (n *SearchTreeNode) Add(word string) {
	cur := n
	for _, c := range word {
		if ok, idx := cur.hasChar(c); ok {
			cur = cur.Children[idx]
		} else {
			cur = cur.addChar(c)
		}
	}
	cur.End = true
}
