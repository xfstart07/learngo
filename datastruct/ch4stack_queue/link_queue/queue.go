// Author: xufei
// Date: 2019-07-29

package link_queue

func New() *LinkQueue {
	return &LinkQueue{}
}

func (s *LinkQueue) Insert(elem int) bool {
	node := new(Node)
	node.data = elem

	s.rear.next = node
	s.rear = node
	return true
}

func (s *LinkQueue) Delete(elem *int) bool {
	if s.front == nil {
		return false
	}
	*elem = s.front.data
	s.front = s.front.next

	return true
}
