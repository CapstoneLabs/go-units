package units

var (
	ElectricCurrent = UnitOptionQuantity("e-current")

	Ampere      = NewUnit("ampere", "A", ElectricCurrent)
	ExaAmpere   = Exa(Ampere, FactorLinear)
	PetaAmpere  = Peta(Ampere, FactorLinear)
	TeraAmpere  = Tera(Ampere, FactorLinear)
	GigaAmpere  = Giga(Ampere, FactorLinear)
	MegaAmpere  = Mega(Ampere, FactorLinear)
	KiloAmpere  = Kilo(Ampere, FactorLinear)
	HectoAmpere = Hecto(Ampere, FactorLinear)
	DecaAmpere  = Deca(Ampere, FactorLinear)
	DeciAmpere  = Deci(Ampere, FactorLinear)
	CentiAmpere = Centi(Ampere, FactorLinear)
	MilliAmpere = Milli(Ampere, FactorLinear)
	MicroAmpere = Micro(Ampere, FactorLinear)
	NanoAmpere  = Nano(Ampere, FactorLinear)
	PicoAmpere  = Pico(Ampere, FactorLinear)
	FemtoAmpere = Femto(Ampere, FactorLinear)
	AttoAmpere  = Atto(Ampere, FactorLinear)
)
