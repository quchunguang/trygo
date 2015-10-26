package main

import (
	"fmt"
	"github.com/jbrukh/bayesian"
)

const (
	Good bayesian.Class = "Good"
	Bad  bayesian.Class = "Bad"
)

func main() {
	classifier := bayesian.NewClassifier(Good, Bad)
	goodStuff := []string{"tall", "rich", "handsome"}
	badStuff := []string{"poor", "smelly", "ugly"}
	classifier.Learn(goodStuff, Good)
	classifier.Learn(badStuff, Bad)

	scores, likely, _ := classifier.LogScores([]string{"tall", "girl"})
	fmt.Println(scores, likely)

	probs, likely, _ := classifier.ProbScores([]string{"tall", "girl"})
	fmt.Println(probs, likely)
}
