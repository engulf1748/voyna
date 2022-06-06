package processor

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/url"
	"os"
	"path/filepath"

	"codeberg.org/voyna/voyna/paths"
	"codeberg.org/voyna/voyna/site"
	"codeberg.org/voyna/voyna/spider"
)

func stringSHA256(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}

func Process(domains []string) {
	ch := make(chan site.Site, len(domains)) // TODO: figure out efficient channel capacity, if any
	for _, domain := range domains {
		u, err := url.Parse(domain)
		if err != nil {
			// TODO
			continue
		}
		go spider.Crawl(u, ch, 1)
	}

	// create storage folder if it does not already exist
	err := os.MkdirAll(paths.CrawlDir(), 0700)
	if err != nil {
		panic(err)
	}
	for {
		select {
		case s := <-ch:
			b, err := json.Marshal(s)
			if err != nil {
				break
			}
			// we cannot save files with URLs as names, for URLs contain "/" among other "special" characters
			fN := filepath.Join(paths.CrawlDir(), stringSHA256(s.String()))
			err = os.WriteFile(fN, b, 0600)
			if err != nil {
				// TODO
			}
		}
	}
}
