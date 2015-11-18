package main

import (
	"fmt"
	"github.com/gocraft/web"
	"net/http"
	"strings"
)

type Context struct {
	HelloCount int
}

func (c *Context) UserRequired(rw web.ResponseWriter, r *web.Request, next web.NextMiddlewareFunc) {
	user := userFromSession(r) // Pretend like this is defined. It reads a session cookie and returns a *User or nil.
	if user != nil {
		c.User = user
		next(rw, r)
	} else {
		rw.Header().Set("Location", "/")
		rw.WriteHeader(http.StatusMovedPermanently)
		// do NOT call next()
	}
}

func (c *Context) SetHelloCount(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HelloCount = 3
	fmt.Println(c.HelloCount)
	next(rw, req)
}

func (c *Context) SayHello(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, strings.Repeat("Hello ", c.HelloCount), "World!")
}

func main() {
	router := web.New(Context{}). // Create your router
					Middleware(web.LoggerMiddleware).     // Use some included middleware
					Middleware(web.ShowErrorsMiddleware). // ...
					Middleware((*Context).SetHelloCount). // Your own middleware!
					Get("/", (*Context).SayHello)         // Add a route
	http.ListenAndServe("127.0.0.1:3000", router) // Start the server!
}
