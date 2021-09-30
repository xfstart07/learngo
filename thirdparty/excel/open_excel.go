package main

import (
	"github.com/tealeg/xlsx"
	"fmt"
)

func main () {
	filePath := "public/excel.xlsx"
	file, err := xlsx.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, sheet := range file.Sheets {
		for _, rows := range sheet.Rows {
			for _, cell := range rows.Cells {
				//fmt.Println(cell.GetStyle())
				fmt.Println(cell.Value)
			}
		}
	}
}
