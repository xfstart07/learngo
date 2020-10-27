// Author: xufei
// Date: 2019-08-07

package v1

import "fmt"

type Node struct {
	key, value int
	prev, next *Node
}

// O(n)
// 执行用时 : 408 ms
// 内存消耗 : 11.4 MB
type LRUCache struct {
	front, rear     *Node
	count, capacity int // 容量
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
		front:    &Node{key: -1, value: -1},
		rear:     &Node{key: -1, value: -1},
	}

	cache.front.next = cache.rear
	cache.rear.prev = cache.front

	return cache
}

func (this *LRUCache) Get(key int) int {
	node := this.find(key)
	if node != nil {
		this.remove(node)
		this.addRear(node)

		return node.value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node := this.find(key)
	if node != nil {
		this.remove(node)

		node.value = value
		this.addRear(node)
		return
	}

	node = &Node{key: key, value: value}
	if this.count == this.capacity {
		this.removeHead()
	}
	this.addRear(node)
}

func (this *LRUCache) find(key int) *Node {
	for node := this.front; node != nil; node = node.next {
		if node.key == key {
			return node
		}
	}
	return nil
}

func (this *LRUCache) print() {
	for node := this.front; node != nil; node = node.next {
		fmt.Println(node.key, node.value)
	}
}

func (this *LRUCache) addRear(node *Node) {
	this.rear.prev.next = node
	node.prev = this.rear.prev

	node.next = this.rear
	this.rear.prev = node
	this.count++
}

func (this *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev

	// 避免内存泄漏
	node.next = nil
	node.prev = nil

	this.count--
}

func (this *LRUCache) removeHead() {
	node := this.front.next
	this.front.next = node.next
	node.next.prev = this.front

	this.count--
}
