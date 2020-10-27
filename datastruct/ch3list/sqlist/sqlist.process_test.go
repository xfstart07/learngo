// Author: xufei
// Date: 2019-07-25

package sqlist

import (
	"testing"
)

func TestGetElem(t *testing.T) {
	type args struct {
		list SqList
		i    int
		e    *int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "线性表获取数据",
			args: args{
				list: SqList{
					data:   [MAXSIZE]int{1, 2, 3, 4, 5},
					length: 5,
				},
				i: 3,
				e: new(int),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetElem(tt.args.list, tt.args.i, tt.args.e); got != tt.want {
				t.Errorf("GetElem() = %v, want %v", got, tt.want)
			} else {
				t.Log("获取值", *tt.args.e)
			}
		})
	}
}

func TestListInsert(t *testing.T) {
	type args struct {
		list *SqList
		i    int
		e    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "插入元素",
			args: args{
				list: &SqList{
					data:   [MAXSIZE]int{1, 2, 3, 4, 5},
					length: 5,
				},
				i: 3,
				e: 6,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListInsert(tt.args.list, tt.args.i, tt.args.e); got != tt.want {
				t.Errorf("ListInsert() = %v, want %v", got, tt.want)
			} else {
				t.Log(tt.args.list)
			}
		})
	}
}

func TestListDelete(t *testing.T) {
	type args struct {
		list *SqList
		i    int
		e    *int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "删除元素",
			args: args{
				list: &SqList{
					data:   [MAXSIZE]int{1, 2, 6, 3, 4, 5},
					length: 6,
				},
				i: 3,
				e: new(int),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListDelete(tt.args.list, tt.args.i, tt.args.e); got != tt.want {
				t.Errorf("ListDelete() = %v, want %v", got, tt.want)
			} else {
				t.Log(tt.args.list)
			}
		})
	}
}
