// Author: xufei
// Date: 2020/3/5

package service

func GetSum(n int64) int64 {
	var sum int64
	var i int64
	for ; i <= n; i++ {
		sum += i
	}

	return sum
}
