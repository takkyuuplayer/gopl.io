package comma

import (
	"strings"
	"testing"
)

func comma(s string) string {
	if strings.HasPrefix(s, "-") {
		return s[:1] + comma(s[1:])
	}
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		return comma(s[:dot]) + s[dot:]
	}

	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}

func TestComma(t *testing.T) {
	if comma("123") != "123" {
		t.Errorf(`comma("123") = %#v, want %#v`, comma("123"), "123")
	}

	if comma("1234") != "1,234" {
		t.Errorf(`comma("1234") = %#v, want %#v`, comma("1234"), "1,234")
	}

	if comma("-1234") != "-1,234" {
		t.Errorf(`comma("-1234") = %#v, want %#v`, comma("-1234"), "-1,234")
	}

	if comma("-1234.1234") != "-1,234.1234" {
		t.Errorf(`comma("-1234.1234") = %#v, want %#v`, comma("-1234.1234"), "-1,234.1234")
	}
}
