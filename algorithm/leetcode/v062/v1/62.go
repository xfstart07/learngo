// Author: xufei
// Date: 2020/4/1

package v1

// DFS，递归尝试每一步
func uniquePaths(m int, n int) int {
	return dfs(1, 1, m, n)
}

func dfs(i, j, m, n int) int {
	if i <= 0 || i > m || j <= 0 || j > n {
		return 0
	}
	if i == m && j == n {
		return 1
	}

	return dfs(i, j+1, m, n) + dfs(i+1, j, m, n)
}
