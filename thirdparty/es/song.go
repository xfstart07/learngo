// Author: xufei
// Date: 2018/12/28

package main

import (
	"context"
	"fmt"
	"learngo/thirdparty/es/service"
	"log"
	"math/rand"
	"time"

	"gopkg.in/olivere/elastic.v6"
)

func main() {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	service.Search(client)

	//bulkInsert(client)
}

// 批量生成数据
func bulkInsert(client *elastic.Client) {
	var songsIDs []string
	for i := 1000; i < 1500; i++ {
		songsIDs = append(songsIDs, fmt.Sprintf("a123%d", i))
	}

	var ktvIDs []int
	for i := 1; i < 100; i++ {
		ktvIDs = append(ktvIDs, i)
	}

	var dates []string
	for i := 0; i < 100; i++ {
		delta := i * 24
		dur, _ := time.ParseDuration(fmt.Sprintf("-%dh", delta))

		dates = append(dates, time.Now().Add(dur).Format("2006-01-02"))
	}

	sub := time.Now()
	for _, song := range songsIDs {
		bulk := client.Bulk()

		for _, date := range dates {
			for j := 0; j < 50; j++ {
				kI := rand.Intn(99)

				song := service.SingData{
					SongID:     song,
					SongPlayer: fmt.Sprintf("刘欢%s", song),
					SongName:   fmt.Sprintf("大河向东流%s", song),
					KtvID:      ktvIDs[kI],
					Count:      rand.Intn(50),
					DateAt:     date,
				}

				req := elastic.NewBulkIndexRequest().Index("dating_songs").Type("songs").Doc(song)
				bulk.Add(req)
			}
		}

		_, err := bulk.Do(context.Background())
		if err != nil {
			log.Printf("bulk err %v", err)
		}
	}

	dur := time.Since(sub)
	log.Println("耗时", dur/time.Second)
}
