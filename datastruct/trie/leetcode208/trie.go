// Author: xufei
// Date: 2019-08-21

package leetcode208

type node struct {
	ch          string
	children    [26]*node
	isWordOfEnd bool
}

//执行用时 : 112 ms
//内存消耗 : 16.9 MB
type Trie struct {
	root *node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	t := Trie{
		root: &node{
			ch:          "/",
			children:    [26]*node{},
			isWordOfEnd: false,
		},
	}
	return t
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	n := this.root

	for i := 0; i < len(word); i++ {
		wid := word[i] - 'a'
		if n.children[wid] == nil {
			n.children[wid] = &node{
				ch:          word[i : i+1],
				children:    [26]*node{},
				isWordOfEnd: false,
			}
		}
		n = n.children[wid]
	}
	n.isWordOfEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	n := this.root

	for i := 0; i < len(word); i++ {
		wid := word[i] - 'a'
		if n.children[wid] == nil {
			return false
		}
		n = n.children[wid]
	}

	return n.isWordOfEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	n := this.root

	for i := 0; i < len(prefix); i++ {
		wid := prefix[i] - 'a'
		if n.children[wid] == nil {
			return false
		}
		n = n.children[wid]
	}

	return true
}
