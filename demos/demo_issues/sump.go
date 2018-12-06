package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func sumP(path string) (sum int) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if !f.IsDir() {
			return nil
		}
		reg := regexp.MustCompile(`[^\d]*(\d+)p.*`)
		match := reg.FindStringSubmatch(path)
		if match == nil {
			return nil
		}
		num, _ := strconv.Atoi(match[1])
		fmt.Println(path, "==>", match[1], "p")
		sum += num
		return nil
	})
	if err != nil {
		fmt.Println("err path")
	}
	fmt.Printf("\nSum = %d p\n", sum)

	return
}

func sumP2(path string) (sum int) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	reg := regexp.MustCompile(`[^\d]*(\d+)p.*`)
	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		match := reg.FindStringSubmatch(f.Name())
		if match == nil {
			continue
		}
		num, _ := strconv.Atoi(match[1])
		sum += num
		fmt.Println(f.Name(), "==>", match[1], "p")
	}
	fmt.Printf("\nSum = %d p\n", sum)
	return
}
