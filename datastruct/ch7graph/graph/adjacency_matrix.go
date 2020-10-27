// Author: xufei
// Date: 2019-08-06

package graph

const (
	MaxVex   = 100
	Infinity = 65535
)

type MGraph struct {
	vexs                  [MaxVex]int
	arc                   [MaxVex][MaxVex]int
	numVertexes, numEdges int
}
