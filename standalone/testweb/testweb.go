package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Login    float64       `json:"login"`
	Nickname string        `json:"nickname"`
	Password string        `json:"password"`
	Username string        `json:"username"`
}

func connect() (session *mgo.Session) {
	connectURL := "localhost"
	session, err := mgo.Dial(connectURL)
	if err != nil {
		log.Fatal("Can't connect to mongo, go error: ", err.Error())
	}
	session.SetSafe(&mgo.Safe{})
	return session
}

func getUsers() (users []User) {
	session := connect()
	defer session.Close()

	iter := session.DB("kzc").C("user").Find(nil).Iter()
	// iter := session.DB("kzc").C("user").Find(bson.M{"username": "admin"}).Iter()
	var user User
	for iter.Next(&user) {
		users = append(users, user)
	}
	return users
}

func Index(w http.ResponseWriter, r *http.Request) {
	var locals = make(map[string]interface{})
	locals["users"] = getUsers()

	if r.Method == "GET" {
		t, err := template.ParseFiles("template/index.html")
		if err != nil {
			return
		}
		err = t.Execute(w, locals)
	}
}

func Favicon(w http.ResponseWriter, r *http.Request) {}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func StaticServer(prefix string, staticDir string) {
	http.Handle(prefix, http.StripPrefix(prefix, http.FileServer(http.Dir(staticDir))))
	return
}

func main() {
	StaticServer("/static/", "./static")
	http.HandleFunc("/", Index)
	http.HandleFunc("/favicon.ico", Favicon)
	err := http.ListenAndServe(":8000", Log(http.DefaultServeMux))
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
