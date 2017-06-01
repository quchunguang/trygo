package main

import (
	"bufio"
	"github.com/tarm/serial"
	"log"
	"net/http"
)

const (
	BAUD   = 9600
	DEVICE = "COM4"
)

var port *serial.Port

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func comm(w http.ResponseWriter, r *http.Request) {
	reader := bufio.NewReader(port)
	reply, err := reader.ReadBytes('\x0a')
	if err != nil {
		log.Print("E:", string(reply))
		w.Write([]byte("000000000000"))
		return
	}

	log.Print("R:", string(reply))
	w.Write(reply[:8])
}

func main() {
	var err error
	c := &serial.Config{Name: DEVICE, Baud: BAUD}
	port, err = serial.OpenPort(c)
	defer port.Close()
	checkErr(err)
	http.HandleFunc("/comm", comm)
	log.Fatal(http.ListenAndServe(":8300", nil))
}
