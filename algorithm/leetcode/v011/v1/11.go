// Author: xufei
// Date: 2020/4/1

package v1

//思路：两层循环尝试每一种可能容纳的值，找出最大值，容纳值计算：min(height[i],height[j])*(j-i)
func maxArea(height []int) int {
	var ans int
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			k := min(height[i], height[j])
			if k*(j-i) > ans {
				ans = k * (j - i)
			}
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
