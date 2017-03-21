package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func main() {
	excelFileName := "./MyXLSXFile.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf(err.Error())
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text, _ := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}
}
