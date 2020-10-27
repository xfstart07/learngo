// Author: xufei
// Date: 2018/11/23

package parser

import (
	"io/ioutil"
	"learngo/crawler/model"
	"testing"
)

func TestReAge(t *testing.T) {
	profile := &model.Profile{}
	profile.Age = ReInt("27岁")
	t.Log(profile)
}

func TestParseProfile(t *testing.T) {
	contents, _ := ioutil.ReadFile("testdata/profile.html")

	type args struct {
		body   []byte
		url    string
		name   string
		gender string
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
	}{
		{
			name: "parse profile",
			args: args{
				body:   contents,
				url:    "http://album.zhenai.com/u/1777694127",
				name:   "Xiao哟哟",
				gender: "女士",
			},
			wantLen: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseProfile(tt.args.body, tt.args.url, tt.args.name, tt.args.gender); len(got.Items) != tt.wantLen {
				t.Errorf("parseProfile() = %v, want %v", len(got.Items), tt.wantLen)
			}
		})
	}
}
