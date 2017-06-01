package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abbot/go-http-auth"
)

func Secret(user, realm string) string {
	users := map[string]string{
		"john": "$1$dlPL2MqE$oQmn16q49SqdmhenQuNgs1", //hello
	}

	if a, ok := users[user]; ok {
		return a
	}
	return ""
}

func doRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>static file server</h1><p><a href='./static'>folder</p>")
}

func handleFileServer(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("static"))
	http.StripPrefix("/static/", fs)
}

func main() {

	authenticator := auth.NewBasicAuthenticator("localhost", Secret)

	// how to secure the FileServer with basic authentication??
	// fs := http.FileServer(http.Dir("static"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/static/", auth.JustCheck(authenticator, handleFileServer))

	http.HandleFunc("/", auth.JustCheck(authenticator, doRoot))

	log.Println(`Listening... http://localhost:3000
 folder is ./static
 authentication in map users`)
	http.ListenAndServe(":3001", nil)
}
