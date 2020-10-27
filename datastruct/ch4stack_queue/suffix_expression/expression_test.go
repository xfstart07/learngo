// Author: xufei
// Date: 2019-07-29

package suffix_expression

import "testing"

func TestSuffix(t *testing.T) {
	type args struct {
		expressions []string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "后缀表达式",
			args: args{
				expressions: []string{"9", "3", "1", "-", "3", "*", "+", "10", "2", "/", "+"},
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "错误后缀表达式",
			args: args{
				expressions: []string{"9", "3", "1", "-", "-", "-", "3", "*", "+", "10", "2", "/", "+"},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Suffix(tt.args.expressions)
			if (err != nil) != tt.wantErr {
				t.Errorf("Suffix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Suffix() = %v, want %v", got, tt.want)
			}
		})
	}
}
