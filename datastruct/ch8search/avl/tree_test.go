// Author: xufei
// Date: 2019-08-01

package avl

import (
	print2 "learngo/datastruct/ch8search/printer"
	"log"
	"testing"
)

func TestBitNode_Insert(t *testing.T) {
	type args struct {
		tree *BiTree
		key  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "插入结点",
			args: args{
				key: 10,
				tree: func() *BiTree {
					t := new(BiTree)
					log.Printf("tree = %+v, tree_point = %p, root_point = %p", t, t, t.root)
					a := []int{3, 2, 1, 4, 5, 6, 7, 10, 9, 8}
					for i := 0; i < len(a); i++ {
						t.root = t.root.Insert(a[i])
					}

					var node *BitNode
					log.Printf("%+v, %p", node, node)
					//node = node.Insert(10) // Receiver 'node' may be 'nil' in call

					return t
				}(),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.tree.root.Insert(tt.args.key); (got != nil) != tt.want {
				t.Errorf("BitNode.Insert() = %v, want %v", got, tt.want)
			} else {
				print2.Preorder(got)
			}
		})
	}
}

func TestBitNode_Delete(t *testing.T) {
	type args struct {
		tree *BiTree
		key  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "删除结点",
			args: args{
				key:  7,
				tree: newTree(),
			},
			want: true,
		},
		{
			name: "删除结点",
			args: args{
				key:  11,
				tree: newTree(),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.tree.root.Delete(tt.args.key); (got != nil) != tt.want {
				t.Errorf("BitNode.Delete() = %v, want %v", got, tt.want)
			} else {
				print2.Preorder(got)
			}
		})
	}
}

func newTree() *BiTree {
	t := new(BiTree)
	log.Printf("tree = %+v, tree_point = %p, root_point = %p", t, t, t.root)
	a := []int{3, 2, 1, 4, 5, 6, 7, 10, 9, 8}
	for i := 0; i < len(a); i++ {
		t.root = t.root.Insert(a[i])
	}
	return t
}
