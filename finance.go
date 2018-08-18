package money

// NominalToEffectiveRate calculates (1+rate/periods)^periods - 1.
func NominalToEffectiveRate(rate Decimal, periods int) Decimal {
	return rate.Div(NewInt(periods)).AddInt(1).PowInt(periods).SubInt(1)
}

// NominalToRealRate calculates (1+rate)/(1+inflation) - 1 (Fisher equation)
func NominalToRealRate(rate, inflation Decimal) Decimal {
	return rate.AddInt(1).Div(inflation.AddInt(1)).SubInt(1)
}

// EffectiveToNominalRate calculates ((1+rate)^(1/periods) - 1) * periods = eff
func EffectiveToNominalRate(rate Decimal, periods int) Decimal {
	return rate.AddInt(1).PowFrac(1, periods).SubInt(1).Mul(NewInt(periods))
}
