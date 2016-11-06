package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"os"
	"strings"
)

func main() {
	var cn163 string
	var ids []string

	flag.StringVar(&cn163, "cn163", "", "Comma separated list of ids.")
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	ids = strings.Split(cn163, ",")

	for _, id := range ids {
		link := "http://cn163.net/archives/" + id
		doc, err := goquery.NewDocument(link)
		if err != nil {
			fmt.Println("Cannot open url")
			os.Exit(2)
		}

		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			href := s.AttrOr("href", "")
			if strings.HasPrefix(href, "ed2k://") {
				uu, _ := url.QueryUnescape(href)
				fmt.Println(uu)
			}
		})
	}
}
