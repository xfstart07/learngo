// Author: xufei
// Date: 2019/1/2

package service

import (
	"context"
	"encoding/json"
	"log"
	"time"

	elastic "gopkg.in/olivere/elastic.v6"
)

type SingData struct {
	SongID     string `json:"song_id"`
	SongName   string `json:"song_name"`
	SongPlayer string `json:"song_player"`
	KtvID      int    `json:"ktv_id"`
	Count      int    `json:"count"`
	DateAt     string `json:"date_at"`
}

func Search(client *elastic.Client) {
	search := client.Search().Index("dating_songs").Type("songs")

	// query
	q := elastic.NewRangeQuery("date_at").Gte("2018-12-01")
	search = search.Query(q)

	// aggs
	agg := elastic.NewTermsAggregation().Field("date_at")
	agg.Order("_term", false)
	agg.Size(365)

	// 获取需要的字段
	songSource := elastic.NewFetchSourceContext(true).Include("song_name", "song_player", "ktv_id")
	songNameHitAgg := elastic.NewTopHitsAggregation().FetchSourceContext(songSource).Size(1)

	// 计算播放量
	sumAgg := elastic.NewSumAggregation().Field("count")

	// 根据歌曲ID
	songsAgg := elastic.NewTermsAggregation().Field("song_id")
	songsAgg.Size(1000)

	// 歌曲
	songsAgg.SubAggregation("sum_count", sumAgg)
	songsAgg.SubAggregation("song_name_hit", songNameHitAgg)
	agg.SubAggregation("songs", songsAgg)

	search = search.Aggregation("dates", agg)

	search.Size(0)

	sub := time.Now()

	result, err := search.Do(context.Background())
	if err != nil {
		log.Printf("err %v", err)
		return
	}

	since := time.Since(sub)
	log.Printf("耗时 %d, %d, %d ", result.TookInMillis, since/time.Second, result.TotalHits())

	sings, err := aggHandle(*result.Aggregations["dates"])
	if err != nil {
		log.Println("err", err)
	}
	//log.Println(sings)
	log.Println(len(sings))
	//log.Println(string(*result.Aggregations["dates"]))
}

type dateBucketResult struct {
	Buckets []dateBucket `json:"buckets"`
}

type dateBucket struct {
	Key         int64            `json:"key"`
	KeyAsString string           `json:"key_as_string"`
	Songs       songBucketResult `json:"songs"`
}

type songBucketResult struct {
	Buckets []songBucket `json:"buckets"`
}

type songBucket struct {
	Key         string             `json:"key"`
	SumCount    sumValue           `json:"sum_count"`
	SongNameHit songNameHitsResult `json:"song_name_hit"`
}

type sumValue struct {
	Value float64 `json:"value"`
}

type songNameHitsResult struct {
	Hits songNameHitResult `json:"hits"`
}

type songNameHitResult struct {
	Hits []songNameHit `json:"hits"`
}

type songNameHit struct {
	Source SingData `json:"_source"`
}

func aggHandle(data []byte) ([]SingData, error) {
	var sings []SingData

	var result dateBucketResult
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	for _, bucket := range result.Buckets {
		sing := SingData{}
		sing.DateAt = bucket.KeyAsString[:10]
		log.Println("日期", bucket.KeyAsString, bucket.KeyAsString[:10])

		log.Println("歌曲总数", len(bucket.Songs.Buckets))
		for _, songBucket := range bucket.Songs.Buckets {
			sing.SongID = songBucket.Key
			sing.Count = int(songBucket.SumCount.Value)

			source := songBucket.SongNameHit.Hits.Hits[0].Source
			sing.SongName = source.SongName
			sing.SongPlayer = source.SongPlayer
			sing.KtvID = source.KtvID

			sings = append(sings, sing)
		}
	}

	return sings, nil
}
