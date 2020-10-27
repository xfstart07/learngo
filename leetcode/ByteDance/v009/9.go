package v009

// https://leetcode-cn.com/explore/interview/card/bytedance/243/array-and-sorting/1034/
// 枚举+DFS
// 每个位置上是 1 的位置，从这个位置开始上下左右的走，每走一步将位置设置为 0
func maxAreaOfIsland(grid [][]int) int {
	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}

			step := 0
			dfs(grid, i, j, &step)
			if step > max {
				max = step
			}
		}
	}

	return max
}

func dfs(grid [][]int, i, j int, step *int) {
	grid[i][j] = 0
	*step++

	if i-1 >= 0 && grid[i-1][j] == 1 {
		dfs(grid, i-1, j, step)
	}
	if i+1 < len(grid) && grid[i+1][j] == 1 {
		dfs(grid, i+1, j, step)
	}
	if j-1 >= 0 && grid[i][j-1] == 1 {
		dfs(grid, i, j-1, step)
	}
	if j+1 < len(grid[i]) && grid[i][j+1] == 1 {
		dfs(grid, i, j+1, step)
	}
}
