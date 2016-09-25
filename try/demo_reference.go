import (
    "fmt"
    "log"
)

type Item struct {
    data int
}
type T struct {
    name  string
    score []Item
}

func change(t *T) {
    var i Item
    i.data = 1
    t.score = append(t.score, i)
    i.data = 2
    t.score = append(t.score, i)
}

func DemoReference() {
    var t = new(T)
    change(t)
    fmt.Println(t.score)
}
