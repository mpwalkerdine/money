package money

// Add calculates a + b.
func (a Decimal) Add(b Decimal) Decimal {
	return wrap(zero().Add(a.value, b.value))
}

// AddInt calculates a + b.
func (a Decimal) AddInt(b int) Decimal {
	return a.Add(NewInt(b))
}

// Sub calculates a - b.
func (a Decimal) Sub(b Decimal) Decimal {
	return wrap(zero().Sub(a.value, b.value))
}

// SubInt calculates a - b.
func (a Decimal) SubInt(b int) Decimal {
	return a.Sub(NewInt(b))
}

// Mul calculates a * b.
func (a Decimal) Mul(b Decimal) Decimal {
	return wrap(zero().Mul(a.value, b.value))
}

// Div calculates a / b.
func (a Decimal) Div(b Decimal) Decimal {
	return wrap(zero().Quo(a.value, b.value))
}
