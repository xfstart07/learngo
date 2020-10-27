// Author: xufei
// Date: 2019-08-01

package binary_sort_tree

import (
	"fmt"
	"learngo/datastruct/ch8search/printer"
)

func (n *BiTNode) replaceNode(parent, elem *BiTNode) error {
	if n == nil {
		return fmt.Errorf("replace node not allow nil")
	}
	if parent.left == n {
		parent.left = elem
	} else {
		parent.right = elem
	}
	return nil
}

func (n *BiTNode) Data() interface{} {
	return n.data
}

func (n *BiTNode) LeftNode() printer.Printer {
	return n.left
}

func (n *BiTNode) RightNode() printer.Printer {
	return n.right
}
