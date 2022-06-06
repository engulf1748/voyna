package spider

import (
	// "fmt"
	"log"
	"net/http"
	"net/url"
	// "os"
	"strings"
	"sync"
	"time"

	"codeberg.org/voyna/voyna/site"

	"golang.org/x/net/html"
	// "github.com/temoto/robotstxt"
)

const MaxDepth = 3

type safeSeen struct {
	s map[string]bool
	// struct embedding--nice!
	sync.Mutex
}

// unlike safeSeen, this deals only with hostnames: we need to prevent tier overriding
type hostSeen struct {
	s map[string]int
	sync.Mutex
}

var seen safeSeen
var hseen hostSeen

func init() {
	seen.s = make(map[string]bool)
	hseen.s = make(map[string]int)
}

// TODO: move to separate file/package
func Get(u *url.URL) (*http.Response, error) {
	req := &http.Request{
		URL:    u,
		Header: make(http.Header),
	}
	req.Header.Set("User-Agent", "GofëBot")
	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	return client.Do(req)
}

func Crawl(u *url.URL, ch chan site.Site, tier int) {
	if !(u.IsAbs() && u.Scheme == "https") {
		return
	}

	// "actual" tier, based on just host-name
	hseen.Lock()
	if t := hseen.s[u.Host]; t != 0 {
		if t < tier {
			tier = t
		} else {
			hseen.s[u.Host] = tier
		}
	} else {
		hseen.s[u.Host] = tier
	}
	hseen.Unlock()

	// TODO: Handle this better
	if tier > 3 {
		return
	}

	var s site.Site

	// codeberg.org/ and codeberg.org are the same, even though their url.URL representations might be different
	domain := strings.TrimSuffix(u.String(), "/")
	// check if "u" was already processed
	seen.Lock()

	if seen.s[domain] == true {
		seen.Unlock()
		return
	}
	seen.s[domain] = true
	seen.Unlock()

	s.URL = u
	s.Tier = tier
	s.IndexTime = time.Now()

	resp, err := Get(u)
	if err != nil || resp.StatusCode != 200 {
		log.Printf("GET failed (or returned non-200) for %s: %v; skipping . . .\n", domain, err)
		return
	}
	defer resp.Body.Close()

	n, err := html.Parse(resp.Body)
	if err != nil {
		log.Printf("html.Parse failed for %s: %v; skipping . . .\n", domain, err)
		return
	}

	processTree(n, &s)

	for _, u := range s.Links {
		// TODO: crawl relative links too
		go Crawl(u, ch, tier+1)
	}

	ch <- s
}
