// Author: xufei
// Date: 2020/3/27

package v122

func maxProfit(prices []int) int {
	var max, tmp int
	for i := 1; i < len(prices); i++ {
		tmp = prices[i] - prices[i-1]
		if tmp > 0 {
			max += tmp
		}
	}
	return max
}
