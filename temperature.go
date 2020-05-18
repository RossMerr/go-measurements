package measurements

import "fmt"

const DegreeSign = "°"

type TemperatureUnit int32

const (
	Celsius TemperatureUnit = iota
	Fahrenheit
	Kelvin
)

var TemperatureUnitTypeName = map[TemperatureUnit]string{
	Celsius:    "C",
	Fahrenheit: "F",
	Kelvin:     "K",
}

var TemperatureUnitTypeValue = map[string]TemperatureUnit{
	"C": Celsius,
	"F": Fahrenheit,
	"K": Kelvin,
}

func (s TemperatureUnit) String() string {
	return TemperatureUnitTypeName[s]
}

type Temperature interface {
	Unit() TemperatureUnit
	Value() float64
	To(unit TemperatureUnit) Temperature
	ToCelsius() Temperature
	ToFahrenheit() Temperature
	ToKelvin() Temperature
	String() string
}

func NewTemperature(unit TemperatureUnit, value float64) Temperature {
	return &temperature{
		unit:  unit,
		value: value,
	}
}

type temperature struct {
	unit  TemperatureUnit
	value float64
}

func (s *temperature) Unit() TemperatureUnit {
	return s.unit
}

func (s *temperature) Value() float64 {
	return s.value
}

func FromCelsius(value float64) Temperature {
	return &temperature{value: value, unit: Celsius}
}

func FromFahrenheit(value float64) Temperature {
	return &temperature{value: value, unit: Fahrenheit}
}

func FromKelvin(value float64) Temperature {
	return &temperature{value: value, unit: Kelvin}
}

func (s *temperature) To(unit TemperatureUnit) Temperature {
	switch unit {
	case Fahrenheit:
		return s.ToFahrenheit()
	case Kelvin:
		return s.ToKelvin()
	default:
		return s.ToCelsius()
	}
}

func (s *temperature) ToCelsius() Temperature {
	switch s.Unit() {
	case Fahrenheit:
		return FromCelsius((s.Value() - 32) * 5 / 9)
	case Kelvin:
		return FromCelsius(s.Value() - 273.15)
	default:
		return FromCelsius(s.Value())
	}
}

func (s *temperature) ToFahrenheit() Temperature {
	switch s.Unit() {
	case Fahrenheit:
		return FromFahrenheit(s.Value())
	case Kelvin:
		return FromFahrenheit((s.Value()-273.15)*9/5 + 32)
	default:
		return FromFahrenheit((s.Value() * 9 / 5) + 32)
	}
}

func (s *temperature) ToKelvin() Temperature {
	switch s.Unit() {
	case Fahrenheit:
		return FromKelvin((s.Value()-32)*5/9 + 273.15)
	case Kelvin:
		return FromKelvin(s.Value())
	default:
		return FromKelvin(s.Value() + 273.15)
	}
}

func (s temperature) String() string {
	if s.unit == Kelvin {
		return fmt.Sprintf("%.2f %s", s.value, s.unit)
	}

	return fmt.Sprintf("%.2f %s%s", s.value, DegreeSign, s.unit)
}
