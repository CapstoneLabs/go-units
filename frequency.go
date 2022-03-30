package units

var (
	Frequency = UnitOptionQuantity("frequency")

	Hertz      = NewUnit("hertz", "Hz", Frequency)
	ExaHertz   = Exa(Hertz, FactorLinear)
	PetaHertz  = Peta(Hertz, FactorLinear)
	TeraHertz  = Tera(Hertz, FactorLinear)
	GigaHertz  = Giga(Hertz, FactorLinear)
	MegaHertz  = Mega(Hertz, FactorLinear)
	KiloHertz  = Kilo(Hertz, FactorLinear)
	HectoHertz = Hecto(Hertz, FactorLinear)
	DecaHertz  = Deca(Hertz, FactorLinear)
	DeciHertz  = Deci(Hertz, FactorLinear)
	CentiHertz = Centi(Hertz, FactorLinear)
	MilliHertz = Milli(Hertz, FactorLinear)
	MicroHertz = Micro(Hertz, FactorLinear)
	NanoHertz  = Nano(Hertz, FactorLinear)
	PicoHertz  = Pico(Hertz, FactorLinear)
	FemtoHertz = Femto(Hertz, FactorLinear)
	AttoHertz  = Atto(Hertz, FactorLinear)
)
