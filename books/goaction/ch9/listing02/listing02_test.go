package listing02_test

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotx = "\u2717"

func TestDownload(t *testing.T) {
	var urls = []struct {
		url        string
		statusCode int
	}{
		{
			url:        "http://www.goinggo.net/feeds/posts/default?alt=rss",
			statusCode: 404,
		},
		{
			url:        "http://zjjzzgl.zjsgat.gov.cn:9090/zahlw/index",
			statusCode: 200,
		},
	}

	for _, u := range urls {
		t.Logf("\tWhen checking %s for status code %d", u.url, u.statusCode)
		{
			resp, err := http.Get(u.url)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.", ballotx, err)
			}
			t.Log("\t\tShould be able to make the Get call checkMark")
			defer resp.Body.Close()

			if resp.StatusCode == u.statusCode {
				t.Logf("\t\tShould receive a %d status.", u.statusCode)
			} else {
				t.Errorf("\t\tShould receive a %d status. %v", u.statusCode, resp.StatusCode)
			}
		}
	}

}
