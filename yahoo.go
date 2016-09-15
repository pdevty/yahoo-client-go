package yahoo

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

const (
	defaultBaseURL    = "http://shopping.yahooapis.jp"
	defaultUserAgent  = "Yahoo AppID: "
	apiRequestTimeout = 30 * time.Second
)

// Client api client for yahoo
type Client struct {
	BaseURL   *url.URL
	AppID     string
	Verbose   bool
	UserAgent string
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

// NewClient returns new yahoo.Client
func NewClient(appid string) *Client {
	u, _ := url.Parse(defaultBaseURL)
	return &Client{u, appid, false, defaultUserAgent + appid}
}

// NewClientWithOptions returns new yahoo.Client
func NewClientWithOptions(appid string, rawurl string, verbose bool) (*Client, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	return &Client{u, appid, verbose, defaultUserAgent + appid}, nil
}

func (c *Client) urlFor(path string) *url.URL {
	newURL, err := url.Parse(c.BaseURL.String())
	if err != nil {
		panic("invalid url passed")
	}

	newURL.Path = path

	return newURL
}

// Request request to yahoo and receive response
func (c *Client) Request(req *http.Request) (resp *http.Response, err error) {
	req.Header.Set("User-Agent", c.UserAgent)

	if c.Verbose {
		dump, err := httputil.DumpRequest(req, true)
		if err == nil {
			log.Printf("%s", dump)
		}
	}

	client := &http.Client{} // same as http.DefaultClient
	client.Timeout = apiRequestTimeout
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	if c.Verbose {
		dump, err := httputil.DumpResponse(resp, true)
		if err == nil {
			log.Printf("%s", dump)
		}
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return resp, fmt.Errorf("API result failed: %s", resp.Status)
	}
	return resp, nil
}

func closeResponse(resp *http.Response) {
	if resp != nil {
		resp.Body.Close()
	}
}
