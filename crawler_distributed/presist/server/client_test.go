// Author: xufei
// Date: 2018/12/18

package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"learngo/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	const host = ":1235"
	go serverRpc(host, "dating_test")
	time.Sleep(1 * time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	profile := model.Profile{
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

	item := engine.Item{
		ID:      profile.ID,
		Payload: profile,
	}

	var result string
	err = client.Call("ItemSaverService.Save", item, &result)
	if err != nil || result != "ok" {
		t.Errorf("err %v, %v", err, result)
	}
}
