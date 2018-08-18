package money

import (
	"fmt"
)

func ExampleNominalToEffectiveRate() {
	annual := Pc(6)
	monthly := NominalToEffectiveRate(annual, 12)
	fmt.Print(monthly.Round(3, ToNearestAway))
	// Output: 0.0617
}

func ExampleNominalToRealRate() {
	annual := Pc(8)
	inflation := Pc(8)
	real := NominalToRealRate(annual, inflation)
	fmt.Print(real)
	// Output: 0
}
