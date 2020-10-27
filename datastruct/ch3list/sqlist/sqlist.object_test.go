// Author: xufei
// Date: 2019-07-25

package sqlist

import "testing"

func TestSqList_GetElem(t *testing.T) {
	type fields struct {
		data   [MAXSIZE]int
		length int
	}
	type args struct {
		i int
		e *int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "获取元素",
			fields: fields{
				data:   [MAXSIZE]int{1, 2, 3, 4, 5},
				length: 5,
			},
			args: args{
				i: 3,
				e: new(int),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &SqList{
				data:   tt.fields.data,
				length: tt.fields.length,
			}
			if got := list.GetElem(tt.args.i, tt.args.e); got != tt.want {
				t.Errorf("SqList.GetElem() = %v, want %v", got, tt.want)
			} else {
				t.Log(*tt.args.e)
			}
		})
	}
}
