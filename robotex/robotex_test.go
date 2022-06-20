package robotex

import (
	"net/url"
	"testing"
)

func TestAllowed(t *testing.T) {
	testcases := []struct {
		u       *url.URL
		allowed bool
	}{
		{parse("https://sr.ht"), true},
		{parse("https://sr.ht/metrics"), false},
		{parse("https://codeberg.org/ar324/gofe"), true},
		{parse("https://codeberg.org/ar324/gofe/commit/cbea66088703398043cb6dab45edfa917134a9c4"), false},
	}

	for _, test := range testcases {
		y, err := Allowed(test.u)
		if err != nil {
			t.Errorf("err on %v: %v", test.u, err)
		}
		if test.allowed != y {
			t.Errorf("mismatch on %v: expected: %v; found: %v", test.u, test.allowed, y)
		} else {
			t.Logf("success on %v: expected: %v; found: %v", test.u, test.allowed, y)
		}
	}
}

func parse(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		return nil
	}
	return u
}
