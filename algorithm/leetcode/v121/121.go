// Author: xufei
// Date: 2020/3/30

package v121

// 题意：股票只能交易一次，找出交易最大值
// 思路：找出交易日的最小值，同时判断当前交易是否能获取到更大值，注意边界值。时间：O(N)
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	ans := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > ans {
			ans = prices[i] - min
		}
		if min > prices[i] {
			min = prices[i]
		}
	}

	return ans
}

// 思路：循环尝试j-i的值，找出最大值，时间：O(n^n)
func maxProfitNN(prices []int) int {
	var ans int
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > ans {
				ans = prices[j] - prices[i]
			}
		}
	}

	return ans
}
