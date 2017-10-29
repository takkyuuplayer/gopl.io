package comma

import "testing"

func TestComma(t *testing.T) {
	if comma("12345") != "12,345" {
		t.Errorf(`comma("12345") = %#v, want %#v`, comma("12345"), "12,345")
	}
}
