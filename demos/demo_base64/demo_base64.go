package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	src := []byte("some data with \x00 and \ufeff")
	des := base64.StdEncoding.EncodeToString(src)
	fmt.Println(des)
	src2, _ := base64.StdEncoding.DecodeString(des)
	fmt.Printf("%q\n", src2)
}
