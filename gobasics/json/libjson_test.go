// Author: xufei
// Date: 2018/12/6

package json

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

type mode struct {
	Name   string `json:"name"`
	Singer string `json:"singer"`
}

func BenchmarkLibJSON(b *testing.B) {
	songs := make([]mode, 0, 100000)

	for i := 0; i < 100000; i++ {
		songs = append(songs, mode{Name: "风继续吹", Singer: "张哥"})
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		data, _ := json.Marshal(songs)

		var songsE []mode
		if err := json.Unmarshal(data, &songsE); err != nil {
			panic(err)
		}
	}
}

var jsonJ = jsoniter.ConfigCompatibleWithStandardLibrary

func BenchmarkJsoniterJSON(b *testing.B) {
	songs := make([]mode, 0, 100000)

	for i := 0; i < 100000; i++ {
		songs = append(songs, mode{Name: "风继续吹", Singer: "张哥"})
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		data, _ := jsonJ.Marshal(songs)

		var songsE []mode
		if err := jsonJ.Unmarshal(data, &songsE); err != nil {
			panic(err)
		}
	}
}

// BenchmarkLibJSON-4        	      10	 119757994 ns/op
// BenchmarkJsoniterJSON-4   	      30	  47837620 ns/op
