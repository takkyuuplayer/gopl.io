package ch04

import (
	"reflect"
	"testing"
	"unicode/utf8"
)

func reverse07(b []byte) []byte {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		switch size {
		case 1:
		case 2, 3:
			b[i], b[i+size-1] = b[i+size-1], b[i]
		case 4:
			b[i], b[i+size-1] = b[i+size-1], b[i]
			b[i+1], b[i+size-2] = b[i+size-2], b[i+1]
		default: // out of utf8 range
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
	a := []byte("string ＵＴＦ−８🍺.")
	b := reverse07(a)

	if !reflect.DeepEqual(a, []byte(".🍺８−ＦＴＵ gnirts")) {
		t.Errorf(`a = %#v, want %#v`, a, []byte(".🍺８−ＦＴＵ gnirts"))
	}

	if !reflect.DeepEqual(b, a) {
		t.Errorf(`b = %#v, want %#v`, b, a)
	}
}
