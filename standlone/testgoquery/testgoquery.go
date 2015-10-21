package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
)

func ExampleScrape() {
	doc, err := goquery.NewDocument("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".reviews-wrap article .review-rhs").Each(func(i int, s *goquery.Selection) {
		band := s.Find("h3").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})
}

func main() {
	ExampleScrape()
}
