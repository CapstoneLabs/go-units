package units

var (
	Mass = UnitOptionQuantity("mass")

	// Metric

	Gram      = NewUnit("gram", "g", Mass)
	ExaGram   = Exa(Gram, FactorLinear)
	PetaGram  = Peta(Gram, FactorLinear)
	TeraGram  = Tera(Gram, FactorLinear)
	GigaGram  = Giga(Gram, FactorLinear)
	MegaGram  = Mega(Gram, FactorLinear)
	KiloGram  = Kilo(Gram, FactorLinear)
	HectoGram = Hecto(Gram, FactorLinear)
	DecaGram  = Deca(Gram, FactorLinear)
	DeciGram  = Deci(Gram, FactorLinear)
	CentiGram = Centi(Gram, FactorLinear)
	MilliGram = Milli(Gram, FactorLinear)
	MicroGram = Micro(Gram, FactorLinear)
	NanoGram  = Nano(Gram, FactorLinear)
	PicoGram  = Pico(Gram, FactorLinear)
	FemtoGram = Femto(Gram, FactorLinear)
	AttoGram  = Atto(Gram, FactorLinear)

	MetricTon = NewUnit("metric ton", "t", Mass, SI)

	// Imperial

	Grain  = NewUnit("grain", "gr", Mass, BI)
	Drachm = NewUnit("drachm", "dr", Mass, BI)
	Ounce  = NewUnit("ounce", "oz", Mass, BI)
	Pound  = NewUnit("pound", "lb", Mass, BI)
	Stone  = NewUnit("stone", "st", Mass, BI)
	Ton    = NewUnit("ton", "t", Mass, BI)
	Slug   = NewUnit("slug", "", Mass, BI)
)

func init() {
	NewRatioConversion(MetricTon, KiloGram, 1000)

	NewRatioConversion(Grain, Gram, 0.06479891)
	NewRatioConversion(Drachm, Gram, 1.7718451953125)
	NewRatioConversion(Ounce, Gram, 28.349523125)
	NewRatioConversion(Pound, Gram, 453.59237)
	NewRatioConversion(Stone, Gram, 6350.29318)
	NewRatioConversion(Ton, Gram, 1016046.9088)
	NewRatioConversion(Slug, Gram, 14593.90294)
}
