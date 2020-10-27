// Author: xufei
// Date: 2020/3/25

package v136

//136. 只出现一次的数字

// 使用 map 标记出现的数字，重复的删掉，最后只剩一个元素
func singleNumber(nums []int) int {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if !m[v] {
			m[v] = true
		} else {
			delete(m, v)
		}
	}

	for k, v := range m {
		if v {
			return k
		}
	}

	return 0
}

func singleNumberXOR(nums []int) int {
	var ans int
	for i := 0; i < len(nums); i++ {
		ans = ans ^ nums[i]
	}

	return ans
}
