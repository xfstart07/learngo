// Author: xufei
// Date: 2019-08-06

package tree

import "log"

// PerOrderTraverse 前序遍历
func PerOrderTraverse(t *BiTNode) {
	if t == nil {
		return
	}

	log.Println(t.data)
	PerOrderTraverse(t.left)
	PerOrderTraverse(t.right)
}

// InOrderTraverse 中序遍历
func InOrderTraverse(t *BiTNode) {
	if t == nil {
		return
	}

	InOrderTraverse(t.left)
	log.Println(t.data)
	InOrderTraverse(t.right)
}

// PostOrderTraverse 后序遍历
func PostOrderTraverse(t *BiTNode) {
	if t == nil {
		return
	}

	PostOrderTraverse(t.left)
	PostOrderTraverse(t.right)
	log.Println(t.data)
}
