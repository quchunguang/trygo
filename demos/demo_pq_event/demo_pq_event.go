package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	"time"
)

func waitForNotification(l *pq.Listener) {
	for {
		select {
		case n := <-l.Notify:
			fmt.Println("Received data from channel [", n.Channel, "] :")
			// Prepare notification payload for pretty print
			var prettyJSON bytes.Buffer
			err := json.Indent(&prettyJSON, []byte(n.Extra), "", "\t")
			if err != nil {
				fmt.Println("Error processing JSON: ", err)
			}
			fmt.Println(string(prettyJSON.Bytes()))
		case <-time.After(90 * time.Second):
			fmt.Println("Received no events for 90 seconds, checking connection")
			go l.Ping()
		}
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var conninfo = "user=postgres dbname=test host=localhost password=123456 sslmode=disable"
	_, err := sql.Open("postgres", conninfo)
	checkErr(err)

	reportProblem := func(ev pq.ListenerEventType, err error) {
		checkErr(err)
	}

	listener := pq.NewListener(conninfo, 10*time.Second, time.Minute, reportProblem)
	err = listener.Listen("events")
	checkErr(err)

	fmt.Println("Start monitoring PostgreSQL...")
	waitForNotification(listener)
}
