// Author: xufei
// Date: 2019-07-29

package infix2suffix

import (
	"reflect"
	"testing"
)

func TestInfix2Suffix(t *testing.T) {
	type args struct {
		exps []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "中缀转后缀",
			args: args{
				exps: []string{"9", "+", "(", "3", "-", "1", ")", "*", "3", "+", "10", "/", "2"},
			},
			want:    []string{"9", "3", "1", "-", "3", "*", "+", "10", "2", "/", "+"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exp := New(tt.args.exps)

			got, err := exp.Infix2Suffix()
			if (err != nil) != tt.wantErr {
				t.Errorf("Infix2Suffix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Infix2Suffix() = %v, want %v", got, tt.want)
			}
		})
	}
}
