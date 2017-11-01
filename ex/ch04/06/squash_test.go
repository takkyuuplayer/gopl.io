package main

import (
	"reflect"
	"testing"
	"unicode"
	"unicode/utf8"
)

func squash(org []byte) []byte {
	s := string(org)

	idx := 0
	for _, v := range s {
		if unicode.IsSpace(v) {
			copy(org[idx:idx+1], []byte(" "))
			idx++
		} else {
			size := utf8.RuneLen(v)
			copy(org[idx:idx+size], []byte(string(v)))
			idx += size
		}
	}

	return org[:idx]
}

func TestSquashForAscii(t *testing.T) {
	a := []byte("a\nb\tc d")
	b := squash(a)

	if string(a) != "a b c d" {
		t.Errorf(`a = %#v, want %#v`, string(a), "a b c d")
	}

	if string(b) != "a b c d" {
		t.Errorf(`b = %#v, want %#v`, string(b), "a b c d")
	}
}

func TestSquashForUTF8(t *testing.T) {
	a := []byte("Hello　世界!")
	b := squash(a)

	if reflect.DeepEqual(a, b) {
		t.Errorf(`a = %#v, does NOT want %#v`, a, b)
	}

	if !reflect.DeepEqual(b, []byte("Hello 世界!")) {
		t.Errorf(`b = %#v, want %#v`, b, []byte("Hello 世界!"))
	}

	after := []byte("Hello　世界!")
	copy(after, []byte("Hello 世界!"))

	if !reflect.DeepEqual(a, after) {
		t.Errorf(`a = %#v, want %#v`, a, after)
	}
}
