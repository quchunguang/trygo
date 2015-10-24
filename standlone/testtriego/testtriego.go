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

	var nodes int = 0
	countTrieNodes(trie, &nodes)
	fmt.Printf("Number of allocated nodes: %d\n", nodes)
}
