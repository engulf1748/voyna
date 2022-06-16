package translator

import (
	"codeberg.org/ar324/gofe/api"
	"codeberg.org/voyna/voyna/search"
)

func Search(query string) api.Result {
	rs := search.Search(query)
	result := api.Result{Links: []*api.Link{}}
	for _, r := range rs {
		result.Links = append(result.Links, &api.Link{Desc: r.Title, URL: r.URL, Context: r.Context})
	}
	return result
}
