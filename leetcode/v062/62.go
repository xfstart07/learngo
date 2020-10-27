// Author: xufei
// Date: 2020/4/1

package v062

// 动态规划，机器人每次只能向下和向右走
// 在二维数组中，f[i][j] 的路径就等于 f[i-1][j] + f[i][j-1]
// 边界值，f[m][1], f[1][n] 的路径都是 1.
func uniquePaths(m int, n int) int {
	f := make([][]int, m)

	for i := 0; i < m; i++ {
		f[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		f[0][j] = 1
	}

	for i := 1; i < m; i++ {
		f[i][0] = 1
		for j := 1; j < n; j++ {
			f[i][j] = f[i-1][j] + f[i][j-1]
		}
	}

	return f[m-1][n-1]
}
