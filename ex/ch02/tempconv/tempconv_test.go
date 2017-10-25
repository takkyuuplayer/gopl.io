package tempconv

import (
	"fmt"
	"testing"
)

func TestAssertString(t *testing.T) {
	if fmt.Sprintf("%s", AbsoluteZeroC) != "-273.15°C" {
		t.Errorf(`AbsoluteZeroC = %#v, want %#v`, fmt.Sprintf("%s", AbsoluteZeroC), "-273.15°C")
	}

	if fmt.Sprintf("%s", CtoF(BoilingC)) != "133°F" {
		t.Errorf(`CtoF(BoilingC) = %#v, want %#v`, fmt.Sprintf("%s", CtoF(BoilingC)), "-273.15°F")
	}

	if fmt.Sprintf("%s", CtoK(BoilingC)) != "373.15°K" {
		t.Errorf(`CtoK(BoilingC) = %#v, want %#v`, fmt.Sprintf("%s", CtoK(BoilingC)), "373.15°F")
	}
}
