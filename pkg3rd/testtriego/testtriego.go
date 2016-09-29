package main

import (
	"fmt"
	"github.com/alediaferia/triego"
)

func countTrieNodes(trie *triego.Trie, i *int) {
	if len(trie.Children) == 0 {
		*i = *i + 1
		return
	}
	for _, v := range trie.Children {
		countTrieNodes(v, i)
	}

	*i = *i + 1
}

func main() {
	trie := triego.NewTrie()

	trie.AppendWord("trial")
	trie.AppendWord("trie")

	fmt.Println("Stored words:")
	for _, w := range trie.Words() {
		fmt.Print(w, " -> ")
	}
	fmt.Print("\n\n")

	// Output: 4
	// That is, 3 actual nodes plus the root node.
	// This is because 1 node is required for the "tri" prefix and just 2 additional nodes for "e" and "al":
	var nodes int = 0
	countTrieNodes(trie, &nodes)
	fmt.Printf("Number of allocated nodes: %d\n", nodes)
}
