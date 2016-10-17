package main

import (
	"encoding/json"
	"net/http"
)

type ZombieThing struct {
	Text string `json:"text"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func zombie(w http.ResponseWriter, r *http.Request) {
	zomb := ZombieThing{"Watch out for this guy!", "Bob Zombie", 12}
	b, err := json.Marshal(zomb)
	if err != nil {
		panic(err)
	}
	w.Write(b)
}

func main() {
	http.HandleFunc("/zombie/", zombie)
	http.ListenAndServe(":9000", nil)
}
