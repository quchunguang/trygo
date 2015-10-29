package main

import (
	"fmt"
	"github.com/awalterschulze/gographviz"
)

func main() {
	graphAst, _ := gographviz.Parse([]byte(`digraph G {}`))
	graph := gographviz.NewGraph()
	gographviz.Analyse(graphAst, graph)
	graph.AddNode("G", "a", nil)
	graph.AddNode("G", "b", nil)
	graph.AddEdge("a", "b", true, nil)
	fmt.Println(graph.String())

	g := gographviz.NewGraph()
	g.SetName("G")
	g.SetDir(true)
	g.AddNode("G", "Hello", nil)
	g.AddNode("G", "World", nil)
	g.AddEdge("Hello", "World", true, nil)
	fmt.Println(g.String())
}
