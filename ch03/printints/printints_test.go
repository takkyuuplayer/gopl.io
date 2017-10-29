package printints

import "testing"

func TestIntsToString(t *testing.T) {
	if intsToString([]int{1, 23, 456, 7890}) != "[1, 23, 456, 7890]" {
		t.Errorf(`intsToString([]int{1,23,456,7890}) = %#v, want %#v`, intsToString([]int{1, 23, 456, 7890}), "[1, 23, 456, 7890]")
	}
}
