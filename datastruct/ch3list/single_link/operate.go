// Author: xufei
// Date: 2019-07-25

package single_link

import "log"

func (l *Node) GetElem(i int, e *int) bool {
	p := l.next
	j := 1
	for ; p != nil && j < i; j++ {
		p = p.next
	}
	if p == nil || j > i {
		return false
	}

	*e = p.data
	return true
}

func (l *Node) ListInsert(i, e int) bool {
	p := l
	j := 1

	for ; p != nil && j < i; j++ {
		p = p.next
	}

	if p == nil || j > i {
		return false
	}

	s := new(Node)
	s.data = e

	s.next = p.next
	p.next = s

	return true
}

// 查找到第 i 个元素
// TODO: 第 1 个元素无法删除
func (l *Node) ListDelete(i int, e *int) bool {
	p := l
	j := 1
	for ; p.next != nil && j < i-1; j++ {
		p = p.next
	}
	log.Printf("i = %d, j = %d", i, j)
	log.Printf("%+v, %+v", p, p.next)

	if p.next == nil || j >= i {
		return false
	}

	del := p.next
	p.next = del.next
	*e = del.data

	return true
}
