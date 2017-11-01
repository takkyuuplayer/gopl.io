package ch04

import (
	"reflect"
	"testing"
)

func reverse03(src *[5]int) {
	for i, j := 0, len(src)-1; i < j; i, j = i+1, j-1 {
		src[i], src[j] = src[j], src[i]
	}
}

func TestReverse(t *testing.T) {
	src := [...]int{1, 2, 3, 4, 5}
	reverse03(&src)

	if !reflect.DeepEqual(src, [...]int{5, 4, 3, 2, 1}) {
		t.Errorf(`src = %#v, want %#v`, src, [...]int{5, 4, 3, 2, 1})
	}
}
