package search

import (
	//"encoding/json"
	//"os"
	//"path/filepath"
	"sort"

	//"codeberg.org/voyna/voyna/paths"
	//"codeberg.org/voyna/voyna/site"
	"codeberg.org/voyna/voyna/spider"
)

type Result struct {
	URL     string
	Context string
	Title   string
	Tier    int
}

type Results []Result

func (r Results) Less(i, j int) bool {
	return r[i].Tier < r[j].Tier
}

func (r Results) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Results) Len() int {
	return len(r)
}

func Search(query string) Results {
	var res Results
	spider.DB.RLock()
	defer spider.DB.RUnlock()
	for _, s := range spider.DB.M {
		var r Result
		if m, c := s.Match(query); m {
			r.Title = s.Title
			r.Context = c
			r.URL = s.URL.String()
			r.Tier = s.Tier
			res = append(res, r)
		}
	}
	// sort according to tier
	sort.Sort(res)
	return res
}
