package money

// NominalToEffectiveRate calculates (1+rate/periods)^periods - 1.
func NominalToEffectiveRate(rate Decimal, periods int) Decimal {
	return rate.Div(NewInt(periods)).AddInt(1).PowInt(periods).SubInt(1)
}
