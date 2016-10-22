package main

import (
	"fmt"
	"net/http"
)

func access(name string) {
	r, _ := http.Get("https://github.com/" + name)
	defer r.Body.Close()
	if r.StatusCode != 200 {
		fmt.Println(name, r.StatusCode)
	}

}
func main() {
	// var i, j byte
	// for i = 'a'; i <= 'z'; i++ {
	// 	fmt.Println(string(i))
	// 	for j = 'a'; j <= 'z'; j++ {
	// 		access(string(i) + string(j))
	// 	}
	// }
	access("dsadadsfdsfd")
}
