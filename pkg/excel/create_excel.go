package main

import (
	"github.com/tealeg/xlsx"
	"fmt"
)

func main() {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Println("创建工作表失败", err)
		return
	}

	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "我是一个表格"

	err = file.Save("public/excel.xlsx")
	fmt.Println(err)
}
