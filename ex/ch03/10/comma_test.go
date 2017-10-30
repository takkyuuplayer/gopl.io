package comma

import (
	"bytes"
	"testing"
)

func comma(s string) string {
	n := len(s)

	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	for i, v := range s {
		if i > 0 && (n-i)%3 == 0 {
			buf.WriteRune(',')
		}
		buf.WriteRune(v)
	}

	return buf.String()
}

func TestComma(t *testing.T) {
	if comma("123") != "123" {
		t.Errorf(`comma("123") = %#v, want %#v`, comma("123"), "123")
	}

	if comma("1234") != "1,234" {
		t.Errorf(`comma("1234") = %#v, want %#v`, comma("1234"), "1,234")
	}

	if comma("12345") != "12,345" {
		t.Errorf(`comma("12345") = %#v, want %#v`, comma("12345"), "12,345")
	}

	if comma("123456") != "123,456" {
		t.Errorf(`comma("123456") = %#v, want %#v`, comma("123456"), "123,456")
	}
}
