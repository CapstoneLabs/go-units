package units

var (
	Speed = UnitOptionQuantity("speed")

	MetersPerSecond   = NewUnit("meters per second", "m/s", Speed)
	KilometersPerHour = NewUnit("kilo meters per hour", "kph", Speed)
	FeetPerSecond     = NewUnit("feet per second", "ft/s", Speed)
	MilesPerHour      = NewUnit("miles per hour", "mph", Speed)
	Knot              = NewUnit("knot", "kt", Speed)
)

func init() {
	NewRatioConversion(KilometersPerHour, MetersPerSecond, 0.2777777777777778)
	NewRatioConversion(FeetPerSecond, MetersPerSecond, 0.30479999999999996)
	NewRatioConversion(MilesPerHour, MetersPerSecond, 0.44704)
	NewRatioConversion(Knot, MetersPerSecond, 0.5144444444444445)
}
