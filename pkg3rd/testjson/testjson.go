package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Message struct {
	Name string
	Text string
}

const jsonStream = `
{"Name": "Ed", "Text": "Knock knock."}
{"Name": "Sam", "Text": "Who's there?"}
{"Name": "Ed", "Text": "Go fmt."}
{"Name": "Sam", "Text": "Go fmt who?"}
{"Name": "Ed", "Text": "Go fmt yourself!"}
`

func main() {
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else {
			fmt.Println(m)
		}
	}
}
