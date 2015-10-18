package main

import (
	"fmt"
	"github.com/deckarep/golang-set"
)

func main() {
	requiredClasses := mapset.NewSet()
	requiredClasses.Add("Cooking")
	requiredClasses.Add("English")
	requiredClasses.Add("Math")
	requiredClasses.Add("Biology")

	scienceSlice := []interface{}{"Biology", "Chemistry"}
	scienceClasses := mapset.NewSetFromSlice(scienceSlice)

	electiveClasses := mapset.NewSet()
	electiveClasses.Add("Welding")
	electiveClasses.Add("Music")
	electiveClasses.Add("Automotive")

	bonusClasses := mapset.NewSet()
	bonusClasses.Add("Go Programming")
	bonusClasses.Add("Python Programming")

	// Show me all the available classes I can take
	allClasses := requiredClasses.Union(scienceClasses).Union(electiveClasses).Union(bonusClasses)
	fmt.Println(allClasses)
	// Set{Cooking, English, Math, Chemistry, Welding, Biology, Music, Automotive, Go Programming, Python Programming}

	// Is cooking considered a science class?
	fmt.Println(scienceClasses.Contains("Cooking"))
	// false

	// Show me all classes that are not science classes, since I hate science.
	fmt.Println(allClasses.Difference(scienceClasses))
	// Set{Music, Automotive, Go Programming, Python Programming, Cooking, English, Math, Welding}

	// Which science classes are also required classes?
	fmt.Println(scienceClasses.Intersect(requiredClasses))
	// Set{Biology}

	// How many bonus classes do you offer?
	fmt.Println(bonusClasses.Cardinality())
	// 2

	// Do you have the following classes? Welding, Automotive and English?
	fmt.Println(allClasses.IsSuperset(mapset.NewSetFromSlice([]interface{}{"Welding", "Automotive", "English"}))) //true
}
