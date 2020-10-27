// Author: xufei
// Date: 2019-07-29

package circular_queue

import (
	"testing"
)

func TestSQueue_Insert(t *testing.T) {
	type args struct {
		queue *SQueue
		elem  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "队列添加元素",
			args: args{
				queue: New(),
				elem:  100,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.queue.Insert(tt.args.elem); got != tt.want {
				t.Errorf("SQueue.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQueue_Delete(t *testing.T) {
	type args struct {
		queue *SQueue
		call  func(q *SQueue)
		elem  *int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "删除队列元素",
			args: args{
				queue: New(),
				call: func(q *SQueue) {
					q.Insert(100)
					q.Insert(101)
				},
				elem: new(int),
			},
			want: true,
		},
		{
			name: "删除队列元素失败，空队列",
			args: args{
				queue: New(),
				call:  func(q *SQueue) {},
				elem:  new(int),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.call(tt.args.queue)

			if got := tt.args.queue.Delete(tt.args.elem); got != tt.want {
				t.Errorf("SQueue.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
