package main

import (
	"fmt"
	"os"
	"text/template"
)

var names = struct {
	Id    int
	Names []string
}{
	1,
	[]string{"Kevin", "Rofail"},
}

func main() {
	t := template.Must(template.New("").Parse("<h1>{{.Id}}<h1><table>{{range .Names}}<tr><td>{{.}}</td></tr>{{end}}</table>"))
	// names := []string{"Kevin", "Rofail"}
	if err := t.Execute(os.Stdout, names); err != nil {
		fmt.Println("Err")
	}
}
