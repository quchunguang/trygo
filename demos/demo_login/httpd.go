package main

import (
	"encoding/json"
	"flag"
	"github.com/gorilla/sessions"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	// "time"
)

var db *mgo.Session
var users *mgo.Collection

var store = sessions.NewCookieStore(
	[]byte("5DFC2BC7F9E8499387DFCD172C8D98FF"),
	[]byte("839F293D10994635AB76B69A8AC90DBA"))

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	// Password2 string `json:"password2"`
	Nickname string `json:"nickname"`
}

type Resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func checkUser(username, password string) (user User, ok bool) {
	err := users.Find(bson.M{"username": username}).One(&user)
	if err == nil && user.Password == password {
		ok = true
		return
	}
	ok = false
	return
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		options := sessions.Options{
			Path:   "/",
			MaxAge: 0,
		}
		session, _ := store.Get(r, "session")
		session.Options = &options

		if user, ok := checkUser(r.PostFormValue("username"), r.PostFormValue("password")); ok {
			log.Println("Login -> ", user.Username)

			session.Values["username"] = user.Username
			session.Values["token"] = uuid.NewV4().String()
			session.Save(r, w)

			// expiration := time.Now().Add(30 * 24 * time.Hour)
			// cookie := &http.Cookie{
			// 	Name:  "nickname",
			// 	Value: url.QueryEscape(user.Nickname),
			// 	Expires: expiration,
			// }
			cookie := sessions.NewCookie("nickname", url.QueryEscape(user.Nickname), &options)
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
	http.Redirect(w, r, "login.html", http.StatusSeeOther)
}

func logout(w http.ResponseWriter, r *http.Request) {
	var cookie http.Cookie
	cookie = http.Cookie{Name: "session", Path: "/", MaxAge: -1}
	http.SetCookie(w, &cookie)
	cookie = http.Cookie{Name: "nickname", Path: "/", MaxAge: -1}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func signup(w http.ResponseWriter, r *http.Request) {
	var err error
	var user User
	var resp Resp

	if r.Method == "POST" {
		r.ParseForm()
		res, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()

		// NOTICE: Not need every segment exist (Password2) exist in User when Unmarshal()
		err = json.Unmarshal(res, &user)
		checkErr(err)

		n, _ := users.Find(bson.M{"username": user.Username}).Count()
		if n > 0 { // Exist already
			resp.Code = 1
			resp.Message = "该用户已存在"
		} else {
			users.Insert(user)
			resp.Code = 0
			resp.Message = "注册成功"
		}
		payload, _ := json.Marshal(resp)
		w.Write(payload)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Initialize MongoDB
	// Indexing with MongoDB client, db.user.createIndex({"username": 1})
	var err error
	db, err = mgo.Dial("localhost")
	checkErr(err)
	defer db.Close()
	users = db.DB("login").C("user")
	checkErr(err)

	// Process command line flags
	port := flag.String("p", "8200", "port to serve on")
	directory := flag.String("d", "www", "the directory of static file to host")
	flag.Parse()

	// Handle HTTP requests
	http.Handle("/", http.FileServer(http.Dir(*directory)))
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/signup", signup)

	// Serve
	log.Printf("Serving on: %s\n", *directory)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
