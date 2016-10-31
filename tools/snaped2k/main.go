package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"os"
	"strings"
)

func main() {
	doc, err := goquery.NewDocument("http://cn163.net/archives/3083/")
	if err != nil {
		fmt.Println("Cannot open url")
		os.Exit(1)
	}

	f, _ := os.Create("out.txt")
	defer f.Close()
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href := s.AttrOr("href", "")
		if strings.HasPrefix(href, "ed2k://") {
			url, _ := url.QueryUnescape(href)
			f.WriteString(url + "\n")
		}
	})
}
