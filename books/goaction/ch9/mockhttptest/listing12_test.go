package mockhttptest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const checkMark = "\u2713"
const ballotx = "\u2717"

var feed = `<?xml version="1.0" encoding="UTF-8"? >
<rss>
<channel>
  <title>Going Go Programming</title>
</channel>
</rss>
`

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")

		fmt.Fprintln(w, feed)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

func TestDownload(t *testing.T) {
	statusCode := http.StatusOK

	server := mockServer()
	defer server.Close()

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", server.URL, statusCode)

		{
			resp, err := http.Get(server.URL)
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
