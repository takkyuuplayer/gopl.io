package main

import (
	"reflect"
	"testing"
)

func rotate(src []int, cnt int) {
	tmp := make([]int, cnt)
	copy(tmp, src[:cnt])

	copy(src, src[cnt:])
	copy(src[len(src)-cnt:], tmp)
}

func TestRotate(t *testing.T) {
	s := []int{0, 1, 2, 3}

	rotate(s, 0)

	if !reflect.DeepEqual(s, []int{0, 1, 2, 3}) {
		t.Errorf(`s = %#v, want %#v`, s, []int{0, 1, 2, 3})
	}

	rotate(s, 1)

	if !reflect.DeepEqual(s, []int{1, 2, 3, 0}) {
		t.Errorf(`s = %#v, want %#v`, s, []int{1, 2, 3, 0})
	}

	rotate(s, 2)

	if !reflect.DeepEqual(s, []int{3, 0, 1, 2}) {
		t.Errorf(`s = %#v, want %#v`, s, []int{3, 0, 1, 2})
	}

}
