// Author: xufei
// Date: 2019-07-29

package link_queue

type Node struct {
	data int
	next *Node
}

type LinkQueue struct {
	front, rear *Node
}
