package v238

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "productExceptSelf",
			args: args{
				nums: []int{1, 2, 0},
			},
			want: []int{0, 0, 2},
		},
		{
			name: "productExceptSelf",
			args: args{
				nums: []int{1, 2, 0, 0},
			},
			want: []int{0, 0, 0, 0},
		},
		{
			name: "productExceptSelf",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: []int{24, 12, 8, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productExceptSelf(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}
