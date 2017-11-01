package ch04

import (
	"reflect"
	"testing"
)

func uniq05(src []string) []string {
	idx := 0
	for _, s := range src {
		if src[idx] == s {
			continue
		}
		idx++
		src[idx] = s
	}

	return src[:idx+1]
}

func TestUniq(t *testing.T) {
	src := []string{"", "a", "b", "b", "c", "c", "c", "b"}

	ret := uniq05(src)

	// in-place
	if !reflect.DeepEqual(src, []string{"", "a", "b", "c", "b", "c", "c", "b"}) {
		t.Errorf(`src = %#v, want %#v`, src, []string{"", "a", "b", "c", "b", "c", "c", "b"})
	}

	// unique
	if !reflect.DeepEqual(ret, []string{"", "a", "b", "c", "b"}) {
		t.Errorf(`ret= %#v, want %#v`, src, []string{"", "a", "b", "c", "b"})
	}
}
