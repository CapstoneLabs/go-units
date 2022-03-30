package units

var (
	Area = UnitOptionQuantity("area")

	// Metric

	SquareMeter      = NewUnit("square meter", "m2", Area, SI)
	SquareExaMeter   = Exa(SquareMeter, FactorSquare)
	SquarePetaMeter  = Peta(SquareMeter, FactorSquare)
	SquareTeraMeter  = Tera(SquareMeter, FactorSquare)
	SquareGigaMeter  = Giga(SquareMeter, FactorSquare)
	SquareMegaMeter  = Mega(SquareMeter, FactorSquare)
	SquareKiloMeter  = Kilo(SquareMeter, FactorSquare)
	SquareHectoMeter = Hecto(SquareMeter, FactorSquare)
	SquareDecaMeter  = Deca(SquareMeter, FactorSquare)
	SquareDeciMeter  = Deci(SquareMeter, FactorSquare)
	SquareCentiMeter = Centi(SquareMeter, FactorSquare)
	SquareMilliMeter = Milli(SquareMeter, FactorSquare)
	SquareMicroMeter = Micro(SquareMeter, FactorSquare)
	SquareNanoMeter  = Nano(SquareMeter, FactorSquare)
	SquarePicoMeter  = Pico(SquareMeter, FactorSquare)
	SquareFemtoMeter = Femto(SquareMeter, FactorSquare)
	SquareAttoMeter  = Atto(SquareMeter, FactorSquare)

	SquareInch = NewUnit("square inch", "in2", Area, BI)
	SquareFoot = NewUnit("square foot", "ft2", Area, BI)
	SquareYard = NewUnit("square yard", "yd2", Area, BI)
)

func init() {
	NewRatioConversion(SquareInch, SquareMeter, 0.00064516)
	NewRatioConversion(SquareFoot, SquareMeter, 0.09290303999999999)
	NewRatioConversion(SquareYard, SquareMeter, 0.83612736)
}
