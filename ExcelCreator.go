package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func makeExcel() {
	xlsx := excelize.NewFile()
	// Create a new sheet.
	// tadada
	index := xlsx.NewSheet("Sheet2")
	// Set value of a cell.
	xlsx.SetCellValue("Sheet2", "A2", "Hello world.")
	xlsx.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("./Workbook.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
