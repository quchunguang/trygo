package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/axgle/mahonia"
)

func main() {
	var argIds, argURL, argSfx, argFile string
	var argGBK bool
	flag.StringVar(&argIds, "ids", "", "Comma separated list of ids, like `11,31,51`. Each joined after the given `--url`.")
	flag.StringVar(&argURL, "url", "cu163.com/", "URL of the site without `http://`. `cu163.com/` by default.")
	flag.StringVar(&argSfx, "sfx", ".html", "Subfix of the url. `.html` by default.")
	flag.StringVar(&argFile, "o", "list.txt", "The output file. `list.txt` by default.")
	flag.BoolVar(&argGBK, "gbk", false, "If output as GBK encoding.")

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	f, _ := os.Create(argFile)
	defer f.Close()

	re := regexp.MustCompile(`ed2k:\/\/\|file\|[^\<\>\|\\\/]+\|\d+\|[0-9a-zA-Z]{32}\|h=[a-z0-9]+\|\/`)
	enc := mahonia.NewEncoder("GBK")
	ids := strings.Split(argIds, ",")
	for _, id := range ids {
		link := "http://" + argURL + id + argSfx
		// doc, err := goquery.NewDocument(link)
		resp, err := http.Get(link)
		if err != nil {
			fmt.Println("Cannot open url")
			os.Exit(2)
		}
		defer resp.Body.Close()
		ret, err := ioutil.ReadAll(resp.Body)

		res := re.FindAllString(string(ret), -1)
		tmp := ""
		for _, u := range res {
			if u == tmp {
				continue
			}
			tmp = u
			u, _ := url.QueryUnescape(u)
			if argGBK {
				u = enc.ConvertString(u)
			}
			fmt.Fprintln(f, u)
		}

		// doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// 	href := s.AttrOr("href", "")
		// 	if strings.HasPrefix(href, "ed2k://") {
		// 		uu, _ := url.QueryUnescape(href)
		// 		fmt.Println(uu)
		// 	}
		// })
	}
}
