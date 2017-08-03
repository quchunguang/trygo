package main

import (
	"fmt"
	"log"
	"net/http"

	"code.google.com/p/go.net/websocket"
)

func Echo(ws *websocket.Conn) {
	var err error

	for {

		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func main() {
	fmt.Println("begin")
	http.Handle("/", http.FileServer(http.Dir("."))) // <-- note this line

	http.Handle("/socket", websocket.Handler(Echo))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	fmt.Println("end")
}
