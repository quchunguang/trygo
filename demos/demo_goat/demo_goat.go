package main

import (
	"net/http"

	"github.com/bahlo/goat"
)

func helloHandler(w http.ResponseWriter, r *http.Request, p goat.Params) {
	goat.WriteJSON(w, map[string]string{
		"hello": p["name"],
	})
}

func main() {
	r := goat.New()
	r.Get("/hello/:name", "hello_url", helloHandler)
	r.Run(":8080")
}
