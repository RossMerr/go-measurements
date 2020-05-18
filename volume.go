package measurements

import "fmt"

const CubicSymbol = "³"

type VolumeType int32

const (
	Milliliter VolumeType = iota
	Litre
	USfluidOunce
	USlegalCup
	USliquidPint
	USLiquidQuart
	USLiquidGallon
	ImperialFluidOunce
	ImperialCup
	ImperialPint
	ImperialQuart
	ImperialGallon
	//CubicMetres
)

var VolumeTypeName = map[VolumeType]string{
	Milliliter:         "ml",
	Litre:              "l",
	USfluidOunce:       "fl oz",
	USlegalCup:         "cp",
	USliquidPint:       "pt",
	USLiquidQuart:      "qt",
	USLiquidGallon:     "gal",
	ImperialFluidOunce: "imp fl oz",
	ImperialCup:        "imp cp",
	ImperialPint:       "imp pt",
	ImperialQuart:      "imp qt",
	ImperialGallon:     "imp gal",
	//CubicMetres:        "m" + CubicSymbol,
}

var VolumeTypeValue = map[string]VolumeType{
	"ml":        Milliliter,
	"l":         Litre,
	"fl oz":     USfluidOunce,
	"cp":        USlegalCup,
	"pt":        USliquidPint,
	"qt":        USLiquidQuart,
	"gal":       USLiquidGallon,
	"imp fl oz": ImperialFluidOunce,
	"imp cp":    ImperialCup,
	"imp pt":    ImperialPint,
	"imp qt":    ImperialQuart,
	"imp gal":   ImperialGallon,
	//"m" + CubicSymbol: CubicMetres,
}

func (s VolumeType) String() string {
	return VolumeTypeName[s]
}

type Volume interface {
	Unit() VolumeType
	Value() float64
	String() string

	To(VolumeType) Volume
	ToMilliliter() Volume
	ToLiter() Volume
	ToUSfluidOunce() Volume
	ToUSlegalCup() Volume
	ToUSliquidPint() Volume
	ToUSLiquidQuart() Volume
	ToUSLiquidGallon() Volume
	ToImperialFluidOunce() Volume
	ToImperialCup() Volume
	ToImperialPint() Volume
	ToImperialQuart() Volume
	ToImperialGallon() Volume
}

type volume struct {
	unit  VolumeType
	value float64
}

func NewVolume(unit VolumeType, value float64) Volume {
	return &volume{
		unit:  unit,
		value: value,
	}
}

func (s *volume) Unit() VolumeType {
	return s.unit
}

func (s *volume) Value() float64 {
	return s.value
}

func (s volume) String() string {
	return fmt.Sprintf("%.2f %s", s.value, s.unit)
}

func FromImperialGallon(value float64) Volume {
	return &volume{unit: ImperialGallon, value: value}
}

func FromImperialQuart(value float64) Volume {
	return &volume{unit: ImperialQuart, value: value}
}

func FromImperialPint(value float64) Volume {
	return &volume{unit: ImperialPint, value: value}
}

func FromImperialCup(value float64) Volume {
	return &volume{unit: ImperialCup, value: value}
}

func FromImperialFluidOunce(value float64) Volume {
	return &volume{unit: ImperialFluidOunce, value: value}
}

func FromUSLiquidGallon(value float64) Volume {
	return &volume{unit: USLiquidGallon, value: value}
}

func FromUSLiquidQuart(value float64) Volume {
	return &volume{unit: USLiquidQuart, value: value}
}

func FromUSliquidPint(value float64) Volume {
	return &volume{unit: USliquidPint, value: value}
}

func FromUSlegalCup(value float64) Volume {
	return &volume{unit: USlegalCup, value: value}
}

func FromUSfluidOunce(value float64) Volume {
	return &volume{unit: USfluidOunce, value: value}
}

func FromLiter(value float64) Volume {
	return &volume{unit: Litre, value: value}
}

func FromMilliliter(value float64) Volume {
	return &volume{unit: Milliliter, value: value}
}

func (s *volume) To(unit VolumeType) Volume {
	switch unit {
	default:
		return s.ToMilliliter()
	case Litre:
		return s.ToLiter()
	case USfluidOunce:
		return s.ToUSfluidOunce()
	case USlegalCup:
		return s.ToUSlegalCup()
	case USliquidPint:
		return s.ToUSliquidPint()
	case USLiquidQuart:
		return s.ToUSLiquidQuart()
	case USLiquidGallon:
		return s.ToUSLiquidGallon()
	case ImperialFluidOunce:
		return s.ToImperialFluidOunce()
	case ImperialCup:
		return s.ToImperialGallon()
	case ImperialPint:
		return s.ToImperialPint()
	case ImperialQuart:
		return s.ToImperialQuart()
	case ImperialGallon:
		return s.ToImperialGallon()
	}
}

func (s *volume) ToMilliliter() Volume {
	switch s.Unit() {
	default:
		return FromMilliliter(s.Value())
	case Litre:
		return FromMilliliter(s.Value() * 1000)
	case USfluidOunce:
		return FromMilliliter(s.Value() * 29.574)
	case USlegalCup:
		return FromMilliliter(s.Value() * 240)
	case USliquidPint:
		return FromMilliliter(s.Value() * 473)
	case USLiquidQuart:
		return FromMilliliter(s.Value() * 946)
	case USLiquidGallon:
		return FromMilliliter(s.Value() * 3785)
	case ImperialFluidOunce:
		return FromMilliliter(s.Value() * 28.413)
	case ImperialCup:
		return FromMilliliter(s.Value() * 284)
	case ImperialPint:
		return FromMilliliter(s.Value() * 568)
	case ImperialQuart:
		return FromMilliliter(s.Value() * 1137)
	case ImperialGallon:
		return FromMilliliter(s.Value() * 4546)
	}
}

func (s *volume) ToLiter() Volume {
	switch s.Unit() {
	default:
		return FromLiter(s.Value() / 1000)
	case Litre:
		return FromLiter(s.Value())
	case USfluidOunce:
		return FromLiter(s.Value() / 33.814)
	case USlegalCup:
		return FromLiter(s.Value() / 4.167)
	case USliquidPint:
		return FromLiter(s.Value() / 2.113)
	case USLiquidQuart:
		return FromLiter(s.Value() / 1.057)
	case USLiquidGallon:
		return FromLiter(s.Value() * 3.785)
	case ImperialFluidOunce:
		return FromLiter(s.Value() / 35.195)
	case ImperialCup:
		return FromLiter(s.Value() / 3.52)
	case ImperialPint:
		return FromLiter(s.Value() / 1.76)
	case ImperialQuart:
		return FromLiter(s.Value() * 1.137)
	case ImperialGallon:
		return FromLiter(s.Value() * 4.546)
	}
}

func (s *volume) ToUSfluidOunce() Volume {
	switch s.Unit() {
	default:
		return FromUSfluidOunce(s.Value() / 29.574)
	case Litre:
		return FromUSfluidOunce(s.Value() * 33.814)
	case USfluidOunce:
		return FromUSfluidOunce(s.Value())
	case USlegalCup:
		return FromUSfluidOunce(s.Value() * 8.115)
	case USliquidPint:
		return FromUSfluidOunce(s.Value() * 16)
	case USLiquidQuart:
		return FromUSfluidOunce(s.Value() * 32)
	case USLiquidGallon:
		return FromUSfluidOunce(s.Value() * 128)
	case ImperialFluidOunce:
		return FromUSfluidOunce(s.Value() / 1.041)
	case ImperialCup:
		return FromUSfluidOunce(s.Value() * 9.608)
	case ImperialPint:
		return FromUSfluidOunce(s.Value() * 19.215)
	case ImperialQuart:
		return FromUSfluidOunce(s.Value() * 38.43)
	case ImperialGallon:
		return FromUSfluidOunce(s.Value() * 154)
	}
}

func (s *volume) ToUSlegalCup() Volume {
	switch s.Unit() {
	default:
		return FromUSlegalCup(s.Value() / 240)
	case Litre:
		return FromUSlegalCup(s.Value() * 4.167)
	case USfluidOunce:
		return FromUSlegalCup(s.Value() / 8.115)
	case USlegalCup:
		return FromUSlegalCup(s.Value())
	case USliquidPint:
		return FromUSlegalCup(s.Value() * 1.972)
	case USLiquidQuart:
		return FromUSlegalCup(s.Value() * 3.943)
	case USLiquidGallon:
		return FromUSlegalCup(s.Value() * 15.773)
	case ImperialFluidOunce:
		return FromUSlegalCup(s.Value() / 8.447)
	case ImperialCup:
		return FromUSlegalCup(s.Value() * 1.184)
	case ImperialPint:
		return FromUSlegalCup(s.Value() * 2.368)
	case ImperialQuart:
		return FromUSlegalCup(s.Value() * 4.736)
	case ImperialGallon:
		return FromUSlegalCup(s.Value() * 18.942)
	}
}

func (s *volume) ToUSliquidPint() Volume {
	switch s.Unit() {
	default:
		return FromUSliquidPint(s.Value() / 473)
	case Litre:
		return FromUSliquidPint(s.Value() * 2.113)
	case USfluidOunce:
		return FromUSliquidPint(s.Value() / 16)
	case USlegalCup:
		return FromUSliquidPint(s.Value() / 1.972)
	case USliquidPint:
		return FromUSliquidPint(s.Value())
	case USLiquidQuart:
		return FromUSliquidPint(s.Value() * 2)
	case USLiquidGallon:
		return FromUSliquidPint(s.Value() * 8)
	case ImperialFluidOunce:
		return FromUSliquidPint(s.Value() / 16.653)
	case ImperialCup:
		return FromUSliquidPint(s.Value() / 1.665)
	case ImperialPint:
		return FromUSliquidPint(s.Value() * 1.201)
	case ImperialQuart:
		return FromUSliquidPint(s.Value() * 2.402)
	case ImperialGallon:
		return FromUSliquidPint(s.Value() * 9.608)
	}
}

func (s *volume) ToUSLiquidQuart() Volume {
	switch s.Unit() {
	default:
		return FromUSLiquidQuart(s.Value() / 946)
	case Litre:
		return FromUSLiquidQuart(s.Value() * 1.057)
	case USfluidOunce:
		return FromUSLiquidQuart(s.Value() / 32)
	case USlegalCup:
		return FromUSLiquidQuart(s.Value() / 3.943)
	case USliquidPint:
		return FromUSLiquidQuart(s.Value() / 2)
	case USLiquidQuart:
		return FromUSLiquidQuart(s.Value())
	case USLiquidGallon:
		return FromUSLiquidQuart(s.Value() * 4)
	case ImperialFluidOunce:
		return FromUSLiquidQuart(s.Value() / 33.307)
	case ImperialCup:
		return FromUSLiquidQuart(s.Value() / 3.331)
	case ImperialPint:
		return FromUSLiquidQuart(s.Value() / 1.665)
	case ImperialQuart:
		return FromUSLiquidQuart(s.Value() * 1.201)
	case ImperialGallon:
		return FromUSLiquidQuart(s.Value() * 4.804)
	}
}

func (s *volume) ToUSLiquidGallon() Volume {
	switch s.Unit() {
	default:
		return FromUSLiquidGallon(s.Value() / 3785)
	case Litre:
		return FromUSLiquidGallon(s.Value() / 3.785)
	case USfluidOunce:
		return FromUSLiquidGallon(s.Value() / 128)
	case USlegalCup:
		return FromUSLiquidGallon(s.Value() / 15.773)
	case USliquidPint:
		return FromUSLiquidGallon(s.Value() / 8)
	case USLiquidQuart:
		return FromUSLiquidGallon(s.Value() / 4)
	case USLiquidGallon:
		return FromUSLiquidGallon(s.Value())
	case ImperialFluidOunce:
		return FromUSLiquidGallon(s.Value() / 133)
	case ImperialCup:
		return FromUSLiquidGallon(s.Value() / 13.323)
	case ImperialPint:
		return FromUSLiquidGallon(s.Value() / 6.661)
	case ImperialQuart:
		return FromUSLiquidGallon(s.Value() / 3.331)
	case ImperialGallon:
		return FromUSLiquidGallon(s.Value() * 1.201)
	}
}

func (s *volume) ToImperialFluidOunce() Volume {
	switch s.Unit() {
	default:
		return FromImperialFluidOunce(s.Value() / 28.413)
	case Litre:
		return FromImperialFluidOunce(s.Value() * 35.195)
	case USfluidOunce:
		return FromImperialFluidOunce(s.Value() * 1.041)
	case USlegalCup:
		return FromImperialFluidOunce(s.Value() * 8.447)
	case USliquidPint:
		return FromImperialFluidOunce(s.Value() * 16.653)
	case USLiquidQuart:
		return FromImperialFluidOunce(s.Value() * 33.307)
	case USLiquidGallon:
		return FromImperialFluidOunce(s.Value() * 133)
	case ImperialFluidOunce:
		return FromImperialFluidOunce(s.Value())
	case ImperialCup:
		return FromImperialFluidOunce(s.Value() * 10)
	case ImperialPint:
		return FromImperialFluidOunce(s.Value() * 20)
	case ImperialQuart:
		return FromImperialFluidOunce(s.Value() * 40)
	case ImperialGallon:
		return FromImperialFluidOunce(s.Value() * 160)
	}
}

func (s *volume) ToImperialCup() Volume {
	switch s.Unit() {
	default:
		return FromImperialCup(s.Value() / 284)
	case Litre:
		return FromImperialCup(s.Value() * 3.52)
	case USfluidOunce:
		return FromImperialCup(s.Value() / 9.608)
	case USlegalCup:
		return FromImperialCup(s.Value() / 1.184)
	case USliquidPint:
		return FromImperialCup(s.Value() * 1.665)
	case USLiquidQuart:
		return FromImperialCup(s.Value() * 3.331)
	case USLiquidGallon:
		return FromImperialCup(s.Value() * 13.323)
	case ImperialFluidOunce:
		return FromImperialCup(s.Value() / 10)
	case ImperialCup:
		return FromImperialCup(s.Value())
	case ImperialPint:
		return FromImperialCup(s.Value() * 2)
	case ImperialQuart:
		return FromImperialCup(s.Value() * 4)
	case ImperialGallon:
		return FromImperialCup(s.Value() * 16)
	}
}

func (s *volume) ToImperialPint() Volume {
	switch s.Unit() {
	default:
		return FromImperialPint(s.Value() / 568)
	case Litre:
		return FromImperialPint(s.Value() * 1.76)
	case USfluidOunce:
		return FromImperialPint(s.Value() / 19.215)
	case USlegalCup:
		return FromImperialPint(s.Value() / 2.368)
	case USliquidPint:
		return FromImperialPint(s.Value() / 1.201)
	case USLiquidQuart:
		return FromImperialPint(s.Value() * 1.665)
	case USLiquidGallon:
		return FromImperialPint(s.Value() * 6.661)
	case ImperialFluidOunce:
		return FromImperialPint(s.Value() / 20)
	case ImperialCup:
		return FromImperialPint(s.Value() / 2)
	case ImperialPint:
		return FromImperialPint(s.Value())
	case ImperialQuart:
		return FromImperialPint(s.Value() * 2)
	case ImperialGallon:
		return FromImperialPint(s.Value() * 8)
	}
}

func (s *volume) ToImperialQuart() Volume {
	switch s.Unit() {
	default:
		return FromImperialQuart(s.Value() / 1137)
	case Litre:
		return FromImperialQuart(s.Value() / 1.137)
	case USfluidOunce:
		return FromImperialQuart(s.Value() / 38.43)
	case USlegalCup:
		return FromImperialQuart(s.Value() / 4.736)
	case USliquidPint:
		return FromImperialQuart(s.Value() / 2.402)
	case USLiquidQuart:
		return FromImperialQuart(s.Value() / 1.201)
	case USLiquidGallon:
		return FromImperialQuart(s.Value() * 3.331)
	case ImperialFluidOunce:
		return FromImperialQuart(s.Value() / 40)
	case ImperialCup:
		return FromImperialQuart(s.Value() / 4)
	case ImperialPint:
		return FromImperialQuart(s.Value() / 2)
	case ImperialQuart:
		return FromImperialQuart(s.Value())
	case ImperialGallon:
		return FromImperialQuart(s.Value() * 4)
	}
}

func (s *volume) ToImperialGallon() Volume {
	switch s.Unit() {
	default:
		return FromImperialGallon(s.Value() / 4546)
	case Litre:
		return FromImperialGallon(s.Value() / 4.546)
	case USfluidOunce:
		return FromImperialGallon(s.Value() / 154)
	case USlegalCup:
		return FromImperialGallon(s.Value() / 18.942)
	case USliquidPint:
		return FromImperialGallon(s.Value() / 9.608)
	case USLiquidQuart:
		return FromImperialGallon(s.Value() / 4.804)
	case USLiquidGallon:
		return FromImperialGallon(s.Value() / 1.201)
	case ImperialFluidOunce:
		return FromImperialGallon(s.Value() / 160)
	case ImperialCup:
		return FromImperialGallon(s.Value() / 16)
	case ImperialPint:
		return FromImperialGallon(s.Value() / 8)
	case ImperialQuart:
		return FromImperialGallon(s.Value() / 4)
	case ImperialGallon:
		return FromImperialGallon(s.Value())
	}
}