// Author: xufei
// Date: 2019-10-17 10:25

package consumer

import (
	"context"
	"fmt"
	"log"
	"time"

	"gopkg.in/olivere/elastic.v6"
)

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

type work struct {
	list []record
	es   *elastic.Client
}

func newWorker(in <-chan record) *work {
	w := &work{
		list: make([]record, 0, 1000),
		es:   newES(),
	}

	go func() {
		tick := time.NewTicker(1 * time.Minute)
		for {
			select {
			case r := <-in:
				w.list = append(w.list, r)
				if len(w.list) == 1000 {
					w.store(w.list)
					w.list = w.list[:0]
					log.Println("写入成功！")
				}
			case <-tick.C:
				log.Println("长度", len(w.list))
				w.store(w.list)
				w.list = w.list[:0]
			}
		}
	}()

	return w
}

func (w *work) store(list []record) {
	if len(list) == 0 {
		return
	}
	var err error

	bulk := w.es.Bulk()
	for _, record := range list {
		var req *elastic.BulkIndexRequest
		req = elastic.NewBulkIndexRequest().Index("test_nsq_records").Type("_doc").Doc(record)

		bulk.Add(req)
	}

	resp, err := bulk.Do(context.Background())
	if err != nil {
		log.Println("写入失败", err)
		return
	}

	if resp.Errors {
		for _, fail := range resp.Failed() {
			err = fmt.Errorf("%v, %w", fail.Error.Reason, err)
		}
	}

	if err != nil {
		log.Println("写入失败", err)
	}
}
