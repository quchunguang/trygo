package trygo

import (
	"fmt"
)

// Score struct
type Score struct {
	data int
}

// T struct
type T struct {
	name   string
	scores []Score
}

func change(t *T) {
	var i Score
	i.data = 1
	t.scores = append(t.scores, i)
	i.data = 2
	t.scores = append(t.scores, i)
}

// DemoRefs func
func DemoRefs() {
	var t = new(T)
	change(t)
	fmt.Println(t.scores)
}
