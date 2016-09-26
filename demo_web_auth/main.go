package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var users = map[string]string{"qcg": "qq", "ynn": "nn"}
var sessions = map[string]time.Time{}
var EXPIRES = time.Hour * 12

func tryLogin(username, password string) (http.Cookie, error) {
	// if exists := db.UserExists(username, password); !exists {
	if pass, exists := users[username]; !exists || pass != password {
		return http.Cookie{},
			errors.New("The username or password you entered isn't correct.")
	}

	sid, err := randString(32)
	if err != nil {
		return http.Cookie{}, err
	}

	sessions[sid] = time.Now().Add(EXPIRES) // set session with expires

	loginCookie := http.Cookie{
		Name:     "id",
		Value:    sid,
		MaxAge:   int(EXPIRES.Seconds()),
		HttpOnly: true,
		Domain:   "localhost",
		Path:     "/",
	}

	return loginCookie, nil
}

func randString(size int) (string, error) {
	buf := make([]byte, size)

	if _, err := rand.Read(buf); err != nil {
		return "", errors.New("Couldn't generate random string")
	}

	return base64.URLEncoding.EncodeToString(buf)[:size], nil
}

func sessionExists(r *http.Request) bool {
	cookie, err := r.Cookie("id")
	if err != nil {
		log.Println(err)
		return false
	}
	expires, exists := sessions[cookie.Value]
	if !exists {
		return false
	}
	if expires.Before(time.Now()) {
		delete(sessions, cookie.Value) // delete session when expires
		return false
	}

	return true
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	if !sessionExists(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	fmt.Fprintf(w, "Authorized!")
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if sessionExists(r) {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			t, _ := template.ParseFiles("login.gtpl")
			t.Execute(w, nil)
		}
	} else {
		r.ParseForm()
		cookie, err := tryLogin(r.Form["username"][0], r.Form["password"][0])
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
