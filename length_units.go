package units

var (
	Length = UnitOptionQuantity("length")

	// Metric

	Meter      = NewUnit("meter", "m", Length, SI, UnitOptionAliases("metre"))
	ExaMeter   = Exa(Meter, FactorLinear)
	PetaMeter  = Peta(Meter, FactorLinear)
	TeraMeter  = Tera(Meter, FactorLinear)
	GigaMeter  = Giga(Meter, FactorLinear)
	MegaMeter  = Mega(Meter, FactorLinear)
	KiloMeter  = Kilo(Meter, FactorLinear)
	HectoMeter = Hecto(Meter, FactorLinear)
	DecaMeter  = Deca(Meter, FactorLinear)
	DeciMeter  = Deci(Meter, FactorLinear)
	CentiMeter = Centi(Meter, FactorLinear)
	MilliMeter = Milli(Meter, FactorLinear)
	MicroMeter = Micro(Meter, FactorLinear)
	NanoMeter  = Nano(Meter, FactorLinear)
	PicoMeter  = Pico(Meter, FactorLinear)
	FemtoMeter = Femto(Meter, FactorLinear)
	AttoMeter  = Atto(Meter, FactorLinear)

	Angstrom = NewUnit("angstrom", "Å", Length, BI, UnitOptionPlural("angstroms"))
	Inch     = NewUnit("inch", "in", Length, BI, UnitOptionPlural("inches"))
	Foot     = NewUnit("foot", "ft", Length, BI, UnitOptionPlural("feet"))
	Yard     = NewUnit("yard", "yd", Length, BI)
	Mile     = NewUnit("mile", "mi", Length, BI)
	League   = NewUnit("league", "lea", Length, BI)
	Furlong  = NewUnit("furlong", "fur", Length, BI)
)

func init() {
	NewRatioConversion(Angstrom, Meter, 0.0000000001)
	NewRatioConversion(Inch, Meter, 0.0254)
	NewRatioConversion(Foot, Meter, 0.30479999999999996)
	NewRatioConversion(Yard, Meter, 0.9144)
	NewRatioConversion(Mile, Meter, 1609.344)
	NewRatioConversion(League, Meter, 4828.032)
	NewRatioConversion(Furlong, Meter, 201.168)
}
