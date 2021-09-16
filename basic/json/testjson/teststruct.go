package testjson

import "encoding/json"

type TestStruct struct {
	Seckey string `json:"seckey"`
	Uuid   string `json:"uuid"`
}

func (s *TestStruct) Print() string {
	t := TestStruct{
		Uuid: s.Uuid,
	}

	b, _ := json.Marshal(t)
	return string(b)
}

func (s *TestStruct) String() string {
	t := TestStruct{
		Uuid: s.Uuid,
	}

	b, _ := json.Marshal(t)
	return string(b)
}
