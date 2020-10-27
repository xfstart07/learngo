// Author: xufei
// Date: 2020/4/1

package v062

import "testing"

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "dfs",
			args: args{
				m: 3,
				n: 2,
			},
			want: 3,
		},
		{
			name: "dfs2",
			args: args{
				m: 7,
				n: 3,
			},
			want: 28,
		},
		{
			name: "dfs3",
			args: args{
				m: 20,
				n: 20,
			},
			want: 35345263800,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
