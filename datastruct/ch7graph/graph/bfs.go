// Author: xufei
// Date: 2019-08-06

package graph

// BFSTraverse 广度优先算法
func BFSTraverse(G MGraph) {
	visited := make([]bool, G.numVertexes)
	queue := make([]int, 0)

	for i := 0; i < G.numVertexes; i++ {
		if visited[i] {
			continue
		}

		visited[i] = true
		queue = append(queue, i)

		for len(queue) > 0 {
			k := queue[0]
			queue = queue[1:]

			for j := 0; j < G.numVertexes; j++ {
				if G.arc[k][j] == 1 && !visited[j] {
					queue = append(queue, i)
				}
			}
		}
	}
}
