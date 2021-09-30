// Author: xufei
// Date: 2019-10-23

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//var a int, b int, c int
	var (
		a int
		b int
		c int
	)

	fmt.Println(a, b, c)

	// 声明变量时指定类型，但一般类型Go可以推断出来，所以基本都是省略的
	var str string = "1111"
	fmt.Println(str)

	// NOTE: 声明时指定类型，这样就表明 jm 变量一定实现了 json.Marshaler 接口的方法。
	// NOTE: EffectiveGo 中说明：“在此声明中，我们调用了一个 *RawMessage 转换并将其赋予了 Marshaler，以此来要求 *RawMessage 实现 Marshaler，这时其属性就会在编译时被检测。 若 json.Marshaler 接口被更改，此包将无法通过编译， 而我们则会注意到它需要更新。”
	var jm json.Marshaler = (*json.RawMessage)(nil)
	_, _ = jm.MarshalJSON()
	//NOTE: 这种方法可以在我们明确需要某个类型实现某个接口时使用，使得在编译时就能发现错误。
}
