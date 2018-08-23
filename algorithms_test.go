package money

import (
	"fmt"
	"testing"
)

func ExampleGoalSeek() {
	initial := New(1)

	// Find x s.t. x^2+2 = 11 (i.e. √9 = 3)
	fmt.Println(GoalSeek(initial, New(11), 2, func(x Decimal) Decimal { return x.PowInt(2).AddInt(2) }))

	// Find x s.t. 0.01x^2+0.8 = 1 (i.e. 10√0.2 ≅ 4.472)
	fmt.Println(GoalSeek(initial, New(1), 4, func(x Decimal) Decimal { return x.PowInt(2).Mul(Pc(1)).Add(Pc(80)) }))

	// Find x s.t. x^2 = -1 (i.e. √-1, no real solution)
	fmt.Println(GoalSeek(initial, New(-1), 3, func(x Decimal) Decimal { return x.PowInt(2) }))

	// Output:
	// 3.0 true
	// 4.472 true
	// 1.00 false
}

func TestGoalSeek(t *testing.T) {
	for i, tc := range []struct {
		name    string
		initial Decimal
		target  Decimal
		prec    int
		f       func(Decimal) Decimal
		want    Decimal
		ok      bool
	}{
		{"x=1", New(1), New(1), 3, func(x Decimal) Decimal { return x }, New(1), true},
		{"x=5", New(10), New(5), 3, func(x Decimal) Decimal { return x }, New(5), true},
		{"x^2-2x=1", New(0), New(1), 3, func(x Decimal) Decimal { return x.PowInt(2).Sub(New(2).Mul(x)) }, Pm(-414), true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := GoalSeek(tc.initial, tc.target, tc.prec, tc.f)
			if !got.Equals(tc.want) || ok != tc.ok {
				t.Errorf("\n#%d wanted (%v,%t), got (%v,%t)", i, tc.want, tc.ok, got, ok)
			}
		})
	}
}
