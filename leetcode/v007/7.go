// Author: xufei
// Date: 2020/3/31

package v007

// 思路：将x每一位都mod 10，然后在乘10得出反转的值，注意[-2^31,2^31-1]边界值。
func reverse(x int) int {
	var ans int
	for x != 0 {
		m := x % 10
		ans = ans*10 + m
		x = x / 10
	}
	max := 2147483647
	min := -2147483648
	if ans > max || ans < min {
		return 0
	}
	return ans
}
