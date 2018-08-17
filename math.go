package money

import (
	"math/big"

	"github.com/ericlagergren/decimal/math"
)

// RoundDP rounds the decimal to the specified number of decimal places.
func (d Decimal) RoundDP(dp int, mode RoundingMode) Decimal {
	r := zero().Set(d.value)

	r.Context.RoundingMode = mode
	r.Quantize(dp)
	r.Context.RoundingMode = d.value.Context.RoundingMode

	return wrap(r)
}

// Round rounds the decimal to the specified number of significant figures.
func (d Decimal) Round(sigfigs int, mode RoundingMode) Decimal {
	r := zero().Set(d.value)

	r.Context.RoundingMode = mode
	r.Round(sigfigs)
	r.Context.RoundingMode = d.value.Context.RoundingMode

	return wrap(r)
}

// PowInt calculates d^i.
func (d Decimal) PowInt(i int) Decimal {
	n := zero().SetUint64(uint64(i))
	return wrap(math.Pow(zero(), d.value, n))
}

// Pow calculates d^n.
func (d Decimal) Pow(n Decimal) Decimal {
	return wrap(math.Pow(zero(), d.value, n.value))
}

// PowFrac calculates d^(num/denom).
func (d Decimal) PowFrac(num, denom int) Decimal {
	r := big.NewRat(int64(num), int64(denom))
	n := zero().SetRat(r)
	return wrap(math.Pow(zero(), d.value, n))
}
