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
