// Author: xufei
// Date: 2019-07-25

package single_link

import (
	"testing"
)

func TestNode_GetElem(t *testing.T) {
	list := new(Node)
	head := list
	for i := 1; i <= 5; i++ {
		node := new(Node)
		node.data = i
		list.next = node
		list = node
	}
	head = head.next

	type args struct {
		i    int
		e    *int
		list *Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "获取列表元素",
			args: args{
				i:    3,
				e:    new(int),
				list: head,
			},
			want: true,
		},
		{
			name: "获取列表元素失败",
			args: args{
				i:    6,
				e:    new(int),
				list: head,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.list.GetElem(tt.args.i, tt.args.e); got != tt.want {
				t.Errorf("Node.GetElem() = %v, want %v", got, tt.want)
			} else {
				t.Log(*tt.args.e)
			}
		})
	}
}

func TestNode_ListInsert(t *testing.T) {
	type args struct {
		i    int
		e    int
		list *Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "插入元素",
			args: args{
				i:    3,
				e:    6,
				list: createList(),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.list.ListInsert(tt.args.i, tt.args.e); got != tt.want {
				t.Errorf("Node.ListInsert() = %v, want %v", got, tt.want)
			} else {
				for node := tt.args.list; node != nil; node = node.next {
					t.Logf("%d", node.data)
				}
			}
		})
	}
}

func TestNode_ListDelete(t *testing.T) {
	type args struct {
		i    int
		e    *int
		list *Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "删除元素",
			args: args{
				i:    3,
				e:    new(int),
				list: createList(),
			},
			want: true,
		},
		{
			name: "删除第 1 元素",
			args: args{
				i:    1,
				e:    new(int),
				list: createList(),
			},
			want: false,
		},
		{
			name: "删除最后一个元素",
			args: args{
				i:    5,
				e:    new(int),
				list: createList(),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.list.ListDelete(tt.args.i, tt.args.e); got != tt.want {
				t.Errorf("Node.ListDelete() = %v, want %v", got, tt.want)
			} else {
				t.Log("删除值", *tt.args.e)
				for node := tt.args.list; node != nil; node = node.next {
					t.Logf("%d", node.data)
				}
			}
		})
	}
}

func createList() *Node {
	list := new(Node)
	head := list
	for i := 1; i <= 5; i++ {
		node := new(Node)
		node.data = i
		list.next = node
		list = node
	}
	return head.next
}
