package main

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"fmt"
)

func main() {
	xlsx, err := excelize.OpenFile("public/excel.xlsx")
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

}
