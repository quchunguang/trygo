package main

import (
	"fmt"
	"github.com/bndr/gopencils"
)

type respStruct struct {
	Login string
	Id    int
	Name  string
}

func main() {
	api := gopencils.Api("https://api.github.com")
	// Users Resource
	users := api.Res("users")

	usernames := []string{"bndr", "torvalds", "coleifer"}

	for _, username := range usernames {
		// Create a new pointer to response Struct
		r := new(respStruct)
		// Get user with id i into the newly created response struct
		_, err := users.Id(username, r).Get()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(r)
		}
	}
}
