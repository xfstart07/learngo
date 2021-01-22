// Author: xufei
// Date: 2020/3/31

package v059

// 递归模拟
// 规则：1.左到右，2. 上到下，3. 右到左，4. 下到上，循环这4个规则
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	set(matrix, 0, 0, 1, 1)
	return matrix
}

func set(matrix [][]int, i, j, k, t int) {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix) || matrix[i][j] > 0 {
		return
	}
	matrix[i][j] = k
	k++

	switch t {
	case 1:
		//	左到右
		if j+1 >= len(matrix) || matrix[i][j+1] > 0 {
			t++
		}
	case 2:
		//	上到下
		if i+1 >= len(matrix) || matrix[i+1][j] > 0 {
			t++
		}
	case 3:
		//	右到左
		if j-1 < 0 || matrix[i][j-1] > 0 {
			t++
		}
	case 4:
		//	下到上
		if i-1 < 0 || matrix[i-1][j] > 0 {
			t = 1
		}
	default:
		return
	}

	switch t {
	case 1:
		set(matrix, i, j+1, k, t)
	case 2:
		set(matrix, i+1, j, k, t)
	case 3:
		set(matrix, i, j-1, k, t)
	case 4:
		set(matrix, i-1, j, k, t)
	}
}
