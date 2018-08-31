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

// GoalSeek attempts to find a value x (to the specified precision), where min <= x <= max,
// such that f(x) = target.
func GoalSeek(min, max, target Decimal, precision int, f func(Decimal) Decimal) (Decimal, bool) {
	// Translate target to zero i.e. g(x) = f(x) - target = 0
	g := func(x Decimal) Decimal {
		return f(x).Sub(target)
	}

	// min is a solution
	if g(min).value.Sign() == 0 {
		return min, true
	}

	// max is a solution
	if g(max).value.Sign() == 0 {
		return max, true
	}

	// Find zero-crossing bracket
	left, right, ok := bracket(min, max, precision, g)

	// No solutions
	if !ok {
		return wrap(zero()), false
	}

	// left bracket is a solution
	if g(left).value.Sign() == 0 {
		return left, true
	}

	// right bracket is a solution
	if g(right).value.Sign() == 0 {
		return right, true
	}

	// bisect the bracket to find the solution
	result, found := bisect(left, right, precision, g)

	return result, found
}

func bracket(min, max Decimal, precision int, g func(Decimal) Decimal) (Decimal, Decimal, bool) {
	// Sweep left to right in smaller and smaller increments until a bracket is found
	// or the values from g are identical to a given precision, then give up.
	inc := max.Sub(min)

	for {
		for left := min; left.LessThan(max); left = left.Add(inc) {
			right := left.Add(inc)
			gleft, gright := g(left), g(right)
			leftSign := gleft.value.Signbit()
			rightSign := gright.value.Signbit()

			// zero-crossing
			if leftSign != rightSign {
				return left, right, true
			}

			// still no crossings but values are now the same, give up
			if left.EqualTo(right, precision) && gleft.EqualTo(gright, precision) {
				return min, max, false
			}
		}
		inc = inc.Mul(half)
	}
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

		if left.EqualTo(right, prec) {
			left.value.Round(prec)
			return left, true
		}
	}

	mid.value.Round(prec)
	return mid, false
}
