package money

import (
	"fmt"
	"testing"
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

func ExampleDecimal_EqualTo() {
	a := NewCents(101)
	b := NewCents(100)
	fmt.Println(a.EqualTo(b, 1))
	fmt.Println(a.EqualTo(b, 2))
	fmt.Println(a.EqualTo(b, 3))
	// Output:
	// true
	// true
	// false
}

func ExampleDecimal_LessThan() {
	fmt.Print(New(1).LessThan(NewCents(100)))
	// Output: false
}

func ExampleDecimal_Format() {
	print := func(format string, d Decimal) {
		fmt.Printf(format+"\n", d)
	}

	print("%s", NewCents(123))
	print("%s", NewScalar(2345, -6))
	print("%v", NewScalar(2345, -6))
	print("%.2f", New(5678))
	print("`%5.2f`", New(7).Div(New(3)))
	print("'%-10.f'", NewCents(-80808))
	print("%.2c", Pm(25))

	// Output:
	// 1.23
	// 2.345e+9
	// 2345000000
	// 5678.00
	// ` 2.33`
	// '-808      '
	// 2.50%
}

func TestDecimalFormat(t *testing.T) {
	for i, tc := range []struct {
		fs   string
		d    Decimal
		want string
	}{
		{"%s", NewCents(123), "1.23"},
		{"%s", NewScalar(2345, -6), "2.345e+9"},
		{"%v", NewScalar(2345, -6), "2345000000"},
		{"%s", NewScalar(3, 9), "3e-9"},
		{"%.1f", NewCents(4567), "45.7"},
		{"%.2f", New(5678), "5678.00"},
		{"%.4f", NewCents(6789), "67.8900"},
		{"`%5.2f`", New(7).Div(New(3)), "` 2.33`"},
		{"'%-10.f'", NewCents(-80808), "'-808      '"},
		{"%.2f", NewScalar(9, 0), "9.00"},
		{"%.2f", NewScalar(10, 2), "0.10"},
		{"%.2f", NewScalar(11, -2), "1100.00"},
		{"%.2f", NewScalar(12, 4), "0.00"},
		{"%.2f", Decimal{}, "0.00"},
		{"%.2f", decr(0, 0), "0.00"},
		{"%.2f", decr(1, 0), "1.00"},
		{"%5.2f", Decimal{}, " 0.00"},
		{"%5.2f", decr(0, 0), " 0.00"},
		{"%5.2f", decr(1, 0), " 1.00"},
		{"%.2c", Decimal{}, "0.00%"},
		{"%10.2c", decr(1, 0), "    100.00%"},
		{"%.2c", decr(101, -1), "101000.00%"},
	} {
		got := fmt.Sprintf(tc.fs, tc.d)
		if got != tc.want {
			t.Errorf("\n#%d\n got: %v\nwant: %v\n", i, got, tc.want)
		}
	}
}

func TestDecimal_EqualTo(t *testing.T) {
	for i, tc := range []struct {
		a, b Decimal
		sf   int
		want bool
	}{
		{New(0), New(0), 1, true},
		{New(1), New(0), 1, false},
		{NewCents(100), New(1), 3, true},
		{NewScalar(123, 2), NewScalar(123, 3), 3, false},
		{NewScalar(1230, 1), NewScalar(123, 0), 3, true},
		{NewScalar(123, -1), NewScalar(1230, 0), 3, true},
	} {
		tc := tc
		t.Run(fmt.Sprintf("#%d %v equals %v [%d sf]", i, tc.a, tc.b, tc.sf), func(t *testing.T) {
			t.Parallel()
			got := tc.a.EqualTo(tc.b, tc.sf)
			if got != tc.want {
				t.Errorf("wanted %t, got %t", tc.want, got)
			}
		})
	}
}
