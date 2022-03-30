package units

var (
	Pressure = UnitOptionQuantity("pressure")

	// SI unit metric
	Pascal      = NewUnit("pascal", "Pa", Pressure, SI)
	ExaPascal   = Exa(Pascal, FactorLinear)
	PetaPascal  = Peta(Pascal, FactorLinear)
	TeraPascal  = Tera(Pascal, FactorLinear)
	GigaPascal  = Giga(Pascal, FactorLinear)
	MegaPascal  = Mega(Pascal, FactorLinear)
	KiloPascal  = Kilo(Pascal, FactorLinear)
	HectoPascal = Hecto(Pascal, FactorLinear)
	DecaPascal  = Deca(Pascal, FactorLinear)
	DeciPascal  = Deci(Pascal, FactorLinear)
	CentiPascal = Centi(Pascal, FactorLinear)
	MilliPascal = Milli(Pascal, FactorLinear)
	MicroPascal = Micro(Pascal, FactorLinear)
	NanoPascal  = Nano(Pascal, FactorLinear)
	PicoPascal  = Pico(Pascal, FactorLinear)
	FemtoPascal = Femto(Pascal, FactorLinear)
	AttoPascal  = Atto(Pascal, FactorLinear)

	// Other
	At       = NewUnit("technical atmosphere", "at", Pressure, BI, UnitOptionPlural("technical atmospheres"))
	Atm      = NewUnit("standard atmosphere", "atm", Pressure, BI, UnitOptionPlural("standard atmospheres"))
	Bar      = NewUnit("bar", "bar", Pressure, BI, UnitOptionPlural("bars"))
	CentiBar = Centi(Bar, FactorLinear)
	MilliBar = Milli(Bar, FactorLinear)
	MicroBar = Micro(Bar, FactorLinear)
	Barye    = NewUnit("barye", "Ba", Pressure, BI, UnitOptionPlural("baryes"))
	InH2O    = NewUnit("inch of Water Column", "inH2O", Pressure, BI)
	InHg     = NewUnit("inch of Mercury", "inHg", Pressure, BI)
	MH2O     = NewUnit("meter of Water Column", "mmH2O", Pressure, BI, UnitOptionPlural("meters of Water Column"))
	MmH2O    = Milli(MH2O, FactorLinear)
	CmH2O    = Centi(MH2O, FactorLinear)
	MHg      = NewUnit("meter of Mercury", "mmHg", Pressure, BI, UnitOptionPlural("meters of Mercury"))
	MmHg     = Milli(MHg, FactorLinear)
	CmHg     = Centi(MHg, FactorLinear)
	Newton   = NewUnit("newton per square meter", "N/m²", Pressure, BI)
	Psi      = NewUnit("pound-force per square inch", "psi", Pressure, BI)
	Torr     = NewUnit("torr", "Torr", Pressure, BI)
)

func init() {
	NewRatioConversion(At, Pascal, 98066.5)
	NewRatioConversion(Atm, Pascal, 101325.2738)
	NewRatioConversion(Bar, Pascal, 98000)
	NewRatioConversion(Barye, Pascal, 0.1)
	NewRatioConversion(InH2O, Pascal, 248.84)
	NewRatioConversion(InHg, Pascal, 3386.38815789)
	NewRatioConversion(MH2O, Pascal, 980.665)
	NewRatioConversion(MHg, Pascal, 13332.2368421)
	NewRatioConversion(Newton, Pascal, 1)
	NewRatioConversion(Psi, Pascal, 6894.757)
	NewRatioConversion(Torr, Pascal, 133.322368421)
}
