package values

import (
	"fmt"
	"net/url"
)

func UsageValues() {
	v := url.Values{}
	v.Set("action", "1")
	v.Add("video_url", "aa.mp4")
	v.Add("qr_url", "12312312312")
	v.Add("position", "203,293,233,233")

	fmt.Println(v.Encode())
}
