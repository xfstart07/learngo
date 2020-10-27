// Author: xufei
// Date: 2020/4/1

package v054

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "spiralOrder",
			args: args{
				matrix: func() [][]int {
					n := 3
					matrix := make([][]int, n)
					total := 0
					for i := 0; i < n; i++ {
						matrix[i] = make([]int, n)
						for j := 0; j < n; j++ {
							total++
							matrix[i][j] = total
						}
					}
					return matrix
				}(),
			},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "spiralOrder",
			args: args{
				matrix: func() [][]int {
					n, m := 3, 4

					matrix := make([][]int, n)
					total := 0
					for i := 0; i < n; i++ {
						matrix[i] = make([]int, m)
						for j := 0; j < m; j++ {
							total++
							matrix[i][j] = total
						}
					}
					return matrix
				}(),
			},
			want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
