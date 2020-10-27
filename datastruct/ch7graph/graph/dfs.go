// Author: xufei
// Date: 2019-08-06

package graph

// 深度遍历操作
func DFSTraverse(G MGraph) {
	visited := make([]bool, G.numVertexes)

	for i := 0; i < G.numVertexes; i++ {
		if !visited[i] {
			DFS(G, visited, i)
		}
	}
}

// 深度优先递归算法
func DFS(G MGraph, visited []bool, i int) {
	visited[i] = true

	for j := 0; j < G.numVertexes; j++ {
		if G.arc[i][j] == 1 && !visited[j] {
			DFS(G, visited, j)
		}
	}
}
