// Author: xufei
// Date: 2019-07-29

package circular_queue

import "math"

func New() *SQueue {
	return &SQueue{
		front: 0,
		rear:  0,
	}
}

func (q *SQueue) Length() int {
	return int(math.Mod(float64(q.rear-q.front+MAXSIZE), MAXSIZE))
}

func (q *SQueue) Insert(elem int) bool {
	tail := int(math.Mod(float64(q.rear+1), MAXSIZE))
	if tail == q.front { // 队列已满
		return false
	}

	q.elements[q.rear] = elem
	q.rear = tail
	return true
}

func (q *SQueue) Delete(elem *int) bool {
	if q.front == q.rear { // 队列为空
		return false
	}

	*elem = q.elements[q.front]
	head := int(math.Mod(float64(q.front+1), MAXSIZE))
	q.front = head
	return true
}
