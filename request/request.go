package request

import (
	"net/http"
	"net/url"
	"time"

	"codeberg.org/voyna/voyna/constants"
)

func Get(u *url.URL) (*http.Response, error) {
	req := &http.Request{
		URL:    u,
		Header: make(http.Header),
	}
	req.Header.Set("User-Agent", constants.UserAgent)
	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	return client.Do(req)
}
