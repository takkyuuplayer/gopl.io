package basename

import "testing"

func TestBasename(t *testing.T) {
	if basename("a") != "a" {
		t.Errorf(`basename("a") = %#v, want %#v`, basename("a"), "a")
	}

	if basename("a.go") != "a" {
		t.Errorf(`basename("a.go") = %#v, want %#v`, basename("a.go"), "a")
	}

	if basename("a/b/c.go") != "c" {
		t.Errorf(`basename("a/b/c.go") = %#v, want %#v`, basename("a/b/c.go"), "c")
	}

	if basename("a/b.c.go") != "b.c" {
		t.Errorf(`basename("a/b.c.go") = %#v, want %#v`, basename("a/b.c.go"), "b.c")
	}
}
