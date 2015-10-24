package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("Sheet1")
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "I am a cell!"

	if err := file.Save("MyXLSXFile.xlsx"); err != nil {
		fmt.Printf(err.Error())
	}
}
