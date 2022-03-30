package units

import (
	"math"
)

type magnitude struct {
	Symbol string
	Prefix string
	Power  float64
}

var (
	keyExa   = "exa"
	keyPeta  = "peta"
	keyTera  = "tera"
	keyGiga  = "giga"
	keyMega  = "mega"
	keyKilo  = "kilo"
	keyHecto = "hecto"
	keyDeca  = "deca"
	keyDeci  = "deci"
	keyCenti = "centi"
	keyMilli = "milli"
	keyMicro = "micro"
	keyNano  = "nano"
	keyPico  = "pico"
	keyFemto = "femto"
	keyAtto  = "atto"

	mags = map[string]magnitude{
		keyExa:   {"E", "exa", 18.0},
		keyPeta:  {"P", "peta", 15.0},
		keyTera:  {"T", "tera", 12.0},
		keyGiga:  {"G", "giga", 9.0},
		keyMega:  {"M", "mega", 6.0},
		keyKilo:  {"k", "kilo", 3.0},
		keyHecto: {"h", "hecto", 2.0},
		keyDeca:  {"da", "deca", 1.0},
		keyDeci:  {"d", "deci", -1.0},
		keyCenti: {"c", "centi", -2.0},
		keyMilli: {"m", "milli", -3.0},
		keyMicro: {"u", "micro", -6.0}, // μ
		keyNano:  {"n", "nano", -9.0},
		keyPico:  {"p", "pico", -12.0},
		keyFemto: {"f", "femto", -15.0},
		keyAtto:  {"a", "atto", -18.0},
	}
)

// Magnitude prefix methods create and return a new Unit, while automatically registering
// conversions to and from the provided base Unit

type MagnitudeFactor = float64

const (
	FactorLinear MagnitudeFactor = iota + 1.0
	FactorSquare
	FactorCubic
)

func magnify(mag string, base Unit, factor MagnitudeFactor, opts ...UnitOption) Unit {
	return mags[mag].makeUnit(base, factor, opts...)
}

func Exa(base Unit, factor MagnitudeFactor, opts ...UnitOption) Unit {
	return magnify(keyExa, base, factor, opts...)
}

func Peta(base Unit, factor MagnitudeFactor, opts ...UnitOption) Unit {
	return magnify(keyPeta, base, factor, opts...)
}

func Tera(base Unit, factor MagnitudeFactor, opts ...UnitOption) Unit {
	return magnify(keyTera, base, factor, opts...)
}

func Giga(base Unit, factor MagnitudeFactor, opts ...UnitOption) Unit {
	return magnify(keyGiga, base, factor, opts...)
}

func Mega(base Unit, factor MagnitudeFactor, opts ...UnitOption) Unit {
	return magnify(keyMega, base, factor, opts...)
}

func Kilo(base Unit, factor MagnitudeFactor, opts ...UnitOption) Unit {
	return magnify(keyKilo, base, factor, opts...)
}

func Hecto(base Unit, factor MagnitudeFactor, opts ...UnitOption) Unit {
	return magnify(keyHecto, base, factor, opts...)
}

func Deca(base Unit, factor MagnitudeFactor, opts ...UnitOption) Unit {
	return magnify(keyDeca, base, factor, opts...)
}

func Deci(base Unit, factor MagnitudeFactor, opts ...UnitOption) Unit {
	return magnify(keyDeci, base, factor, opts...)
}

func Centi(base Unit, factor MagnitudeFactor, opts ...UnitOption) Unit {
	return magnify(keyCenti, base, factor, opts...)
}

func Milli(base Unit, factor MagnitudeFactor, opts ...UnitOption) Unit {
	return magnify(keyMilli, base, factor, opts...)
}

func Micro(base Unit, factor MagnitudeFactor, opts ...UnitOption) Unit {
	return magnify(keyMicro, base, factor, opts...)
}

func Nano(base Unit, factor MagnitudeFactor, opts ...UnitOption) Unit {
	return magnify(keyNano, base, factor, opts...)
}

func Pico(base Unit, factor MagnitudeFactor, opts ...UnitOption) Unit {
	return magnify(keyPico, base, factor, opts...)
}

func Femto(base Unit, factor MagnitudeFactor, opts ...UnitOption) Unit {
	return magnify(keyFemto, base, factor, opts...)
}

func Atto(base Unit, factor MagnitudeFactor, opts ...UnitOption) Unit {
	return magnify(keyAtto, base, factor, opts...)
}

// Create magnitude unit and conversion given a base unit
func (mag magnitude) makeUnit(base Unit, factor MagnitudeFactor, addOpts ...UnitOption) Unit {
	name := mag.Prefix + base.Name
	symbol := mag.Symbol + base.Symbol

	// set system to metric by default
	opts := []UnitOption{SI}

	// create prefixed aliases if needed
	for _, alias := range base.aliases {
		magAlias := mag.Prefix + alias
		opts = append(opts, UnitOptionAliases(magAlias))
	}

	// append any supplemental options
	for _, opt := range addOpts {
		opts = append(opts, opt)
	}

	// append quantity name opt
	opts = append(opts, UnitOptionQuantity(base.Quantity))

	u := NewUnit(name, symbol, opts...)

	// only create conversions to and from base unit
	ratio := 1.0 * math.Pow(10.0, mag.Power*factor)
	NewRatioConversion(u, base, ratio)

	return u
}
