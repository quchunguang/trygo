package main

// $ curl -s https://api.douban.com/v2/book/1220562 > douban_book.json
// $ go get github.com/ChimeraCoder/gojson/gojson
// $ curl -s https://api.douban.com/v2/book/1220562 | gojson -name=Payload

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Payload struct {
	Alt         string   `json:"alt"`
	AltTitle    string   `json:"alt_title"`
	Author      []string `json:"author"`
	AuthorIntro string   `json:"author_intro"`
	Binding     string   `json:"binding"`
	Catalog     string   `json:"catalog"`
	ID          string   `json:"id"`
	Image       string   `json:"image"`
	Images      struct {
		Large  string `json:"large"`
		Medium string `json:"medium"`
		Small  string `json:"small"`
	} `json:"images"`
	Isbn10      string `json:"isbn10"`
	Isbn13      string `json:"isbn13"`
	OriginTitle string `json:"origin_title"`
	Pages       string `json:"pages"`
	Price       string `json:"price"`
	Pubdate     string `json:"pubdate"`
	Publisher   string `json:"publisher"`
	Rating      struct {
		Average   string `json:"average"`
		Max       int    `json:"max"`
		Min       int    `json:"min"`
		NumRaters int    `json:"numRaters"`
	} `json:"rating"`
	Subtitle string `json:"subtitle"`
	Summary  string `json:"summary"`
	Tags     []struct {
		Count int    `json:"count"`
		Name  string `json:"name"`
		Title string `json:"title"`
	} `json:"tags"`
	Title      string   `json:"title"`
	Translator []string `json:"translator"`
	URL        string   `json:"url"`
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func WriteJson(filename string, p Payload) {
	f, err := os.Create(filename)
	checkErr(err)
	defer f.Close()
	res, err := json.MarshalIndent(p, "", "    ")
	checkErr(err)
	f.Write(res)
}

func ReadJson(filename string) (ret Payload) {
	f, err := os.Open(filename)
	defer f.Close()
	checkErr(err)
	b, err := ioutil.ReadAll(f)
	checkErr(err)
	json.Unmarshal(b, &ret)
	return
}

func main() {
	book := ReadJson("douban_book.json")
	WriteJson("douban_book.json", book)
	fmt.Printf("Title: %s\nAuthor: %s\n",
		book.Title,
		book.Author[0])
}
