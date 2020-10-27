// Author: xufei
// Date: 2019-06-26

package maps

import "fmt"

func MakeLenMap() {
	// make map 设置了 len，表示 map 已经存在 len 个元素了。
	keys := make([]string, 10)
	for key, value := range keys {
		fmt.Println(key, value)
	}
	fmt.Println(len(keys))
}
