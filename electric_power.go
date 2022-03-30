package units

var (
	ElectricPower = UnitOptionQuantity("e-power")

	Watt      = NewUnit("watt", "W", ElectricPower)
	ExaWatt   = Exa(Watt, FactorLinear)
	PetaWatt  = Peta(Watt, FactorLinear)
	TeraWatt  = Tera(Watt, FactorLinear)
	GigaWatt  = Giga(Watt, FactorLinear)
	MegaWatt  = Mega(Watt, FactorLinear)
	KiloWatt  = Kilo(Watt, FactorLinear)
	HectoWatt = Hecto(Watt, FactorLinear)
	DecaWatt  = Deca(Watt, FactorLinear)
	DeciWatt  = Deci(Watt, FactorLinear)
	CentiWatt = Centi(Watt, FactorLinear)
	MilliWatt = Milli(Watt, FactorLinear)
	MicroWatt = Micro(Watt, FactorLinear)
	NanoWatt  = Nano(Watt, FactorLinear)
	PicoWatt  = Pico(Watt, FactorLinear)
	FemtoWatt = Femto(Watt, FactorLinear)
	AttoWatt  = Atto(Watt, FactorLinear)
)
