package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

type Person struct {
	Id      int
	Name    string
	Age     int
	Emails  []string
	Company string
	Role    string
}

type OnlineUser struct {
	User      []*Person
	LoginTime string
}

var onlineUser OnlineUser

func HandlerTest(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PostFormValue("test"))
	for _, person := range onlineUser.User {
		if person.Id == id {
			t, err := template.ParseFiles("reply.html")
			checkError(err)

			err = t.Execute(w, &person)
			checkError(err)
		}
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	dumx := Person{
		Id:      1001,
		Name:    "zoro",
		Age:     27,
		Emails:  []string{"dg@gmail.com", "dk@hotmail.com"},
		Company: "Omron",
		Role:    "SE"}

	chxd := Person{Id: 1002, Name: "chxd", Age: 27, Emails: []string{"test@gmail.com", "d@hotmail.com"}}

	onlineUser = OnlineUser{User: []*Person{&dumx, &chxd}}

	t, err := template.ParseFiles("tmpl.html")
	checkError(err)

	err = t.Execute(w, onlineUser)
	checkError(err)
}

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/test", HandlerTest)
	http.ListenAndServe(":8888", nil)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
