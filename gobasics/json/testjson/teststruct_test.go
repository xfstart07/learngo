package testjson

import (
	"encoding/json"
	"testing"
)

func TestJson(t *testing.T) {
	tmp := &TestStruct{
		Uuid:   "xxxxuuid",
		Seckey: "seckeyxxxx",
	}

	body, _ := json.Marshal(tmp)
	t.Logf("marshal: %v", string(body))

	t.Logf("toString: %v", tmp)
}
