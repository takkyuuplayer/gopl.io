package weightconv

import (
	"fmt"
	"testing"
)

func TestAssertString(t *testing.T) {
	k := Kilogram(1)

	if fmt.Sprintf("%s", k) == "1.0 kg" {
		t.Errorf(`k = %#v, do NOT want %#v`, fmt.Sprintf("%s", k), "1.0 kg")
	}

	p := Pound(1)

	if fmt.Sprintf("%s", p) == "1.0 lb" {
		t.Errorf(`p = %#v, do NOT want %#v`, fmt.Sprintf("%s", p), "1.0 lb")
	}
}

func TestConversion(t *testing.T) {
	k := Kilogram(1)

	if KtoP(k) != 1/KilogramPerPound {
		t.Errorf(`KtoP(k) = %#v, want %#v`, KtoP(k), 1/KilogramPerPound)
	}

	p := Pound(1)

	if PtoK(p) != KilogramPerPound {
		t.Errorf(`PtoK(p) = %#v, want %#v`, PtoK(p), KilogramPerPound)
	}
}
