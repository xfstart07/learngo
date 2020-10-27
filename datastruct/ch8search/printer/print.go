// Author: xufei
// Date: 2019-08-02

package printer

import (
	"log"
	"reflect"
)

type Printer interface {
	Data() interface{}
	LeftNode() Printer
	RightNode() Printer
}

// Preorder 中序打印
func Preorder(t Printer) {
	if reflect.ValueOf(t).IsNil() {
		return
	}

	log.Println(t.Data()) // 前序
	Preorder(t.LeftNode())
	//log.Println(t.data) // 中序
	Preorder(t.RightNode())
	//log.Println(t.data) // 后序
}
