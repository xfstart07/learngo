// Author: xufei
// Date: 2018/12/12

package persist

import (
	"context"
	"learngo/crawler/model"
	"log"
	"testing"

	"gopkg.in/olivere/elastic.v6"
)

func Test_save(t *testing.T) {
	profile := &model.Profile{
		Name:       "柠檬",
		ID:         "1532019062",
		Marital:    "未婚",
		Age:        31,
		Gender:     "女士",
		Astrology:  "天秤座(09.23-10.22)",
		Height:     163,
		Weight:     45,
		WorkCity:   "北京朝阳区",
		Income:     "2-5万",
		Education:  "大专",
		URL:        "http://album.zhenai.com/u/1532019062",
		Nation:     "汉族",
		Native:     "黑龙江哈尔滨",
		IsBuyHouse: "已购房",
		IsBuyCar:   "已买车",
	}

	type args struct {
		item interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test elasticsearch Save",
			args: args{
				item: profile,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			profile, ok := tt.args.item.(*model.Profile)
			if !ok {
				t.Error("item not profile")
			}
			client, err := elastic.NewClient(elastic.SetSniff(false))
			if err != nil {
				panic(err)
			}

			resp, err := client.Index().
				Index("dating_profile").
				Type("zhenai_test").
				Id(profile.ID).
				BodyJson(profile).
				Do(context.Background())
			if err != nil {
				panic(err)
			}

			log.Printf("%+v", resp)
		})
	}
}
