package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Luxurioust/excelize"
)

func main() {
	xlsx := excelize.CreateFile()

	xlsx.SetCellValue("Sheet1", "A1", "名字")
	xlsx.SetCellValue("Sheet1", "B1", "年龄")
	xlsx.SetCellValue("Sheet1", "C1", "大小")

	for i := 0; i < 10; i++ {

		num := 2 + i
		f := strconv.Itoa(num)
		a := "A" + f
		b := "B" + f
		c := "C" + f

		xlsx.SetCellValue("Sheet1", a, "查查")
		xlsx.SetCellValue("Sheet1", b, "20")
		xlsx.SetCellValue("Sheet1", c, "123")
	}

	// Save xlsx file by the given path.
	err := xlsx.WriteTo("./excelize2.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
