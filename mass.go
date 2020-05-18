package measurements

import "fmt"

type MassUnit int32

const (
	Kilogram MassUnit = iota
	Gram
	Pound
	Ounce
)

var MassUnitName = map[MassUnit]string{
	Kilogram: "kg",
	Gram:     "g",
	Pound:    "lb",
	Ounce:    "oz",
}

var MassUnitValue = map[string]MassUnit{
	"kg": Kilogram,
	"g":  Gram,
	"lb": Pound,
	"oz": Ounce,
}

func (s MassUnit) String() string {
	return MassUnitName[s]
}

type Mass interface {
	Unit() MassUnit
	Value() float64
	String() string

	To(unit MassUnit) Mass
	ToKilogram() Mass
	ToGram() Mass
	ToPound() Mass
	ToOunce() Mass
}

type mass struct {
	unit  MassUnit
	value float64
}

func NewMass(unit MassUnit, value float64) Mass {
	return &mass{
		unit:  unit,
		value: value,
	}
}

func (s *mass) Unit() MassUnit {
	return s.unit
}

func (s *mass) Value() float64 {
	return s.value
}

func (s mass) String() string {
	return fmt.Sprintf("%.2f %s", s.value, s.unit)
}

func FromOunce(value float64) Mass {
	return &mass{Ounce, value}
}

func FromPound(value float64) Mass {
	return &mass{Pound, value}
}

func FromGram(value float64) Mass {
	return &mass{Gram, value}
}

func FromKilogram(value float64) Mass {
	return &mass{Kilogram, value}
}

func (s *mass) To(unit MassUnit) Mass {
	switch unit {
	default: // Kilogram
		return s.ToKilogram()
	case Gram:
		return s.ToGram()
	case Pound:
		return s.ToPound()
	case Ounce:
		return s.ToOunce()
	}
}

func (s *mass) ToKilogram() Mass {
	switch s.Unit() {
	default: // Kilogram
		return FromKilogram(s.Value())
	case Gram:
		return FromKilogram(s.Value() / 1000)
	case Pound:
		return FromKilogram(s.Value() / 2.205)
	case Ounce:
		return FromKilogram(s.Value() / 35.274)
	}
}

func (s *mass) ToGram() Mass {
	switch s.Unit() {
	default: // Kilogram
		return FromGram(s.Value() * 1000)
	case Gram:
		return FromGram(s.Value())
	case Pound:
		return FromGram(s.Value() * 454)
	case Ounce:
		return FromGram(s.Value() * 28.35)
	}
}

func (s *mass) ToPound() Mass {
	switch s.Unit() {
	default: // Kilogram
		return FromPound(s.Value() * 2.205)
	case Gram:
		return FromPound(s.Value() / 454)
	case Pound:
		return FromPound(s.Value())
	case Ounce:
		return FromPound(s.Value() / 16)
	}
}

func (s *mass) ToOunce() Mass {
	switch s.Unit() {
	default: // Kilogram
		return FromOunce(s.Value() * 35.274)
	case Gram:
		return FromOunce(s.Value() / 28.35)
	case Pound:
		return FromOunce(s.Value() * 16)
	case Ounce:
		return FromOunce(s.Value())
	}
}