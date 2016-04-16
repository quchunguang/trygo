package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
)

const indexPage = `
<h1>Logged in as %s</h1>
<form method="post" action="/logout">
    <button type="submit">Logout</button>
</form>
<a href="/internal">Internal</a>
`

func indexPageHandler(response http.ResponseWriter, request *http.Request) {
	username, nickname, login := getSession(request)
	if username != "" && login > 0 {
		fmt.Fprintf(response, indexPage, nickname)
	} else {
		http.Redirect(response, request, "/login", 302)
	}
}

const internalPage = `
<h1>Internal</h1>
<hr>
<small>Nickname: %s</small>
<form method="post" action="/logout">
    <button type="submit">Logout</button>
</form>
`

func internalPageHandler(response http.ResponseWriter, request *http.Request) {
	username, nickname, login := getSession(request)
	if username != "" && login > 0 {
		fmt.Fprintf(response, internalPage, nickname)
	} else {
		http.Redirect(response, request, "/login", 302)
	}
}

const loginForm = `
<h1>Login</h1>
<form method="post" action="/login">
    <label for="name">User name</label>
    <input type="text" id="name" name="name">
    <label for="password">Password</label>
    <input type="password" id="password" name="password">
    <button type="submit">Login</button>
</form>
`

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Login    int           `json:"login"`    // 0: no right, >0: common right
	Nickname string        `json:"nickname"` // display name
	Username string        `json:"username"` // login name
	Password string        `json:"password"` // login password
}

func connect() (session *mgo.Session) {
	connectURL := "localhost"
	session, err := mgo.Dial(connectURL)
	if err != nil {
		fmt.Println("Can't connect to mongo, go error: ", err.Error())
	}
	session.SetSafe(&mgo.Safe{})
	return session
}

func checkUser(username, password string) *User {
	session := connect()
	defer session.Close()

	var user User
	db_user := session.DB("kzc").C("user")
	cond := bson.M{"username": username, "password": password}
	if err := db_user.Find(cond).One(&user); err != nil {
		return nil
	} else {
		return &user
	}
}

func loginGetHandler(response http.ResponseWriter, request *http.Request) {
	username, _, login := getSession(request)
	if username != "" && login > 0 {
		http.Redirect(response, request, "/", 302)
	} else {
		fmt.Fprintf(response, loginForm)
	}
}

func loginPostHandler(response http.ResponseWriter, request *http.Request) {
	username := request.FormValue("name")
	password := request.FormValue("password")
	redirectTarget := "/login"
	if username != "" && password != "" {
		if user := checkUser(username, password); user != nil {
			setSession(*user, response)
			redirectTarget = "/"
		}
	}
	http.Redirect(response, request, redirectTarget, 302)
}

func logoutHandler(response http.ResponseWriter, request *http.Request) {
	clearSession(response)
	http.Redirect(response, request, "/", 302)
}

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func setSession(user User, response http.ResponseWriter) {
	value := map[string]string{
		"username": user.Username,
		"nickname": user.Nickname,
		"login":    strconv.Itoa(user.Login),
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func getSession(request *http.Request) (username, nickname string, login int) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			if login, err = strconv.Atoi(cookieValue["login"]); err != nil {
				return
			}
			username = cookieValue["username"]
			nickname = cookieValue["nickname"]
		}
	}
	return
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/", indexPageHandler)
	router.HandleFunc("/internal", internalPageHandler)
	router.HandleFunc("/login", loginGetHandler).Methods("GET")
	router.HandleFunc("/login", loginPostHandler).Methods("POST")
	router.HandleFunc("/logout", logoutHandler).Methods("POST")

	http.Handle("/", router)
	http.ListenAndServe(":8000", nil)
}
