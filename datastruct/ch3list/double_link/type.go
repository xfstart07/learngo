// Author: xufei
// Date: 2019-07-26

package double_link

type DulNode struct {
	data  int
	prior *DulNode
	next  *DulNode
}
