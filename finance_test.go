package money

import (
	"fmt"
	"testing"
)

func ExampleNominalToEffectiveRate() {
	nominal := Pc(6)
	effective := NominalToEffectiveRate(nominal, 12)
	fmt.Print(effective.Round(3, ToNearestAway))
	// Output: 0.0617
}

func ExampleNominalToRealRate() {
	nominal := Bp(521)  // 5.21%
	inflation := Pm(25) // 2.50%
	real := NominalToRealRate(nominal, inflation)
	fmt.Print(real.Round(3, ToNearestAway))
	// Output: 0.0264
}

func ExampleRealToNominalRate() {
	real := Bp(264)     // 2.64%
	inflation := Pm(25) // 2.50%
	nominal := RealToNominalRate(real, inflation)
	fmt.Print(nominal.Round(3, ToNearestAway))
	// Output: 0.0521
}

func ExampleEffectiveToNominalRate() {
	effective := Bp(617) // 6.17%
	nominal := EffectiveToNominalRate(effective, 12)
	fmt.Print(nominal.Round(3, ToNearestAway))
	// Output: 0.0600
}

func ExampleEffectiveToPeriodicRate() {
	annual := Pc(10)
	daily := EffectiveToPeriodicRate(annual, 365)
	fmt.Print(daily.Round(6, ToNearestAway))
	// Output: 0.000261158
}

func ExampleFutureValue() {
	amount := New(1500) // £1,500
	rate := Pm(43)      // 4.3%
	periods := 4        // Compounded quarterly
	duration := 6       // 6 years

	final := FutureValue(amount, rate, duration, periods)
	fmt.Print(final.RoundDP(2, ToNearestEven))
	// Output: 1938.84
}

func ExampleRecompoundRate() {
	annualQuarterly := Pc(4)
	annualDaily := RecompoundRate(annualQuarterly, 4, 365)
	fmt.Print(annualDaily.Round(3, ToNearestAway))
	// Output: 0.0398
}

func ExampleFutureValueOrdinaryAnnuity() {
	amountPerPeriod := New(1000) // £1000 at the end of each year
	rate := Pc(5)                // 5% PA
	periods := 5                 // 5 years
	result := FutureValueOrdinaryAnnuity(amountPerPeriod, rate, periods)
	fmt.Printf("%.2f", result)
	// Output: 5525.63
}

func ExampleFutureValueAnnuityDue() {
	amountPerPeriod := New(1000) // £1000 at the start of each year
	rate := Pc(5)                // 5% PA
	periods := 5                 // 5 years
	result := FutureValueAnnuityDue(amountPerPeriod, rate, periods)
	fmt.Printf("%.2f", result)
	// Output: 5801.91
}

func ExampleDeflate() {
	amount := New(1000) // £1000
	years := 20         // in 20 years
	inflation := Pm(25) // with 2.5% inflation p.a.
	fmt.Printf("%.2f", Deflate(amount, inflation, years))
	// Output: 610.27
}

func TestEffectiveToPeriodicRate(t *testing.T) {
	principle := New(10000)
	annual := Pc(10)
	daily := EffectiveToPeriodicRate(annual, 365)
	expected := principle.Mul(annual.AddInt(1))

	actual := principle
	for i := 0; i < 365; i++ {
		actual = actual.Mul(daily.AddInt(1)).RoundDP(3, ToNearestEven)
	}

	if !actual.Equals(expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
