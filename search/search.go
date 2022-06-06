package search

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"

	"codeberg.org/voyna/voyna/paths"
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
	f, err := os.Open(paths.CrawlDir())
	defer f.Close()
	if err != nil {
		panic(err)
	}
	dN, err := f.Readdirnames(0)
	if err != nil {
		panic(err)
	}
	for _, n := range dN {
		b, err := os.ReadFile(filepath.Join(paths.CrawlDir(), n))
		if err != nil {
			continue
		}
		var site spider.Site
		err = json.Unmarshal(b, &site)
		if err != nil {
			continue
		}
		var r Result
		if m, c := site.Match(query); m {
			r.Title = site.Title
			r.Context = c
			r.URL = site.URL.String()
			r.Tier = site.Tier
			res = append(res, r)
		}
	}
	// sort according to tier
	sort.Sort(res)
	return res
}
