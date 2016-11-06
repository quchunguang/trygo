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
	var argIds, argURL string
	var ids []string

	flag.StringVar(&argIds, "ids", "", "Comma separated list of ids, like `11,31,51`. Each joined after the given `--url`.")
	flag.StringVar(&argURL, "url", "cn163.net/archives/", "URL of the site without `http://`. `cn163.net/archives/` by default.")
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	ids = strings.Split(argIds, ",")

	for _, id := range ids {
		link := "http://" + argURL + id
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
