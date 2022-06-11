package robotex

import (
	"net/url"
	"testing"
)

type object struct {
	u       *url.URL
	allowed bool
}

func TestAllowed(t *testing.T) {
	var objects []object
	u1, _ := url.Parse("https://sr.ht")
	objects = append(objects, object{u1, true})
	u2, _ := url.Parse("https://sr.ht/metrics")
	objects = append(objects, object{u2, false})
	u3, _ := url.Parse("https://codeberg.org/ar324/gofe")
	objects = append(objects, object{u3, true})
	u4, _ := url.Parse("https://codeberg.org/ar324/gofe/commit/cbea66088703398043cb6dab45edfa917134a9c4")
	objects = append(objects, object{u4, false})

	for _, object := range objects {
		y, err := Allowed(object.u)
		if err != nil {
			t.Fatalf("err on %v: %v", object.u, err)
		}
		if object.allowed != y {
			t.Fatalf("mismatch on %v: expected: %v; found: %v", object.u, object.allowed, y)
		} else {
			t.Logf("success on %v: expected: %v; found: %v", object.u, object.allowed, y)
		}
	}
}
