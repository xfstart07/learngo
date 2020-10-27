// Author: xufei
// Date: 2019/2/27

package app

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"gopkg.in/olivere/elastic.v6"
)

type songBase struct {
	SongCode   string `json:"song_code"`
	Player     string `json:"player"`
	Title      string `json:"title"`
	CreateTime string `json:"create_time"`
	KtvId      int    `json:"ktv_id,omitempty"`
	KtvName    string `json:"ktv_name,omitempty"`
	KtvCode    string `json:"ktv_code,omitempty"`
	RoomId     string `json:"room_id,omitempty"`
}

func CreateWrite(client *elastic.Client, jobs <-chan string) {
	go func() {
		for job := range jobs {
			writer(client, job)
		}
	}()
}

func writer(client *elastic.Client, body string) error {
	songs := []songBase{}
	err := json.Unmarshal([]byte(body), &songs)
	if err != nil {
		log.Println("Unmarshal json err", err)
		return err
	}
	if len(songs) < 3000 {
		songs = append(songs, songs...)
		songs = append(songs, songs...)
	}

	now := time.Now()

	bulk := client.Bulk()
	for song := range songs {
		req := elastic.NewBulkIndexRequest().Index("test_songs").Type("_doc").Doc(song)
		bulk.Add(req)
	}

	_, err = bulk.Do(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("处理单个耗时", time.Since(now))
	return nil
}
