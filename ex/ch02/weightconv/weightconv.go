package weightconv

import (
	"fmt"
)

type Kilogram float64
type Pound float64

const KilogramPerPound = 0.45359237

func (p Pound) String() string {
	return fmt.Sprintf("%g lb", p)
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%g kg", k)
}

func PtoK(p Pound) Kilogram {
	return Kilogram(p * KilogramPerPound)
}

func KtoP(k Kilogram) Pound {
	return Pound(k * 1.0 / KilogramPerPound)
}
