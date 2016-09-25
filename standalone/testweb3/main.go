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
var sessions = map[string]bool{}

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

	sessions[sid] = true

	loginCookie := http.Cookie{
		Name:     "id",
		Value:    sid,
		MaxAge:   int((time.Hour * 12).Seconds()),
		HttpOnly: true,
		Domain:   "mydomain.com",
		Path:     "/admin/",
	}

	return loginCookie, nil
}

func randString(size int) (string, error) {
	buf := make([]byte, size)

	if _, err := rand.Read(buf); err != nil {
		log.Println(err)
		return "", errors.New("Couldn't generate random string")
	}

	return base64.URLEncoding.EncodeToString(buf)[:size], nil
}

func sessionExists(r *http.Request) bool {
	cookie, err := r.Cookie("id")
	if err == http.ErrNoCookie {
		return false
	} else if err != nil {
		log.Println(err)
		return false
	}

	if _, exists := sessions[cookie.Value]; !exists {
		return false
	}

	return true
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	if !sessionExists(r) {
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
	}
	fmt.Fprintf(w, "Hello astaxie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		if sessionExists(r) {
			http.Redirect(w, r, "/", http.StatusOK)
		} else {
			t, _ := template.ParseFiles("login.gtpl")
			t.Execute(w, nil)
		}
	} else {
		cookie, err := tryLogin(r.Form["username"][0], r.Form["password"][0])
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
		}
		http.SetCookie(w, &cookie)
		fmt.Println("username:", r.Form["username"][0])
		fmt.Println("password:", r.Form["password"][0])
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
