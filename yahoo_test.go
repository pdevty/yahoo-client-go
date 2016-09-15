package yahoo

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if h := req.Header.Get("User-Agent"); h != defaultUserAgent+"dummy-appid" {
			t.Errorf("User-Agent shoud be '%s' but %s", defaultUserAgent+"dummy-appid", h)
		}
	}))
	defer ts.Close()

	client, _ := NewClientWithOptions("dummy-appid", ts.URL, false)

	req, _ := http.NewRequest("GET", client.urlFor("/").String(), nil)
	client.Request(req)
}
