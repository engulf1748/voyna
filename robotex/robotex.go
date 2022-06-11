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
	un, err := url.Parse("https://" + u.Host + "/robots.txt") // TODO: do something more sensible
	resp, err := requests.Get(un)
	var robots *robotstxt.RobotsData
	if err != nil {
		return false, err
	}
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
