// Author: xufei
// Date: 2019-08-01

package binary_sort_tree

import (
	"testing"
)

func TestSearchBST(t *testing.T) {
	type args struct {
		t      *BiTree
		key    int
		parent *BiTNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "二叉树查找",
			args: args{
				t:      new(BiTree),
				key:    89,
				parent: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got := tt.args.t.Search(tt.args.t.root, tt.args.key, tt.args.parent); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBiTNode_Insert(t *testing.T) {
	type args struct {
		key int
		t   *BiTree
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "插入",
			args: args{
				key: 10,
				t:   new(BiTree),
			},
			want: true,
		},
		{
			name: "插入失败，已存在",
			args: args{
				key: 10,
				t: func() *BiTree {
					n := new(BiTree)
					n.Insert(10)
					return n
				}(),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.t.Insert(tt.args.key); got != tt.want {
				t.Errorf("BiTNode.Insert() = %v, want %v", got, tt.want)
			} else {
				Print(tt.args.t.root)
			}
		})
	}
}

func TestBiTree_DeleteBST(t *testing.T) {
	type args struct {
		tree   *BiTree
		parent *BiTNode
		key    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "删除节点",
			args: args{
				tree: func() *BiTree {
					tree := new(BiTree)
					a := []int{62, 88, 58, 47, 35, 73, 51, 99, 37, 93}
					for _, v := range a {
						tree.Insert(v)
					}
					return tree
				}(),
				parent: nil,
				key:    47,
			},
			want: true,
		},
		{
			name: "删除节点",
			args: args{
				tree: func() *BiTree {
					tree := new(BiTree)
					a := []int{62, 88, 58, 47, 35, 73, 51, 99, 37, 93}
					for _, v := range a {
						tree.Insert(v)
					}
					return tree
				}(),
				parent: nil,
				key:    37,
			},
			want: true,
		},
		{
			name: "删除节点",
			args: args{
				tree: func() *BiTree {
					tree := new(BiTree)
					a := []int{62, 88, 58, 47, 35, 73, 51, 99, 37, 93}
					for _, v := range a {
						tree.Insert(v)
					}
					return tree
				}(),
				parent: nil,
				key:    58,
			},
			want: true,
		},
		{
			name: "删除节点",
			args: args{
				tree: func() *BiTree {
					tree := new(BiTree)
					a := []int{62}
					for _, v := range a {
						tree.Insert(v)
					}
					return tree
				}(),
				parent: nil,
				key:    62,
			},
			want: true,
		},
		{
			name: "删除根节点",
			args: args{
				tree: func() *BiTree {
					tree := new(BiTree)
					a := []int{62, 88, 58, 47, 35, 73, 51, 99, 37, 93}
					for _, v := range a {
						tree.Insert(v)
					}
					return tree
				}(),
				parent: nil,
				key:    62,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.tree.Delete(tt.args.tree.root, tt.args.parent, tt.args.key); got != tt.want {
				t.Errorf("BiTree.Delete() = %v, want %v", got, tt.want)
			} else {
				Print(tt.args.tree.root)
			}
		})
	}
}
