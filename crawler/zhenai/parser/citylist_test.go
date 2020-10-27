// Author: xufei
// Date: 2018/11/28

package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, _ := ioutil.ReadFile("testdata/citylist.html")

	type args struct {
		body []byte
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
	}{
		{
			name: "parse city list",
			args: args{
				body: contents,
			},
			wantLen: 470,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseCityList(tt.args.body, ""); len(got.Items) != tt.wantLen {
				t.Errorf("ParseCityList() = %v, want %v", len(got.Items), tt.wantLen)
			}
		})
	}
}
