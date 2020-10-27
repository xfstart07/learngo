// Author: Xu Fei
// Date: 2018/8/19
package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(*p)

	//p2 := &int  // 不允许通过 &int 来获取地址，因为 int 是不允许获取地址的。
	//fmt.Println(*p2)
}
