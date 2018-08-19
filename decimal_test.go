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

func ExampleDecimal_Format() {
	n := 0
	print := func(format string, d Decimal) {
		n++
		fmt.Printf("%2d) "+format+"\n", n, d)
	}

	print("%s", NewCents(123))
	print("%s", NewScalar(2345, -6))
	print("%v", NewScalar(2345, -6))
	print("%s", NewScalar(3, 9))
	print("%.1f", NewCents(4567))
	print("%.2f", New(5678))
	print("%.4f", NewCents(6789))
	print("`%5.2f`", New(7).Div(New(3)))
	print("'%-10.f'", NewCents(-80808))
	print("%.2f", NewScalar(9, 0))
	print("%.2f", NewScalar(10, 2))
	print("%.2f", NewScalar(11, -2))
	print("%.2f", NewScalar(12, 4))

	// Output:
	//  1) 1.23
	//  2) 2.345e+9
	//  3) 2345000000
	//  4) 3e-9
	//  5) 45.7
	//  6) 5678.00
	//  7) 67.8900
	//  8) ` 2.33`
	//  9) '-808      '
	// 10) 9.00
	// 11) 0.10
	// 12) 1100.00
	// 13) 0.00
}
