package robotex

import (
	"fmt"
	"net/http"
	"net/url"

	"codeberg.org/voyna/voyna/constants"
	"codeberg.org/voyna/voyna/request"

	"github.com/temoto/robotstxt"
)

func Allowed(u *url.URL) (bool, error) {
	// TODO: allow http
	if !(u.IsAbs() && u.Scheme == "https") {
		return false, fmt.Errorf("non-absolute or non-https link")
	}

	un := *u
	// TODO: learn how to set the path component in an idiomatic way
	un.RawPath, un.Path = "/robots.txt", "/robots.txt"
	un.RawQuery, un.Fragment = "", ""

	resp, err := request.Get(&un)
	var robots *robotstxt.RobotsData
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode != http.StatusNotFound {
			return false, err
		}
		// assuming everything is allowed if there's no robots.txt
		robots, err = robotstxt.FromString("User-agent: *\nDisallow:")
		if err != nil {
			return false, nil
		}
	} else {
		robots, err = robotstxt.FromResponse(resp)
		if err != nil {
			return false, nil
		}
	}
	p := u.EscapedPath()
	return robots.TestAgent(p, constants.UserAgent), nil
}
