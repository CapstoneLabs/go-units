package units

var (
	Volume = UnitOptionQuantity("volume")

	// Metric

	Liter      = NewUnit("liter", "l", Volume, SI, UnitOptionAliases("litre"))
	ExaLiter   = Exa(Liter, FactorLinear)
	PetaLiter  = Peta(Liter, FactorLinear)
	TeraLiter  = Tera(Liter, FactorLinear)
	GigaLiter  = Giga(Liter, FactorLinear)
	MegaLiter  = Mega(Liter, FactorLinear)
	KiloLiter  = Kilo(Liter, FactorLinear)
	HectoLiter = Hecto(Liter, FactorLinear)
	DecaLiter  = Deca(Liter, FactorLinear)
	DeciLiter  = Deci(Liter, FactorLinear)
	CentiLiter = Centi(Liter, FactorLinear)
	MilliLiter = Milli(Liter, FactorLinear)
	MicroLiter = Micro(Liter, FactorLinear)
	NanoLiter  = Nano(Liter, FactorLinear)
	PicoLiter  = Pico(Liter, FactorLinear)
	FemtoLiter = Femto(Liter, FactorLinear)
	AttoLiter  = Atto(Liter, FactorLinear)

	CubicMeter      = NewUnit("cubic meter", "m3", Volume, SI)
	CubicExaMeter   = Exa(CubicMeter, FactorCubic)
	CubicPetaMeter  = Peta(CubicMeter, FactorCubic)
	CubicTeraMeter  = Tera(CubicMeter, FactorCubic)
	CubicGigaMeter  = Giga(CubicMeter, FactorCubic)
	CubicMegaMeter  = Mega(CubicMeter, FactorCubic)
	CubicKiloMeter  = Kilo(CubicMeter, FactorCubic)
	CubicHectoMeter = Hecto(CubicMeter, FactorCubic)
	CubicDecaMeter  = Deca(CubicMeter, FactorCubic)
	CubicDeciMeter  = Deci(CubicMeter, FactorCubic)
	CubicCentiMeter = Centi(CubicMeter, FactorCubic)
	CubicMilliMeter = Milli(CubicMeter, FactorCubic)
	CubicMicroMeter = Micro(CubicMeter, FactorCubic)
	CubicNanoMeter  = Nano(CubicMeter, FactorCubic)
	CubicPicoMeter  = Pico(CubicMeter, FactorCubic)
	CubicFemtoMeter = Femto(CubicMeter, FactorCubic)
	CubicAttoMeter  = Atto(CubicMeter, FactorCubic)

	// Imperial

	Quart      = NewUnit("quart", "qt", Volume, BI)
	Pint       = NewUnit("pint", "pt", Volume, BI)
	Gallon     = NewUnit("gallon", "gal", Volume, BI)
	FluidOunce = NewUnit("fluid ounce", "fl oz", Volume, BI, UnitOptionAliases("floz"))

	CubicInch = NewUnit("cubic inch", "in3", Volume, BI)
	CubicFoot = NewUnit("cubic foot", "ft3", Volume, BI)
	CubicYard = NewUnit("cubic yard", "yd3", Volume, BI)

	// US

	FluidQuart          = NewUnit("fluid quart", "", Volume, US)
	FluidPint           = NewUnit("fluid pint", "", Volume, US)
	FluidGallon         = NewUnit("fluid gallon", "", Volume, US)
	CustomaryFluidOunce = NewUnit("customary fluid ounce", "", Volume, US)
	Barrel              = NewUnit("barrel", "bbl", Volume, US)
)

func init() {
	NewRatioConversion(CubicMeter, Liter, 1000.0)

	// Imperial

	NewRatioConversion(Quart, Liter, 1.1365225)
	NewRatioConversion(Pint, Liter, 0.56826125)
	NewRatioConversion(Gallon, Liter, 4.54609)
	NewRatioConversion(FluidOunce, MilliLiter, 28.4130625)

	NewRatioConversion(CubicInch, Liter, 0.016387063999999996)
	NewRatioConversion(CubicFoot, Liter, 28.316846592)
	NewRatioConversion(CubicYard, Liter, 764.554857984)

	// US

	NewRatioConversion(FluidQuart, Liter, 0.946352946)
	NewRatioConversion(FluidPint, Liter, 0.473176473)
	// NewRatioConversion(FluidGallon, Liter, 3.7854117839999998)
	NewRatioConversion(FluidGallon, Liter, 3.7854117839999994)
	NewRatioConversion(CustomaryFluidOunce, MilliLiter, 29.5735295625)
	NewRatioConversion(Barrel, Liter, 119.24047119599999)
}
