package main

import (
	"fmt"
	"time"
)

// 时间的解析，格式化带上时区
func main() {

	kindOf()

}

func defaultTime() {
	//timeLayout := "2006-01-02 15:04:05 MST"
	timeLayout := "2006-01-02 15:04:05"

	fmt.Println("时区", time.Local)
	fmt.Println(time.Now().Format(timeLayout))

	timeStr := time.Now().Format(timeLayout)

	//loc, _ := time.LoadLocation("Local")
	parseTime, _ := time.Parse(timeLayout, timeStr)
	fmt.Println(parseTime)

	fmt.Println(time.Now().UnixNano())
}

func kindOf() {
	timeLayout := "2006-1-2 15:04:05"

	temp := "2016-2-17 3:18:00"

	str, err := time.Parse(timeLayout, temp)
	fmt.Println(str, err)
}
