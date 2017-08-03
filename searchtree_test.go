package trygo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var words = []string{"cat", "call", "be", "bee", "zoo", "beg"}
var nowords = []string{"caty", "cal", "bebe", "b", "oo", "bg"}

func TestAdd(t *testing.T) {
	tree := NewSearchTree()
	for _, word := range words {
		tree.Add(word)
	}

	for _, word := range words {
		assert.True(t, tree.Contain(word))
	}
	for _, word := range nowords {
		assert.False(t, tree.Contain(word))
	}
}

func TestPrint(t *testing.T) {
	tree := NewSearchTree()
	for _, word := range words {
		tree.Add(word)
	}
	tree.Print()
	fmt.Println()
}
