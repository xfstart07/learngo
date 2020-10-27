// Author: xufei
// Date: 2019-08-05

package merge_sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "归并排序",
			args: args{
				list: []int{50, 10, 90, 30, 70, 40, 80, 60, 20},
			},
			want: []int{10, 20, 30, 40, 50, 60, 70, 80, 90},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
