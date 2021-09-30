// Author: xufei
// Date: 2019/2/26

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/olivere/elastic.v6"
)

type songBase struct {
	SongCode   string `json:"song_code"`
	Player     string `json:"player"`
	Title      string `json:"title"`
	CreateTime string `json:"create_time"`
	KtvId      int    `json:"ktv_id"`
	KtvName    string `json:"ktv_name"`
	KtvCode    string `json:"ktv_code"`
	RoomId     string `json:"room_id"`
}

func main() {

	songs := []songBase{}
	for i := 0; i < 3000; i++ {
		songs = append(songs, songBase{
			SongCode:   "$hd632245",
			Player:     "徐子崴 李谷一",
			Title:      "你养我长大我陪你变老",
			CreateTime: "2019-01-22T14:53:40.083912+08:00",
			KtvId:      321,
			KtvName:    "测试版权盒绑定",
			KtvCode:    "K0000320",
			RoomId:     "A01",
		})
	}

	now := time.Now()

	client, _ := newES()
	for i := 0; i < 1000; i++ {
		bulk := client.Bulk()

		for song := range songs {
			req := elastic.NewBulkIndexRequest().Index("test_songs").Type("_doc").Doc(song)
			bulk.Add(req)
		}

		_, err := bulk.Do(context.Background())
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println(time.Since(now))
}

func newES() (*elastic.Client, error) {
	return elastic.NewClient(elastic.SetURL("http://weixinote.dev:9200"),
		elastic.SetSniff(false),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
}
