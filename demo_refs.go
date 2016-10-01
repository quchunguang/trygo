package trygo

import (
    "fmt"
)

type Score struct {
    data int
}
type T struct {
    name  string
    scores []Score
}

func change(t *T) {
    var i Score
    i.data = 1
    t.scores = append(t.scores, i)
    i.data = 2
    t.scores = append(t.scores, i)
}

func DemoRefs() {
    var t = new(T)
    change(t)
    fmt.Println(t.scores)
}
