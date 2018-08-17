package money

import (
	"fmt"
)

func ExampleNew() {
	fmt.Print(New(50))
	// Output: 50.00
}

func ExampleNewCents() {
	fmt.Print(NewCents(1234))
	// Output: 12.34
}

func ExampleNewScalar() {
	fmt.Print(NewScalar(42, 0))
	//Output: 42
}

func ExamplePc() {
	fmt.Print(Pc(50))
	// Output: 0.5
}

func ExamplePm() {
	fmt.Print(Pm(20))
	// Output: 0.02
}

func ExampleBp() {
	fmt.Print(Bp(10))
	// Output: 0.001
}

func ExampleDecimal_Equals() {
	a := New(1)
	b := NewCents(100)
	fmt.Print(a.Equals(b))
	// Output: true
}
