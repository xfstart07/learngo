// Author: xufei
// Date: 2018/12/26

package lib

func init() {
	println("sum.init")
}

func Sum(x ...int) int {
	n := 0
	for _, i := range x {
		n += i
	}

	return n
}
