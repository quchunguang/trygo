package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strconv"
)

const DOUBAN_BASE_URL = "https://movie.douban.com/subject/"

// ExampleScrape(24860563)
func ExampleScrape(id int) {
	reqUrl := DOUBAN_BASE_URL + strconv.Itoa(id) + "/"
	doc, err := goquery.NewDocument(reqUrl)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".recommendations-bd dl").Each(func(i int, s *goquery.Selection) {
		url := s.Find("dt a").AttrOr("href", "")
		band := s.Find("dd a").Text()
		fmt.Printf("%d: %s %s\n", i, band, url)
	})
}
