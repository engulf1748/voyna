package spider

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"

	"codeberg.org/voyna/voyna/site"
)

var ignore = map[string]bool{
	"script": true,
	"style":  true,
	"link":   true,
}

// Processes an html.Node an places data in the passed Site pointer
func processTree(n *html.Node, s *site.Site) {
	var links []*url.URL
	// s.Content
	var b strings.Builder
	// s.Keywords
	var keywords []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode || n.Type == html.DocumentNode {
			if n.Data == "a" {
				for _, v := range n.Attr {
					if v.Key == "href" {
						// TODO: Handle relative URLs
						// TODO: deal with <button>'s being used to the effect of <a>
						u, err := url.Parse(v.Val)
						if err != nil {
							continue
						}
						links = append(links, u)
						break
					}
				}
			} else if n.Data == "meta" {
				m := make(map[string]string) // Why on earth is n.Attr a []string!
				for _, v := range n.Attr {
					m[v.Key] = v.Val
				}
				if m["name"] == "keywords" {
					for _, v := range strings.Split(m["content"], ",") {
						keywords = append(keywords, v)
					}
				}
			} else if n.Data == "title" {
				s.Title = n.FirstChild.Data // assuming nothing is embedded inside <title> except a text node
			} else if ignore[n.Data] {
				return
			} else {
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					f(c)
				}
			}
		} else if n.Type == html.TextNode {
			b.WriteString(n.Data)
		}
	}
	f(n)
	s.Content = b.String()
	s.Links = links
	s.Keywords = keywords
}