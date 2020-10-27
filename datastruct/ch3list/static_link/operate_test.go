// Author: xufei
// Date: 2019-07-26

package static_link

import (
	"log"
	"testing"
)

func TestInitList(t *testing.T) {
	space := new(StaticLinkList)
	type args struct {
		space *StaticLinkList
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "生成链表",
			args: args{
				space: space,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitList(tt.args.space); got != tt.want {
				t.Errorf("InitList() = %v, want %v", got, tt.want)
			} else {
				for i := 0; i < MAXSIZE; i++ {
					log.Println(tt.args.space[i].data, " ", tt.args.space[i].cur)
				}
			}
		})
	}
}

func TestListInsert(t *testing.T) {
	type args struct {
		list *StaticLinkList
		i    int
		e    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "插入数据",
			args: args{
				list: func() *StaticLinkList {
					space := new(StaticLinkList)
					InitList(space)
					ListInsert(space, 1, 100)
					return space
				}(),
				i: 2,
				e: 101,
			},
			want: true,
		},
		{
			name: "在空链表中，插入无效位置数据",
			args: args{
				list: func() *StaticLinkList {
					space := new(StaticLinkList)
					InitList(space)
					return space
				}(),
				i: 4,
				e: 104,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < MAXSIZE; i++ {
				log.Println(tt.args.list[i].data, " ", tt.args.list[i].cur)
			}
			log.Println("执行前...")

			if got := ListInsert(tt.args.list, tt.args.i, tt.args.e); got != tt.want {
				t.Errorf("ListInsert() = %v, want %v", got, tt.want)
			} else {
				for i := 0; i < MAXSIZE; i++ {
					log.Println(tt.args.list[i].data, " ", tt.args.list[i].cur)
				}
				log.Println("执行前...")
			}
		})
	}
}

func TestListDelete(t *testing.T) {
	type args struct {
		list *StaticLinkList
		i    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "删除数据",
			args: args{
				list: func() *StaticLinkList {
					space := new(StaticLinkList)
					InitList(space)
					ListInsert(space, 1, 100)
					ListInsert(space, 2, 102)
					ListInsert(space, 3, 103)
					return space
				}(),
				i: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < MAXSIZE; i++ {
				log.Println(tt.args.list[i].data, " ", tt.args.list[i].cur)
			}
			log.Println("执行前...")

			if got := ListDelete(tt.args.list, tt.args.i); got != tt.want {
				t.Errorf("ListDelete() = %v, want %v", got, tt.want)
			} else {
				for i := 0; i < MAXSIZE; i++ {
					log.Println(tt.args.list[i].data, " ", tt.args.list[i].cur)
				}
				log.Println("执行前...")
			}
		})
	}
}
