package main

import (
	"fmt"
	"github.com/smartystreets/mafsa"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Create a BuildTree and insert your items in lexicographical order
	bt := mafsa.New()
	bt.Insert("cities") // an error will be returned if input is < last input
	bt.Insert("city")
	bt.Insert("pities")
	bt.Insert("pity")
	bt.Finish()

	err := bt.Save("filename")
	checkErr(err)
	mt, err := mafsa.Load("filename")
	checkErr(err)

	fmt.Println("Does tree contain 'cities'?", mt.Contains("cities"))
	fmt.Println("Does tree contain 'pitiful'?", mt.Contains("pitiful"))

	// You can traverse down to a certain node, if it exists
	fmt.Printf("'y' node is at: %p\n", mt.Traverse([]rune("city")))

	// To traverse the tree and get the number of elements inserted
	// before the prefix specified
	node, idx := mt.IndexedTraverse([]rune("pit"))
	if node != nil {
		fmt.Printf("Index number for 'pit': %d\n", idx)
	}

	myData := []string{
		"The plural of city",
		"Noun; a large town",
		"The state of pitying",
		"A feeling of sorrow and compassion",
	}
	node, idx = mt.IndexedTraverse([]rune("pities"))
	if node != nil && node.Final {
		fmt.Println(myData[idx-1])
	}
}
