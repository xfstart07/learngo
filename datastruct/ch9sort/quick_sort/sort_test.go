// Author: xufei
// Date: 2019-08-05

package quick_sort

import (
	"reflect"
	"testing"
)

func TestQSort(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "快排序",
			args: args{
				list: []int{50, 10, 90, 30, 70, 40, 80, 60, 20},
			},
			want: []int{10, 20, 30, 40, 50, 60, 70, 80, 90},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QSort(tt.args.list, 0, len(tt.args.list)-1)
			t.Log(tt.args.list)
			if !reflect.DeepEqual(tt.want, tt.args.list) {
				t.Errorf("QSort() = %v, want %v", tt.args.list, tt.want)
			}
		})
	}
}
