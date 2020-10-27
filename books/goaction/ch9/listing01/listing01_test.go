package listing01

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotx = "\u2717"

func TestDownload(t *testing.T) {
	url := "http://www.goinggo.net/feeds/posts/default?alt=rss"
	statusCode := 200

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusCode)

		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.", ballotx, err)
			}
			t.Log("\t\tShould be able to make the Get call checkMark")
			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a %d status.", statusCode)
			} else {
				t.Errorf("\t\tShould receive a %d status. %v", statusCode, resp.StatusCode)
			}
		}
	}
}
