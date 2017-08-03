package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Process command line flags
	port := flag.String("p", "8300", "port to serve on")
	directory := flag.String("d", "www", "the directory of static file to host")
	flag.Parse()

	// Handle HTTP requests
	http.Handle("/", http.FileServer(http.Dir(*directory)))

	// Serve
	log.Printf("Serving on: %s\n", *directory)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
