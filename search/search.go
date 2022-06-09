package search

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"

	"codeberg.org/voyna/voyna/paths"
	"codeberg.org/voyna/voyna/site"
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
	f, err := os.Open(paths.CorpusDir)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	dN, err := f.Readdirnames(0)
	if err != nil {
		panic(err)
	}
	for _, n := range dN {
		b, err := os.ReadFile(filepath.Join(paths.CorpusDir, n))
		if err != nil {
			continue
		}
		var s site.Site
		err = json.Unmarshal(b, &s)
		if err != nil {
			continue
		}
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
