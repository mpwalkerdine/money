// Package money is a convenience wrapper for "github.com/ericlagergren/decimal".
//
// It makes decimal values returned from this package immutable, at the expense of
// reduced memory efficiency. See https://golang.org/pkg/math/big/ for why the API
// is designed in the way it is. We forego that benefit in calling code, but where
// possible this package will attempt to minimise allocations during calculations.
// In the majority of cases however, monetary values fit inside the compact uint64
// value used by the underlying decimal package.
package money

import (
	"fmt"
	"strings"

	eld "github.com/ericlagergren/decimal"
)

type dec = *eld.Big

// RoundingMode determines how rounding is performed.
type RoundingMode = eld.RoundingMode

// Rounding constants.
const (
	ToNearestEven RoundingMode = eld.ToNearestEven
	ToNearestAway RoundingMode = eld.ToNearestAway
	ToZero        RoundingMode = eld.ToZero
	AwayFromZero  RoundingMode = eld.AwayFromZero
	ToNegativeInf RoundingMode = eld.ToNegativeInf
	ToPositiveInf RoundingMode = eld.ToPositiveInf
)

// Decimal is an immutable, arbitrary precision decimal number.
type Decimal struct {
	value dec
}

// New creates a monetary decimal for a given unit number.
func New(units int64) Decimal { return dec2(units * 100) }

// NewCents creates a monetary decimal for a given cent value.
func NewCents(cents int64) Decimal { return dec2(cents) }

// NewInt creates an integer decimal with zero scale.
func NewInt(i int) Decimal { return decr(int64(i), 0) }

// NewScalar creates an arbitrary scalar value.
func NewScalar(value int64, scale int) Decimal { return decr(value, scale) }

// Pc creates a decimal percent i.e. 2% = Pc(2) = 0.02.
func Pc(v int64) Decimal { return decr(v, 2) }

// Pm creates a decimal permille i.e. 3‰ = Pm(3) = 0.003.
func Pm(v int64) Decimal { return decr(v, 3) }

// Bp creates a new "Basis Point" / permyriad decimal i.e. 4‱ = Bp(4) = 0.0004.
func Bp(v int64) Decimal { return decr(v, 4) }

// Format implements the fmt.Formatter interface.
// Verbs are the same as for the underlying decimal.Big, except %v and %d are the same as %f.
// %c will multiply by 100, use %f and append '%'.
// If a precision is requested for negative scale decimals, these are appended.
func (d Decimal) Format(s fmt.State, c rune) {
	if d.value == nil || d.value.Sign() == 0 {
		// Handle special-case zero, we don't want Go operating mode treatment.
		d.value = new(eld.Big)
	} else {
		// Copy to manipulate
		d.value = zero().Copy(d.value)
	}

	if c == 'c' {
		defer fmt.Fprint(s, "%")
		d.value.Mul(d.value, eld.New(100, 0))
	}

	if strings.ContainsRune("vdc", c) {
		c = 'f'
	}

	if prec, hasPrec := s.Precision(); hasPrec && d.value.Scale() < prec {
		d.value.Quantize(prec)
	}

	d.value.Format(s, c)
}

// Equals returns true if the two numbers represent the same value.
func (d Decimal) Equals(other Decimal) bool {
	return !d.value.IsNaN(0) && !other.value.IsNaN(0) && d.value.Cmp(other.value) == 0
}

// EqualTo returns true if two numbers are equal to the specified significant figures.
func (d Decimal) EqualTo(other Decimal, sigfigs int) bool {
	a := zero().Copy(d.value).Round(sigfigs)
	b := zero().Copy(other.value).Round(sigfigs)
	return a.Sub(a, b).Sign() == 0
}

// LessThan than returns true if the receiver is less than the argument.
func (d Decimal) LessThan(other Decimal) bool {
	v, o := d.value, other.value
	if v == nil {
		v = zero()
	}
	if o == nil {
		o = zero()
	}
	return !v.IsNaN(0) && !o.IsNaN(0) && v.Cmp(o) < 0
}

func zero() dec {
	z := new(eld.Big)
	z.Context = eld.Context128
	z.Context.OperatingMode = eld.Go
	return z
}

func wrap(value dec) Decimal              { return Decimal{value} }
func dec2(value int64) Decimal            { return wrap(zero().SetMantScale(value, 2)) }
func decr(value int64, scale int) Decimal { return wrap(zero().SetMantScale(value, scale).Reduce()) }
