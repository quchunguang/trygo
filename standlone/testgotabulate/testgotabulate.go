package main

import (
	"fmt"
	"github.com/bndr/gotabulate"
)

var STRING_ARRAY = []string{"test string", "test string 2", "test", "row", "bndr"}

func tbl1() {
	// Create Some Fake Rows
	row_1 := []interface{}{"john", 20, "ready"}
	row_2 := []interface{}{"bndr", 23, "ready"}

	// Create an object from 2D interface array
	t := gotabulate.Create([][]interface{}{row_1, row_2})

	// Set the Headers (optional)
	t.SetHeaders([]string{"age", "status"})

	// Set the Empty String (optional)
	t.SetEmptyString("None")

	// Set Align (Optional)
	t.SetAlign("right")

	// Print the result: grid, or simple
	fmt.Println(t.Render("grid"))
}

func tbl2() {
	// Some Strings
	string_1 := []string{"TV", "1000$", "Sold"}
	string_2 := []string{"PC", "50%", "on Hold"}

	// Create Object
	tabulate := gotabulate.Create([][]string{string_1, string_2})

	// Set Headers
	tabulate.SetHeaders([]string{"Type", "Cost", "Status"})

	// Render
	fmt.Println(tabulate.Render("simple"))
}

func tbl3() {
	tabulate := gotabulate.Create([][]string{[]string{"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus laoreet vestibulum pretium. Nulla et ornare elit. Cum sociis natoque penatibus et magnis",
		"Vivamus laoreet vestibulum pretium. Nulla et ornare elit. Cum sociis natoque penatibus et magnis", "zzLorem ipsum", " test", "test"}, []string{"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus laoreet vestibulum pretium. Nulla et ornare elit. Cum sociis natoque penatibus et magnis",
		"Vivamus laoreet vestibulum pretium. Nulla et ornare elit. Cum sociis natoque penatibus et magnis", "zzLorem ipsum", " test", "test"}, STRING_ARRAY, []string{"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus laoreet vestibulum pretium. Nulla et ornare elit. Cum sociis natoque penatibus et magnis",
		"Vivamus laoreet vestibulum pretium. Nulla et ornare elit. Cum sociis natoque penatibus et magnis", "zzLorem ipsum", " test", "test"}, STRING_ARRAY})

	tabulate.SetHeaders([]string{"Header 1", "header 2", "header 3", "header 4"})
	// Set Max Cell Size
	tabulate.SetMaxCellSize(16)

	// Turn On String Wrapping
	tabulate.SetWrapStrings(true)

	// Render the table
	fmt.Println(tabulate.Render("grid"))
}

func main() {
	tbl1()
	tbl2()
	tbl3()
}
