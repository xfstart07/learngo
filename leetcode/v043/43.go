package v043

import (
	"strconv"
	"strings"
)

// 大数模拟乘法
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	nums := make([]int, len(num1)+len(num2))

	k := 0
	for i := len(num1) - 1; i >= 0; i-- {
		x, _ := strconv.Atoi(string(num1[i]))

		p := k
		for j := len(num2) - 1; j >= 0; j-- {
			y, _ := strconv.Atoi(string(num2[j]))

			nums[p] += x * y
			p++
		}
		k++
	}

	s := ""
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 10 {
			v := nums[i] / 10
			nums[i] = nums[i] % 10
			nums[i+1] += v
		}
		s = strconv.Itoa(nums[i]) + s
	}
	s = strings.TrimLeft(s, "0")

	return s
}
