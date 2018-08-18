package money

import (
	"fmt"
)

func ExampleNominalToEffectiveRate() {
	nominal := Pc(6)
	effective := NominalToEffectiveRate(nominal, 12)
	fmt.Print(effective.Round(3, ToNearestAway))
	// Output: 0.0617
}

func ExampleNominalToRealRate() {
	nominal := Pc(8)
	inflation := Pc(8)
	real := NominalToRealRate(nominal, inflation)
	fmt.Print(real)
	// Output: 0
}

func ExampleEffectiveToNominalRate() {
	effective := Bp(617)
	nominal := EffectiveToNominalRate(effective, 12)
	fmt.Print(nominal.Round(3, ToNearestAway))
	// Output: 0.0600
}
