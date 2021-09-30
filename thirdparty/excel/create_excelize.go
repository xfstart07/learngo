package main

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"fmt"
)

func main() {
	xlsx, err := excelize.OpenFile("public/excel2.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取工作表，打印出数据
	rows := xlsx.GetRows("Sheet1")
	for _, row := range rows {
		for _, cell := range row {
			fmt.Println(cell)
		}
	}

	sheetName := xlsx.GetSheetName(1)

	cell := xlsx.GetCellValue(sheetName, "A1")
	fmt.Println(cell)

	// 合并 A1-E1
	xlsx.MergeCell(sheetName, "A1", "E1")

	// A1-E1 的内容居中
	style, err := xlsx.NewStyle(`{"alignment":{"horizontal":"center"}}`)
	if err != nil {
		fmt.Println("创建样式失败", err)
		return
	}
	xlsx.SetCellStyle(sheetName, "A1", "E1", style)

	xlsx.Save()
}
