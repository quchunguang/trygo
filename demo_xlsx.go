package trygo

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func DemoXlsx() {
	excelFileName := "/home/qcg/share/20140912.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				s, _:=cell.String()
				fmt.Printf("%s\n", s)
			}
		}
	}
}
