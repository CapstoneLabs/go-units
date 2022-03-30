package units

var (
	Time = UnitOptionQuantity("time")

	Second      = NewUnit("second", "s", Time)
	ExaSecond   = Exa(Second, FactorLinear)
	PetaSecond  = Peta(Second, FactorLinear)
	TeraSecond  = Tera(Second, FactorLinear)
	GigaSecond  = Giga(Second, FactorLinear)
	MegaSecond  = Mega(Second, FactorLinear)
	KiloSecond  = Kilo(Second, FactorLinear)
	HectoSecond = Hecto(Second, FactorLinear)
	DecaSecond  = Deca(Second, FactorLinear)
	DeciSecond  = Deci(Second, FactorLinear)
	CentiSecond = Centi(Second, FactorLinear)
	MilliSecond = Milli(Second, FactorLinear)
	MicroSecond = Micro(Second, FactorLinear)
	NanoSecond  = Nano(Second, FactorLinear)
	PicoSecond  = Pico(Second, FactorLinear)
	FemtoSecond = Femto(Second, FactorLinear)
	AttoSecond  = Atto(Second, FactorLinear)

	Minute = NewUnit("minute", "min", Time)
	Hour   = NewUnit("hour", "hr", Time)
	Day    = NewUnit("day", "d", Time)
	Month  = NewUnit("month", "", Time)
	Year   = NewUnit("year", "yr", Time)

	Decade     = NewUnit("decade", "", Time)
	Century    = NewUnit("century", "", Time)
	Millennium = NewUnit("millennium", "", Time)

	// more esoteric time units
	PlanckTime = NewUnit("planck time", "𝑡ₚ", Time)
	Fortnight  = NewUnit("fortnight", "", Time)
	Score      = NewUnit("score", "", Time)
)

func init() {
	NewRatioConversion(Minute, Second, 60.0)
	NewRatioConversion(Hour, Second, 3600.0)
	NewRatioConversion(Day, Hour, 24.0)
	NewRatioConversion(Month, Day, 30.0)
	NewRatioConversion(Year, Day, 365.25)

	NewRatioConversion(Decade, Year, 10.0)
	NewRatioConversion(Century, Year, 100.0)
	NewRatioConversion(Millennium, Year, 1000.0)

	NewRatioConversion(PlanckTime, Second, 5.39e-44)
	NewRatioConversion(Fortnight, Day, 14)
	NewRatioConversion(Score, Year, 20.0)
}
