package money

import (
	"fmt"
)

func ExampleDecimal_RoundDP() {
	fmt.Print(Bp(1234).RoundDP(1, ToPositiveInf))
	// Output: 0.2
}

func ExampleDecimal_Round() {
	fmt.Print(NewCents(123999999).Round(3, ToNegativeInf))
	// Output: 1230000
}

func ExampleDecimal_PowInt() {
	fmt.Print(New(3).PowInt(3))
	// Output: 27.000000
}

func ExampleDecimal_Pow() {
	fmt.Print(New(16).Pow(Pc(50)).RoundDP(2, ToNearestAway))
	// Output: 4.00
}

func ExampleDecimal_PowFrac() {
	fmt.Print(New(27).PowFrac(1, 3).RoundDP(2, ToNearestEven))
	// Output: 3.00
}

func ExampleMax() {
	fmt.Print(Max(New(1), NewCents(200), NewInt(1)))
	// Output: 2.00
}
