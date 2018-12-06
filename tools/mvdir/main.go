package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`(.*)\.S\d\dE\d\d(.*)`)

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Println(err)
	}

	pre := ""
	for _, file := range files {
		f := file.Name()
		if stat, err := os.Stat(f); err != nil || stat.IsDir() {
			continue
		}
		d := re.ReplaceAllString(f, "$1")
		if d != pre {
			pre = d
			newpath := filepath.Join(".", d)
			os.Mkdir(newpath, os.ModePerm)
		}
		os.Rename(filepath.Join(".", f), filepath.Join(".", d, f))
	}
}
