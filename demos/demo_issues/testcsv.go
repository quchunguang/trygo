package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"
)

func testcsv() {
	in := `first_name;last_name;username
	Rob;Pike;rob
	Ken;Thompson;ken
	Robert;Griesemer;gri`
	r := csv.NewReader(strings.NewReader(in))
	r.Comma = ';'
	r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(records)
}
