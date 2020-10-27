// Author: xufei
// Date: 2018/11/24

package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var client = http.Client{}

// 限速
var rateLimit = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimit
	if url == "" {
		return nil, fmt.Errorf("url is blank")
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error Status %d, url: %s", resp.StatusCode, url)
	}

	return ioutil.ReadAll(resp.Body)
}
