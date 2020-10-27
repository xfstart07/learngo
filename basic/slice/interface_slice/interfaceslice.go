// Author: xufei
// Date: 2019/4/21

package main

import (
	"fmt"
	"reflect"
)

type BulkList []interface{}

type MyStruct struct {
	Name string
}

func main() {
	list := []string{"a", "b"}

	myStructSlice := []*MyStruct{}
	myStructSlice = append(myStructSlice, &MyStruct{
		Name: "leon",
	})
	myStructSlice = append(myStructSlice, &MyStruct{
		Name: "xu",
	})

	handle(list)          // len = 1
	handle(myStructSlice) // len =1

	//handle(myStructSlice...) // Cannot use 'myStructSlice' (type []*MyStruct) as type []interface{}

	handle(myStructSlice[0], myStructSlice[1]) // len = 2 这种方式可以正确输出，但是非常局限

	// value print
	myStructSliceValue := []MyStruct{}
	myStructSliceValue = append(myStructSliceValue, MyStruct{
		Name: "leon",
	})
	myStructSliceValue = append(myStructSliceValue, MyStruct{
		Name: "xu",
	})
	handleValue(myStructSliceValue) // len = 1

	bulkList := BulkList{}
	for _, r := range myStructSlice {
		bulkList = append(bulkList, r)
	}
	handleBulkList(bulkList) // len = 2
}

func handle(list ...interface{}) {
	fmt.Println(fmt.Sprintf("len=%d", len(list)))
	for _, value := range list {
		fmt.Println(value, reflect.TypeOf(value).String())
	}
}

func handleValue(list ...interface{}) {
	fmt.Println(fmt.Sprintf("len=%d", len(list)))
	for _, value := range list {
		fmt.Println(value)
	}
}

func handleBulkList(list BulkList) {
	fmt.Println(fmt.Sprintf("len=%d", len(list)))
	for _, value := range list {
		fmt.Println(value)
	}
}
