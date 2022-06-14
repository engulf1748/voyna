package spider

import (
	// "fmt"
	// "log"
	"net/url"
	// "os"
	// "runtime"
	// "strings"
	"sync"
	"time"

	"codeberg.org/voyna/voyna/log4j"
	"codeberg.org/voyna/voyna/request"
	"codeberg.org/voyna/voyna/robotex"
	"codeberg.org/voyna/voyna/site"

	"golang.org/x/net/html"
)

const (
	MaxDepth  = 3
	RateLimit = 10000
)

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

// Keeps track of the URLs we've seen so far.
var seen safeSeen

// Keep track of the hosts we've seen so far.
var hseen hostSeen

var rateLimitCh chan bool

func init() {
	seen.s = make(map[string]bool)
	hseen.s = make(map[string]int)

	rateLimitCh = make(chan bool, RateLimit)
}

func Crawl(u *url.URL, ch chan site.Site, tier int) {
	rateLimitCh <- true
	defer func() {
		<-rateLimitCh
	}()

	if !(u.IsAbs() && u.Scheme == "https") {
		log4j.Logger.Printf("ignoring %q; not HTTPS or absolute link", u.String())
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
	if tier > MaxDepth {
		log4j.Logger.Printf("ignoring %q; tier: %q", u.String(), tier)
		return
	}

	// TODO: Handle cases such as xyz.com/page and xyz.com/page/
	surl := u.String()
	// check if "u" was already processed
	seen.Lock()

	if seen.s[surl] == true {
		seen.Unlock()
		return
	}
	seen.s[surl] = true
	seen.Unlock()

	allowed, err := robotex.Allowed(u)
	if err != nil {
		log4j.Logger.Printf("Allowed failed for %q: %v; moving on . . .\n", u, err)
		return
	}
	if !allowed {
		log4j.Logger.Printf("robots.txt disallowed crawling %q; moving on . . .\n", u)
		return
	}

	resp, err := request.Get(u)
	if err != nil || resp.StatusCode != 200 {
		log4j.Logger.Printf("GET failed (or returned non-200) for %q: %v; moving on . . .\n", u, err)
		return
	}
	defer resp.Body.Close()

	n, err := html.Parse(resp.Body)
	if err != nil {
		log4j.Logger.Printf("html.Parse failed for %s: %v; moving on . . .\n", u, err)
		return
	}

	s := site.Site{
		URL:       u,
		Tier:      tier,
		IndexTime: time.Now(),
	}

	processTree(n, &s)

	for _, u := range s.Links {
		// TODO: crawl relative links too
		go Crawl(u, ch, tier+1)
	}

	ch <- s
}
