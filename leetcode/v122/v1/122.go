// Author: xufei
// Date: 2020/3/27

package v1

func maxProfit(prices []int) int {
	return calc(prices, 0)
}

// 暴力算法：会超时，算法时间复杂度 O(n^n)
// 从第s个位置，最优解是多少
func calc(prices []int, s int) int {
	if s >= len(prices) {
		return 0
	}

	max := 0
	// i 买，j 卖
	for i := s; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			tmp := calc(prices, j+1) + prices[j] - prices[i]

			if tmp > max {
				max = tmp
			}
		}
	}

	return max
}
