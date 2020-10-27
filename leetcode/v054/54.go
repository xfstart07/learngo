// Author: xufei
// Date: 2020/4/1

package v054

// 递归模拟
func spiralOrder(matrix [][]int) []int {
	ans := make([]int, 0)
	see := make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		see[i] = make([]bool, len(matrix[i]))
	}

	set(matrix, see, &ans, 0, 0, 1)

	return ans
}

func set(matrix [][]int, see [][]bool, ans *[]int, i, j, t int) {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[i]) || see[i][j] {
		return
	}
	*ans = append(*ans, matrix[i][j])
	see[i][j] = true

	switch t {
	case 1:
		//	左到右
		if j+1 >= len(matrix[i]) || see[i][j+1] {
			t++
		}
	case 2:
		//	上到下
		if i+1 >= len(matrix) || see[i+1][j] {
			t++
		}
	case 3:
		//	右到左
		if j-1 < 0 || see[i][j-1] {
			t++
		}
	case 4:
		//	下到上
		if i-1 < 0 || see[i-1][j] {
			t = 1
		}
	default:
		return
	}

	switch t {
	case 1:
		set(matrix, see, ans, i, j+1, t)
	case 2:
		set(matrix, see, ans, i+1, j, t)
	case 3:
		set(matrix, see, ans, i, j-1, t)
	case 4:
		set(matrix, see, ans, i-1, j, t)
	}
}
