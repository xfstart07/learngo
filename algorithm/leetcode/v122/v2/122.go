// Author: xufei
// Date: 2020/3/27

package v2

// 峰谷
func maxProfit(prices []int) int {
	var max int

	n := 0
	for n < len(prices) {
		//寻找峰谷
		//	寻找谷值
		for n < len(prices)-1 && prices[n] >= prices[n+1] {
			n++
		}
		valley := prices[n]

		// 寻找峰值
		for n < len(prices)-1 && prices[n] < prices[n+1] {
			n++
		}
		peak := prices[n]
		n++

		max = max + peak - valley
	}

	return max
}
