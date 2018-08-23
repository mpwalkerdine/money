package money

import (
	"math"
)

var (
	maxVal = wrap(zero().SetUint64(math.MaxUint64))
	minVal = wrap(zero().Neg(maxVal.value))
	ten    = wrap(zero().SetMantScale(10, 0))
	half   = wrap(zero().SetMantScale(5, 1))
)

// GoalSeek attempts to find a value x (to the specified precision), such that f(x) = target,
// using initial as a starting point.
func GoalSeek(initial, target Decimal, precision int, f func(Decimal) Decimal) (Decimal, bool) {
	// Translate target to zero g(x) = f(x) - target = 0
	g := func(x Decimal) Decimal {
		return f(x).Sub(target)
	}

	// Maybe we got lucky
	if g(initial).value.Sign() == 0 {
		return initial, true
	}

	left, right, ok := bracket(initial, g)

	// No solutions
	if !ok {
		return initial, false
	}

	result, found := bisect(left, right, precision, g)

	return result, found
}

func bracket(initial Decimal, g func(Decimal) Decimal) (Decimal, Decimal, bool) {
	guessSign := g(initial).value.Signbit()

	// Find a bracket by expanding outwards from initial in 0.1, 1, 10, etc. increments
	var left, right = initial, initial
	for i := NewScalar(1, 1); minVal.LessThan(left) || right.LessThan(maxVal); i = i.Mul(ten) {
		right = initial.Add(i)
		if g(right).value.Signbit() != guessSign {
			left = initial
			return left, right, true
		}

		left = initial.Sub(i)
		if g(left).value.Signbit() != guessSign {
			right = initial
			return left, right, true
		}
	}

	return initial, initial, false
}

func bisect(left, right Decimal, prec int, g func(Decimal) Decimal) (Decimal, bool) {
	var mid, test Decimal
	leftSign := g(left).value.Signbit()

	for i := 0; i < 1000; i++ {
		mid = left.Add(right).Mul(half)
		test = g(mid)
		if test.value.Sign() == 0 {
			mid.value.Round(prec)
			return mid, true
		}

		if test.value.Signbit() == leftSign {
			left = mid
		} else {
			right = mid
		}

		if test.value.Sub(right.value, left.value).SetScale(test.value.Scale()-prec).Quantize(prec).Sign() == 0 {
			left.value.Round(prec)
			return left, true
		}
	}

	mid.value.Round(prec)
	return mid, false
}
