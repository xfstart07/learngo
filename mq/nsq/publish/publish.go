// Author: xufei
// Date: 2019-10-15 16:13

package publish

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/nsqio/go-nsq"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

type record struct {
	KtvID         string `json:"ktv_id"`
	RoomIP        string `json:"room_ip"`
	EndAt         string `json:"end_at"`         // 结束播放时间
	Duration      int    `json:"duration"`       // 实际播放时长
	TotalDuration int    `json:"total_duration"` // 文件总时长，秒
	SongID        string `json:"song_id"`
	SongName      string `json:"song_name"`
	Singer        string `json:"singer"`
}

func NewProduct(songID *int32) error {
	topicName := "publish"
	msgCount := 300

	config := nsq.NewConfig()
	config.LookupdPollInterval = 10 * time.Second
	producer, err := nsq.NewProducer("weixinote.dev:4150", config)
	if err != nil {
		return err
	}
	defer producer.Stop()

	var testData [][]byte
	for i := 0; i < msgCount; i++ {
		si := strconv.Itoa(i)
		atomic.AddInt32(songID, 1)
		songIDstr := fmt.Sprintf("%d", *songID)
		rec := record{
			KtvID:         "K123213" + si,
			RoomIP:        "A11" + si,
			EndAt:         time.Now().Format("2006-01-02 15:04:05"),
			Duration:      rand.Intn(300),
			TotalDuration: 300,
			SongID:        "SY" + songIDstr,
			SongName:      "十年" + si,
			Singer:        "chenyixun" + si,
		}
		bytes, _ := json.Marshal(rec)

		testData = append(testData, bytes)
	}

	err = producer.MultiPublish(topicName, testData)
	//for _, data := range testData {
	//	err = producer.Publish(topicName, data)
	//}
	if err != nil {
		return err
	}

	return nil
}
