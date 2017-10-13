package main

import "testing"

const boilingF = 212.0

func TestVarDeclaration(t *testing.T) {
	var f = boilingF
	var c = (f - 32) * 5 / 9

	if f != 212 {
		t.Errorf(`f = %#v, want %#v`, f, 212.0)
	}

	if c != 100 {
		t.Errorf(`c = %#v, want %#v`, c, 100.0)
	}
}
