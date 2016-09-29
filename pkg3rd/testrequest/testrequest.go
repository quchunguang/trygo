package main

import (
	"fmt"
	"github.com/mozillazg/request"
	"net/http"
)

func main() {
	c := new(http.Client)
	req := request.NewRequest(c)
	resp, _ := req.Get("http://httpbin.org/get")
	j, _ := resp.Json()
	defer resp.Body.Close()
	fmt.Println(j)
	fmt.Println(resp.Body)
}
