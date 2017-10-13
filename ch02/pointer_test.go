package main

import "testing"

func TestPointer(t *testing.T) {
	x := 1
	p := &x

	if *p != 1 {
		t.Errorf(`*p = %#v, want %#v`, *p, 1)
	}

	*p = 2

	if x != 2 {
		t.Errorf(`p = %#v, want %#v`, p, 2)
	}
}

func Test2DifferentPointers(t *testing.T) {
	var x, y int

	if &x != &x {
		t.Errorf(`&x = %#v, want %#v`, &x, &x)
	}

	if &x == &y {
		t.Errorf(`&x = %#v, do NOT want %#v`, &x, &y)
	}

	if &x == nil {
		t.Errorf(`&x = %#v, do NOT want %#v`, &x, nil)
	}
}

func TestReturnPointerOfLocalVariable(t *testing.T) {
	f := func() *int {
		v := 1

		return &v
	}

	if *f() != 1 {
		t.Errorf(`*f() = %#v, want %#v`, *f(), 1)
	}

	if f() == f() {
		t.Errorf(`f() = %#v, do NOT want %#v`, f(), f())
	}
}

func TestSideEffect(t *testing.T) {
	incr := func(p *int) int {
		*p++
		return *p
	}

	v := 1

	if incr(&v) != 2 {
		t.Errorf(`incr(v) = %#v, want %#v`, 2, 2)
	}

	if v != 2 {
		t.Errorf(`v = %#v, want %#v`, v, 2)
	}
}
