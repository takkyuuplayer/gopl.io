package ch04

import (
	"reflect"
	"testing"
	"unicode/utf8"
)

func reverse07(b []byte) []byte {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		if size > 1 { // 3 byte at most
			b[i], b[i+size-1] = b[i+size-1], b[i]
		}
		i += size
	}

	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}

	return b
}

func TestAscii(t *testing.T) {
	a := []byte("ascii")
	b := reverse07(a)

	if !reflect.DeepEqual(a, []byte("iicsa")) {
		t.Errorf(`a = %#v, want %#v`, a, []byte("iicsa"))
	}

	if !reflect.DeepEqual(b, a) {
		t.Errorf(`b = %#v, want %#v`, b, a)
	}
}

func TestUTF8(t *testing.T) {
	a := []byte("string ＵＴＦ−８.")
	b := reverse07(a)

	if !reflect.DeepEqual(a, []byte(".８−ＦＴＵ gnirts")) {
		t.Errorf(`a = %#v, want %#v`, a, []byte(".８−ＦＴＵ gnirts"))
	}

	if !reflect.DeepEqual(b, a) {
		t.Errorf(`b = %#v, want %#v`, b, a)
	}
}
