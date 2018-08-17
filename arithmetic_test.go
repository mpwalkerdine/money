package money

import (
	"fmt"
	"testing"
)

func ExampleDecimal_Add() {
	a := New(5)
	b := New(10)
	fmt.Print(a.Add(b))
	// Output: 15.00
}

func ExampleDecimal_AddInt() {
	fmt.Println(NewCents(99).AddInt(1))
	// Output: 1.99
}

func ExampleDecimal_Sub() {
	a := New(5)
	b := New(10)
	fmt.Print(a.Sub(b))
	// Output: -5.00
}

func ExampleDecimal_SubInt() {
	fmt.Print(NewCents(499).SubInt(4))
	// Output: 0.99
}

func ExampleDecimal_Mul() {
	a := New(50)
	b := Pc(50)
	fmt.Print(a.Mul(b))
	// Output: 25.000
}

func ExampleDecimal_Div() {
	a := New(1)
	b := New(3)
	fmt.Print(a.Div(b))
	// Output: 0.3333333333333333333333333333333334
}

func TestMul(t *testing.T) {
	tests := []struct {
		a        Decimal
		b        Decimal
		expected Decimal
	}{
		{New(0), New(0), New(0)},
		{New(1), New(0), New(0)},
		{New(1), New(1), New(1)},
		{New(-1), New(1), New(-1)},
		{New(-1), New(-1), New(1)},
		{New(100), Pm(1), decr(1, 1)},
		{Bp(1234), New(1), Bp(1234)},
		{Bp(1234), Bp(1), decr(1234, 8)},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v * %v", tc.a, tc.b), func(t *testing.T) {
			actual := tc.a.Mul(tc.b)
			if !actual.Equals(tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
