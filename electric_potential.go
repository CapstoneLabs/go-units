package units

var (
	ElectricPotential = UnitOptionQuantity("e-potential")

	Volt      = NewUnit("volt", "V", ElectricPotential)
	ExaVolt   = Exa(Volt, FactorLinear)
	PetaVolt  = Peta(Volt, FactorLinear)
	TeraVolt  = Tera(Volt, FactorLinear)
	GigaVolt  = Giga(Volt, FactorLinear)
	MegaVolt  = Mega(Volt, FactorLinear)
	KiloVolt  = Kilo(Volt, FactorLinear)
	HectoVolt = Hecto(Volt, FactorLinear)
	DecaVolt  = Deca(Volt, FactorLinear)
	DeciVolt  = Deci(Volt, FactorLinear)
	CentiVolt = Centi(Volt, FactorLinear)
	MilliVolt = Milli(Volt, FactorLinear)
	MicroVolt = Micro(Volt, FactorLinear)
	NanoVolt  = Nano(Volt, FactorLinear)
	PicoVolt  = Pico(Volt, FactorLinear)
	FemtoVolt = Femto(Volt, FactorLinear)
	AttoVolt  = Atto(Volt, FactorLinear)
)
