package money

// NominalToEffectiveRate calculates (1+rate/periods)^periods - 1.
//
// A nominal rate in this context is just the periodic rate multiplied
// by the number of periods, which will be less than the effective rate
// due to the nature of compounding.
//
// For example, 6% compounded monthly means 0.5% is applied each month,
// which gives an effective rate of 6.17%.
func NominalToEffectiveRate(rate Decimal, periods int) Decimal {
	return rate.Div(NewInt(periods)).AddInt(1).PowInt(periods).SubInt(1)
}

// NominalToRealRate calculates (1+rate)/(1+inflation) - 1 (Fisher equation).
//
// A nominal rate in this context is the "before inflation rate".
func NominalToRealRate(rate, inflation Decimal) Decimal {
	return rate.AddInt(1).Div(inflation.AddInt(1)).SubInt(1)
}

// RealToNominalRate calculates (rate + 1)(1+inflation) - 1 (Fisher equation).
//
// A nominal rate in this context is the "before inflation rate".
func RealToNominalRate(rate, inflation Decimal) Decimal {
	return rate.AddInt(1).Mul(inflation.AddInt(1)).SubInt(1)
}

// EffectiveToNominalRate calculates ((1+rate)^(1/periods) - 1) * periods.
//
// This is the inverse of NominalToEffectiveRate.
func EffectiveToNominalRate(rate Decimal, periods int) Decimal {
	return EffectiveToPeriodicRate(rate, periods).Mul(NewInt(periods))
}

// EffectiveToPeriodicRate calculates (1+rate)^(1/periods) - 1.
//
// Given an effective rate, e.g. the amount of interest paid in a year,
// this calculates the rate used for periodic payments to achieve the
// same rate overall.
func EffectiveToPeriodicRate(rate Decimal, periods int) Decimal {
	return rate.AddInt(1).PowFrac(1, periods).SubInt(1)
}

// FutureValue calculates amount * (1+rate/periods)^(duration*periods).
//
// In other words, given an amount, interest rate and compounding frequency,
// this calculates what the future value will be at the end of the term.
func FutureValue(amount, rate Decimal, duration, periods int) Decimal {
	return amount.Mul(rate.Div(NewInt(periods)).AddInt(1).PowInt(duration * periods))
}

// RecompoundRate converts a rate from one compounding basis to another.
//
// For example, 4% compounded quarterly is a 3.98% compounded daily.
func RecompoundRate(rate Decimal, current, new int) Decimal {
	return rate.Div(NewInt(current)).AddInt(1).PowFrac(current, new).SubInt(1).Mul(NewInt(new))
}

// FutureValueOrdinaryAnnuity calculates amountPerPeriod * [ ((1+rate)^periods - 1) / rate ].
//
// This is the accumulated value when payments are made at the end of each period.
// Note that rate should be per period.
func FutureValueOrdinaryAnnuity(amountPerPeriod, rate Decimal, periods int) Decimal {
	return amountPerPeriod.Mul(rate.AddInt(1).PowInt(periods).SubInt(1).Div(rate))
}

// FutureValueAnnuityDue calculates amountPerPeriod * [ ((1+rate)^periods - 1) / rate ] * (1+rate).
//
// This is the accumulated value when payments are made at the beginning of each period.
// Note that rate should be per period.
func FutureValueAnnuityDue(amountPerPeriod, rate Decimal, periods int) Decimal {
	return FutureValueOrdinaryAnnuity(amountPerPeriod, rate, periods).Mul(rate.AddInt(1))
}

// Deflate calculates amount / (1 + inflation)^periods.
//
// This expresses a future value (after the given number of periods) in today's money.
// The inflation rate must be per unit of period
// i.e. if periods is 5 years, inflation must be per annum e.g. 0.05 for 5% pa.
func Deflate(amount, inflation Decimal, periods int) Decimal {
	return amount.Div(inflation.AddInt(1).PowInt(periods))
}
