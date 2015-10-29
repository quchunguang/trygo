package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str := "c29tZSBkYXRhIHdpdGggACBhbmQg77u/"
	data, _ := base64.StdEncoding.DecodeString(str)
	fmt.Printf("%q", data)
}
