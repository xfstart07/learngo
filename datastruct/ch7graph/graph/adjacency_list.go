// Author: xufei
// Date: 2019-08-06

package graph

type EdgeNode struct {
	adjvex int
	weight int
	next   *EdgeNode
}

type VertexNode struct {
	data      int
	firstEdge *EdgeNode
}

type AdjList [MaxVex]VertexNode

type GraphAdjList struct {
	AdjList               AdjList
	numVertexes, numEdges int
}
