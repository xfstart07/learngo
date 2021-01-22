// Author: xufei
// Date: 2020/4/1

package v011

// 思路：当不考虑高度的情况下，宽度越大越好，
// 我们可以从数组的头尾开始计算容器的容纳值，之后对高度低的一端进行移位，看是否有更高的高度，依次来寻找最大值
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	var ans int
	for i < j {
		k := min(height[i], height[j]) * (j - i)
		if k > ans {
			ans = k
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return ans
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
