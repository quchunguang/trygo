package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("../test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Md5: %x\n", md5.Sum(data))
	fmt.Printf("Sha1: %x\n", sha1.Sum(data))
	fmt.Printf("Sha256: %x\n", sha256.Sum256(data))
	fmt.Printf("Sha512: %x\n", sha512.Sum512(data))
}
