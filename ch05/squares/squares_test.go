package main

import "testing"

func squares() func() int {
	var x int

	return func() int {
		defer func() {
			x++
		}()
		return x * x
	}
}

func TestSquares(t *testing.T) {
	x := squares()

	a := x()

	if a != 0 {
		t.Errorf(`a = %#v, want %#v`, a, 0)
	}

	a = x()

	if a != 1 {
		t.Errorf(`a = %#v, want %#v`, a, 1)
	}

	a = x()

	if a != 4 {
		t.Errorf(`a = %#v, want %#v`, a, 4)
	}
}
