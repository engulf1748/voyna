package processor

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"codeberg.org/voyna/voyna/spider"
)

func stringSHA256(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}

func Process(domains []string) {
	ch := make(chan spider.Site, len(domains)) // TODO: figure out efficient channel capacity, if any
	for _, domain := range domains {
		u, err := url.Parse(domain)
		if err != nil {
			// TODO
			continue
		}
		go spider.Crawl(u, ch, 1)
	}

	// create data/database folder if it does not already exist
	err := os.MkdirAll("data/database", 0700)
	if err != nil {
		panic(err)
	}
	for {
		select {
		case site := <-ch:
			b, err := json.Marshal(site)
			if err != nil {
				break
			}
			// we cannot save files with URLs as names, for URLs contain "/" among other "special" characters
			err = os.WriteFile(fmt.Sprintf("data/database/%s", stringSHA256(site.String())), b, 0600)
			if err != nil {
				// TODO
			}
		}
	}
}
