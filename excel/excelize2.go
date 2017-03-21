package main

import (
	"fmt"
	"os"

	"github.com/Luxurioust/excelize"
)

func main() {
	xlsx := excelize.CreateFile()

	xlsx.SetCellValue("Sheet1", "A1", "草年末")
	xlsx.SetCellValue("Sheet1", "B1", "b1")
	xlsx.SetCellValue("Sheet1", "C1", "C1")

	xlsx.SetCellValue("Sheet1", "A2", "23")
	xlsx.SetCellValue("Sheet1", "B2", "b2")
	xlsx.SetCellValue("Sheet1", "C2", "C2")

	// Save xlsx file by the given path.
	err := xlsx.WriteTo("./excelize2.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
