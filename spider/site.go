package spider

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

// Represents a website as seen by spider.
type Site struct {
	URL   *url.URL
	Links []*url.URL

	// We will confirm this ourselves during requests.
	// I think it's a helpful indicator for the front-end,
	// should they choose to accentuate this fact.
	Secure bool

	// Will this be useful? I was thinking adding
	// both invalid links or broken links to this.
	// We could also separate these.
	BrokenLinks []string

	// This will later be updated, I assume, to a
	// []string instead. The tokenizer will probably
	// parse words individually like that, maybe.
	// For now, this is fine.
	Content string

	Title string

	Tier int

	// it should be difficult to abuse keywords in a tier-based system
	Keywords []string

	References int

	// To keep track of page freshness
	IndexTime time.Time
}

// Allows printing a Site in an easy-to-read manner
func (s Site) String() string {
	// TODO: add other fields
	return fmt.Sprintf("Keywords: %v\nTitle: %s\n", s.Keywords, s.Title)
}

// Performs simple (and verging on stupid) keyword matching
func (s *Site) Match(query string) (bool, string) {
	query = strings.ToLower(query)
	qwords := strings.Split(query, " ")
	// if s.Title contains the query, let it take precedence
	words := strings.Split(strings.ToLower(s.Title), " ")
	for _, word := range words {
		for _, qword := range qwords {
			if word == qword {
				return true, s.Title
			}
		}
	}

	words = strings.Split(s.Content, " ")
	contextLevel := 3
	for i, word := range words {
		for _, qword := range qwords {
			if word == qword {
				var low, high int
				if low = i - contextLevel; low < 0 {
					low = 0
				}
				if high = i + contextLevel; high > len(words) {
					high = len(words)
				}
				return true, strings.Join(words[low:high], " ")
			}
		}
	}

	for _, v := range s.Keywords {
		for _, qword := range qwords {
			if v == qword {
				return true, v
			}
		}
	}

	return false, ""
}
