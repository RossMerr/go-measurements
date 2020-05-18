package measurements

import "fmt"

type PressureUnit int32

const (
	Torr PressureUnit = iota
	Bar
	Pascal
	PoundForcePerSquareInch
)

var PressureUnitName = map[PressureUnit]string{
	Torr:                    "Torr",
	Bar:                     "bar",
	Pascal:                  "Pa",
	PoundForcePerSquareInch: "psi",
}

var PressureUnitValue = map[string]PressureUnit{
	"Torr": Torr,
	"bar":  Bar,
	"Pa":   Pascal,
	"psi":  PoundForcePerSquareInch,
}

func (s PressureUnit) String() string {
	return PressureUnitName[s]
}

type Pressure interface {
	Unit() PressureUnit
	Value() float64
	String() string

	To(PressureUnit) Pressure
	ToTorr() Pressure
	ToBar() Pressure
	ToPascal() Pressure
	ToPoundForcePerSquareInch() Pressure
}

type pressure struct {
	unit  PressureUnit
	value float64
}

func NewPressure(unit PressureUnit, value float64) Pressure {
	return &pressure{
		unit:  unit,
		value: value,
	}
}

func (s *pressure) Unit() PressureUnit {
	return s.unit
}

func (s *pressure) Value() float64 {
	return s.value
}

func (s pressure) String() string {
	return fmt.Sprintf("%.2f %s", s.value, s.unit)
}

func FromTorr(value float64) Pressure {
	return &pressure{unit: Torr, value: value}
}

func FromBar(value float64) Pressure {
	return &pressure{unit: Bar, value: value}
}

func FromPascal(value float64) Pressure {
	return &pressure{unit: Pascal, value: value}
}

func FromPoundForcePerSquareInch(value float64) Pressure {
	return &pressure{unit: PoundForcePerSquareInch, value: value}
}

func (s *pressure) To(unit PressureUnit) Pressure {
	switch unit {
	default: // Torr
		return s.ToTorr()
	case Bar:
		return s.ToBar()
	case Pascal:
		return s.ToPascal()
	case PoundForcePerSquareInch:
		return s.ToPoundForcePerSquareInch()
	}
}

func (s *pressure) ToTorr() Pressure {
	switch s.Unit() {
	default: // Torr
		return FromTorr(s.Value())
	case Bar:
		return FromTorr(s.Value() * 750)
	case Pascal:
		return FromTorr(s.Value() / 133)
	case PoundForcePerSquareInch:
		return FromTorr(s.Value() * 51.715)
	}
}

func (s *pressure) ToBar() Pressure {
	switch s.Unit() {
	default: // Torr
		return FromBar(s.Value() / 750)
	case Bar:
		return FromBar(s.Value())
	case Pascal:
		return FromBar(s.Value() / 100000)
	case PoundForcePerSquareInch:
		return FromBar(s.Value() / 14.504)
	}
}

func (s *pressure) ToPascal() Pressure {
	switch s.Unit() {
	default: // Torr
		return FromPascal(s.Value() * 133)
	case Bar:
		return FromPascal(s.Value() * 100000)
	case Pascal:
		return FromPascal(s.Value())
	case PoundForcePerSquareInch:
		return FromPascal(s.Value() * 6895)
	}
}

func (s *pressure) ToPoundForcePerSquareInch() Pressure {
	switch s.Unit() {
	default: // Torr
		return FromPoundForcePerSquareInch(s.Value() / 51.715)
	case Bar:
		return FromPoundForcePerSquareInch(s.Value() * 14.504)
	case Pascal:
		return FromPoundForcePerSquareInch(s.Value() / 6895)
	case PoundForcePerSquareInch:
		return FromPoundForcePerSquareInch(s.Value())
	}
}
