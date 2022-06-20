package constants

import (
	"errors"
)

const (
	UserAgent = "VoynaBot"
)

var ErrNilURL = errors.New("nil *url.URL")
