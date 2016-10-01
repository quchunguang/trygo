package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

func main() {
	input := []byte(`# The Title Level 1
This is a simple **common style** [markdown](https://daringfireball.net/projects/markdown/) text.

* item1 - *OKOK*
* item2 - [Google](http://google.com)
* item3 - [Baidu](http://baidu.com)
`)
	// Generate HTML from Markdown
	html := blackfriday.MarkdownCommon(input)

	// Using UGCPolicy to filter HTML Generated
	html = bluemonday.UGCPolicy().SanitizeBytes(html)
	fmt.Println(string(html))

	// Query HTML using jquery style goquery
	buf := bytes.NewBuffer(html)
	query, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		panic(err)
	}
	query.Find("li a").Each(func(_ int, items *goquery.Selection) {
		url, exists := items.Attr("href")
		if exists {
			fmt.Println(url)
		}
	})
}
