// Author: xufei
// Date: 2019-07-31

package sequential

func Sequential_Search(a []int, n, key int) int {
	for i := 1; i <= n; i++ {
		if a[i] == key {
			return i
		}
	}
	return 0
}

// 有哨兵顺序查找
func Sequential_Search2(a []int, n, key int) int {
	a[0] = key
	var i int
	for i = n; a[i] != key; i-- {
	}
	return i
}
