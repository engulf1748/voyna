package spider

import (
	"errors"
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

var ErrNilNode = errors.New("nil *html.Node")

// Processes an html.Node an places data in the passed Site pointer. Ensure s.URL is filled in.
func processTree(n *html.Node, s *site.Site) error {
	if n == nil {
		return ErrNilNode
	}
	var links []*url.URL
	// s.Content
	var b strings.Builder
	// s.Keywords
	var keywords []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n == nil {
			return
		}
		if n.Type == html.ElementNode || n.Type == html.DocumentNode {
			if n.Data == "a" {
				for _, v := range n.Attr {
					if v.Key == "href" {
						// TODO: deal with <button>'s being used to the effect of <a>
						u, err := url.Parse(v.Val)
						if err != nil {
							continue
						}
						if u.Path == "" { // fragments, I suppose; TODO: what to do about this?
							continue
						}
						u = s.URL.ResolveReference(u) // returns a copy of u if it is an absolute URL
						links = append(links, u)
						break
					}
				}
			} else if n.Data == "meta" {
				m := make(map[string]string) // Why on earth is n.Attr a []Attribute!
				for _, v := range n.Attr {
					m[v.Key] = v.Val
				}
				if m["name"] == "keywords" {
					for _, v := range strings.Split(m["content"], ",") {
						keywords = append(keywords, v)
					}
				} else if m["name"] == "description" {
					s.MetaDescription = m["content"]
				}
			} else if n.Data == "title" {
				if n.FirstChild == nil {
					// TODO: panic: runtime error: invalid memory address or nil pointer dereference
					return
				}
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
	return nil
}
