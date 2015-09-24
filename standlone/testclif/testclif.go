package main

import "gopkg.in/ukautz/clif.v0"

func cmdHello(out clif.Output) { out.Printf("Hello World\n") }

func main() {
	cli := clif.New("My App", "1.0.0", "An example application")
	cli.New("hello", "The obligatory hello world", cmdHello)
	cli.Run()
}
