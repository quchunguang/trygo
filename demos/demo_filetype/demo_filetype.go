package main

import (
	"fmt"
	"gopkg.in/h2non/filetype.v0"
	"io/ioutil"
)

func main() {
	buf, _ := ioutil.ReadFile("sample.jpg")

	// Simple file type checking
	kind, unkwown := filetype.Match(buf)
	if unkwown != nil {
		fmt.Printf("Unkwown: %s", unkwown)
		return
	}
	fmt.Printf("File type: %s. MIME: %s\n",
		kind.Extension, kind.MIME.Value)

	// Check type class
	if filetype.IsImage(buf) {
		fmt.Println("File is an image")
	} else {
		fmt.Println("Not an image")
	}

	// Check if file is supported by extension
	if filetype.IsSupported("jpg") {
		fmt.Println("Extension supported")
	} else {
		fmt.Println("Extension not supported")
	}

	// Check if file is supported by extension
	if filetype.IsMIMESupported("image/jpeg") {
		fmt.Println("MIME type supported")
	} else {
		fmt.Println("MIME type not supported")
	}

	// File header
	// We only have to pass the file header = first 261 bytes
	head := buf[:261]
	if filetype.IsImage(head) {
		fmt.Println("File is an image")
	} else {
		fmt.Println("Not an image")
	}

	// Add additional file type matchers
	fooType := filetype.NewType("foo", "foo/foo")

	// Register the new matcher and its type
	filetype.AddMatcher(fooType, func(buf []byte) bool {
		return len(buf) > 1 && buf[0] == 0x01 && buf[1] == 0x02
	})

	// Check if the new type is supported by extension
	if filetype.IsSupported("foo") {
		fmt.Println("New supported type: foo")
	}

	// Check if the new type is supported by MIME
	if filetype.IsMIMESupported("foo/foo") {
		fmt.Println("New supported MIME type: foo/foo")
	}

	// Try to match the file
	fooFile := []byte{0x01, 0x02}
	newkind, _ := filetype.Match(fooFile)
	if newkind == filetype.Unknown {
		fmt.Println("Unknown file type")
	} else {
		fmt.Printf("File type matched: %s\n", newkind.Extension)
	}
}
