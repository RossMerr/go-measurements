package measurements_test

import (
	"reflect"
	"testing"

	"github.com/RossMerr/go-measurements"
)

func Test_volume_ToMilliliter(t *testing.T) {
	type fields struct {
		unit  measurements.VolumeType
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   measurements.Volume
	}{
		{
			name: "Milliliter to Milliliter",
			fields: fields{
				unit:  measurements.Milliliter,
				value: 10,
			},
			want: measurements.FromMilliliter(10),
		},
		{
			name: "Litre to Milliliter",
			fields: fields{
				unit:  measurements.Litre,
				value: 0.01,
			},
			want: measurements.FromMilliliter(10),
		},
		{
			name: "USfluidOunce to Milliliter",
			fields: fields{
				unit:  measurements.USfluidOunce,
				value: 0.33814,
			},
			want: measurements.FromMilliliter(10),
		},
		{
			name: "USlegalCup to Milliliter",
			fields: fields{
				unit:  measurements.USlegalCup,
				value: 0.0416667,
			},
			want: measurements.FromMilliliter(10),
		},
		{
			name: "USliquidPint to Milliliter",
			fields: fields{
				unit:  measurements.USliquidPint,
				value: 0.0211338,
			},
			want: measurements.FromMilliliter(10),
		},
		{
			name: "USLiquidQuart to Milliliter",
			fields: fields{
				unit:  measurements.USLiquidQuart,
				value: 0.0105669,
			},
			want: measurements.FromMilliliter(10),
		},
		{
			name: "USLiquidGallon to Milliliter",
			fields: fields{
				unit:  measurements.USLiquidGallon,
				value: 0.00264172,
			},
			want: measurements.FromMilliliter(10),
		},
		{
			name: "ImperialFluidOunce to Milliliter",
			fields: fields{
				unit:  measurements.ImperialFluidOunce,
				value: 0.351951,
			},
			want: measurements.FromMilliliter(10),
		},
		{
			name: "ImperialCup to Milliliter",
			fields: fields{
				unit:  measurements.ImperialCup,
				value: 0.0351951,
			},
			want: measurements.FromMilliliter(10),
		},
		{
			name: "ImperialPint to Milliliter",
			fields: fields{
				unit:  measurements.ImperialPint,
				value: 0.0175975,
			},
			want: measurements.FromMilliliter(10),
		},
		{
			name: "ImperialQuart to Milliliter",
			fields: fields{
				unit:  measurements.ImperialQuart,
				value: 0.00879877,
			},
			want: measurements.FromMilliliter(10),
		},
		{
			name: "ImperialGallon to Milliliter",
			fields: fields{
				unit:  measurements.ImperialGallon,
				value: 0.00219969,
			},
			want: measurements.FromMilliliter(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := measurements.NewVolume(tt.fields.unit, tt.fields.value)
			if got := s.ToMilliliter(); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("ToMilliliter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_volume_ToLiter(t *testing.T) {
	type fields struct {
		unit  measurements.VolumeType
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   measurements.Volume
	}{
		{
			name: "Milliliter to Litre",
			fields: fields{
				unit:  measurements.Milliliter,
				value: 10000,
			},
			want: measurements.FromLiter(10),
		},
		{
			name: "Litre to Litre",
			fields: fields{
				unit:  measurements.Litre,
				value: 10,
			},
			want: measurements.FromLiter(10),
		},
		{
			name: "USfluidOunce to Litre",
			fields: fields{
				unit:  measurements.USfluidOunce,
				value: 338.14,
			},
			want: measurements.FromLiter(10),
		},
		{
			name: "USlegalCup to Litre",
			fields: fields{
				unit:  measurements.USlegalCup,
				value: 41.6667,
			},
			want: measurements.FromLiter(10),
		},
		{
			name: "USliquidPint to Litre",
			fields: fields{
				unit:  measurements.USliquidPint,
				value: 21.1338,
			},
			want: measurements.FromLiter(10),
		},
		{
			name: "USLiquidQuart to Litre",
			fields: fields{
				unit:  measurements.USLiquidQuart,
				value: 10.5669,
			},
			want: measurements.FromLiter(10),
		},
		{
			name: "USLiquidGallon to Litre",
			fields: fields{
				unit:  measurements.USLiquidGallon,
				value: 2.64172,
			},
			want: measurements.FromLiter(10),
		},
		{
			name: "ImperialFluidOunce to Litre",
			fields: fields{
				unit:  measurements.ImperialFluidOunce,
				value: 351.951,
			},
			want: measurements.FromLiter(10),
		},
		{
			name: "ImperialCup to Litre",
			fields: fields{
				unit:  measurements.ImperialCup,
				value: 35.1951,
			},
			want: measurements.FromLiter(10),
		},
		{
			name: "ImperialPint to Litre",
			fields: fields{
				unit:  measurements.ImperialPint,
				value: 17.5975,
			},
			want: measurements.FromLiter(10),
		},
		{
			name: "ImperialQuart to Litre",
			fields: fields{
				unit:  measurements.ImperialQuart,
				value: 8.79877,
			},
			want: measurements.FromLiter(10),
		},
		{
			name: "ImperialGallon to Litre",
			fields: fields{
				unit:  measurements.ImperialGallon,
				value: 2.19969,
			},
			want: measurements.FromLiter(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := measurements.NewVolume(tt.fields.unit, tt.fields.value)
			if got := s.ToLiter(); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("ToLiter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_volume_ToUSfluidOunce(t *testing.T) {
	type fields struct {
		unit  measurements.VolumeType
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   measurements.Volume
	}{
		{
			name: "Milliliter to USfluidOunce",
			fields: fields{
				unit:  measurements.Milliliter,
				value: 295.735,
			},
			want: measurements.FromUSfluidOunce(10),
		},
		{
			name: "Litre to USfluidOunce",
			fields: fields{
				unit:  measurements.Litre,
				value: 0.295735,
			},
			want: measurements.FromUSfluidOunce(10),
		},
		{
			name: "USfluidOunce to USfluidOunce",
			fields: fields{
				unit:  measurements.USfluidOunce,
				value: 10,
			},
			want: measurements.FromUSfluidOunce(10),
		},
		{
			name: "USlegalCup to USfluidOunce",
			fields: fields{
				unit:  measurements.USlegalCup,
				value: 1.23223,
			},
			want: measurements.FromUSfluidOunce(10),
		},
		{
			name: "USliquidPint to USfluidOunce",
			fields: fields{
				unit:  measurements.USliquidPint,
				value: 0.625,
			},
			want: measurements.FromUSfluidOunce(10),
		},
		{
			name: "USLiquidQuart to USfluidOunce",
			fields: fields{
				unit:  measurements.USLiquidQuart,
				value: 0.3125,
			},
			want: measurements.FromUSfluidOunce(10),
		},
		{
			name: "USLiquidGallon to USfluidOunce",
			fields: fields{
				unit:  measurements.USLiquidGallon,
				value: 0.078125,
			},
			want: measurements.FromUSfluidOunce(10),
		},
		{
			name: "ImperialFluidOunce to USfluidOunce",
			fields: fields{
				unit:  measurements.ImperialFluidOunce,
				value: 10.4084,
			},
			want: measurements.FromUSfluidOunce(10),
		},
		{
			name: "ImperialCup to USfluidOunce",
			fields: fields{
				unit:  measurements.ImperialCup,
				value: 1.04084,
			},
			want: measurements.FromUSfluidOunce(10),
		},
		{
			name: "ImperialPint to USfluidOunce",
			fields: fields{
				unit:  measurements.ImperialPint,
				value: 0.520421,
			},
			want: measurements.FromUSfluidOunce(10),
		},
		{
			name: "ImperialQuart to USfluidOunce",
			fields: fields{
				unit:  measurements.ImperialQuart,
				value: 0.260211,
			},
			want: measurements.FromUSfluidOunce(10),
		},
		{
			name: "ImperialGallon to USfluidOunce",
			fields: fields{
				unit:  measurements.ImperialGallon,
				value: 0.0650527,
			},
			want: measurements.FromUSfluidOunce(10.02),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := measurements.NewVolume(tt.fields.unit, tt.fields.value)
			if got := s.ToUSfluidOunce(); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("ToUSfluidOunce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_volume_ToUSlegalCup(t *testing.T) {
	type fields struct {
		unit  measurements.VolumeType
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   measurements.Volume
	}{
		{
			name: "Milliliter to USlegalCup",
			fields: fields{
				unit:  measurements.Milliliter,
				value: 2400,
			},
			want: measurements.FromUSlegalCup(10),
		},
		{
			name: "Litre to USlegalCup",
			fields: fields{
				unit:  measurements.Litre,
				value: 2.4,
			},
			want: measurements.FromUSlegalCup(10),
		},
		{
			name: "USfluidOunce to USlegalCup",
			fields: fields{
				unit:  measurements.USfluidOunce,
				value: 81.1537,
			},
			want: measurements.FromUSlegalCup(10),
		},
		{
			name: "USlegalCup to USlegalCup",
			fields: fields{
				unit:  measurements.USlegalCup,
				value: 10,
			},
			want: measurements.FromUSlegalCup(10),
		},
		{
			name: "USliquidPint to USlegalCup",
			fields: fields{
				unit:  measurements.USliquidPint,
				value: 5.0721,
			},
			want: measurements.FromUSlegalCup(10),
		},
		{
			name: "USLiquidQuart to USlegalCup",
			fields: fields{
				unit:  measurements.USLiquidQuart,
				value: 2.53605,
			},
			want: measurements.FromUSlegalCup(10),
		},
		{
			name: "USLiquidGallon to USlegalCup",
			fields: fields{
				unit:  measurements.USLiquidGallon,
				value: 0.634013,
			},
			want: measurements.FromUSlegalCup(10),
		},
		{
			name: "ImperialFluidOunce to USlegalCup",
			fields: fields{
				unit:  measurements.ImperialFluidOunce,
				value: 84.4682,
			},
			want: measurements.FromUSlegalCup(10),
		},
		{
			name: "ImperialCup to USlegalCup",
			fields: fields{
				unit:  measurements.ImperialCup,
				value: 8.44682,
			},
			want: measurements.FromUSlegalCup(10),
		},
		{
			name: "ImperialPint to USlegalCup",
			fields: fields{
				unit:  measurements.ImperialPint,
				value: 4.22341,
			},
			want: measurements.FromUSlegalCup(10),
		},
		{
			name: "ImperialQuart to USlegalCup",
			fields: fields{
				unit:  measurements.ImperialQuart,
				value: 2.1117,
			},
			want: measurements.FromUSlegalCup(10),
		},
		{
			name: "ImperialGallon to USlegalCup",
			fields: fields{
				unit:  measurements.ImperialGallon,
				value: 0.527926,
			},
			want: measurements.FromUSlegalCup(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := measurements.NewVolume(tt.fields.unit, tt.fields.value)
			if got := s.ToUSlegalCup(); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("ToUSlegalCup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_volume_ToUSliquidPint(t *testing.T) {
	type fields struct {
		unit  measurements.VolumeType
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   measurements.Volume
	}{
		{
			name: "Milliliter to USliquidPint",
			fields: fields{
				unit:  measurements.Milliliter,
				value: 4731.76,
			},
			want: measurements.FromUSliquidPint(10),
		},
		{
			name: "Litre to USliquidPint",
			fields: fields{
				unit:  measurements.Litre,
				value: 4.73176,
			},
			want: measurements.FromUSliquidPint(10),
		},
		{
			name: "USfluidOunce to USliquidPint",
			fields: fields{
				unit:  measurements.USfluidOunce,
				value: 160,
			},
			want: measurements.FromUSliquidPint(10),
		},
		{
			name: "USlegalCup to USliquidPint",
			fields: fields{
				unit:  measurements.USlegalCup,
				value: 19.7157,
			},
			want: measurements.FromUSliquidPint(10),
		},
		{
			name: "USliquidPint to USliquidPint",
			fields: fields{
				unit:  measurements.USliquidPint,
				value: 10,
			},
			want: measurements.FromUSliquidPint(10),
		},
		{
			name: "USLiquidQuart to USliquidPint",
			fields: fields{
				unit:  measurements.USLiquidQuart,
				value: 5,
			},
			want: measurements.FromUSliquidPint(10),
		},
		{
			name: "USLiquidGallon to USliquidPint",
			fields: fields{
				unit:  measurements.USLiquidGallon,
				value: 1.25,
			},
			want: measurements.FromUSliquidPint(10),
		},
		{
			name: "ImperialFluidOunce to USliquidPint",
			fields: fields{
				unit:  measurements.ImperialFluidOunce,
				value: 166.535,
			},
			want: measurements.FromUSliquidPint(10),
		},
		{
			name: "ImperialCup to USliquidPint",
			fields: fields{
				unit:  measurements.ImperialCup,
				value: 16.6535,
			},
			want: measurements.FromUSliquidPint(10),
		},
		{
			name: "ImperialPint to USliquidPint",
			fields: fields{
				unit:  measurements.ImperialPint,
				value: 8.32674,
			},
			want: measurements.FromUSliquidPint(10),
		},
		{
			name: "ImperialQuart to USliquidPint",
			fields: fields{
				unit:  measurements.ImperialQuart,
				value: 4.16337,
			},
			want: measurements.FromUSliquidPint(10),
		},
		{
			name: "ImperialGallon to USliquidPint",
			fields: fields{
				unit:  measurements.ImperialGallon,
				value: 1.04084,
			},
			want: measurements.FromUSliquidPint(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := measurements.NewVolume(tt.fields.unit, tt.fields.value)
			if got := s.ToUSliquidPint(); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("ToUSliquidPint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_volume_ToUSLiquidQuart(t *testing.T) {
	type fields struct {
		unit  measurements.VolumeType
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   measurements.Volume
	}{
		{
			name: "Milliliter to USLiquidQuart",
			fields: fields{
				unit:  measurements.Milliliter,
				value: 9463.53,
			},
			want: measurements.FromUSLiquidQuart(10),
		},
		{
			name: "Litre to USLiquidQuart",
			fields: fields{
				unit:  measurements.Litre,
				value: 9.4635299994212,
			},
			want: measurements.FromUSLiquidQuart(10),
		},
		{
			name: "USfluidOunce to USLiquidQuart",
			fields: fields{
				unit:  measurements.USfluidOunce,
				value: 320,
			},
			want: measurements.FromUSLiquidQuart(10),
		},
		{
			name: "USlegalCup to USLiquidQuart",
			fields: fields{
				unit:  measurements.USlegalCup,
				value: 39.4314,
			},
			want: measurements.FromUSLiquidQuart(10),
		},
		{
			name: "USliquidPint to USLiquidQuart",
			fields: fields{
				unit:  measurements.USliquidPint,
				value: 20.00001382,
			},
			want: measurements.FromUSLiquidQuart(10),
		},
		{
			name: "USLiquidQuart to USLiquidQuart",
			fields: fields{
				unit:  measurements.USLiquidQuart,
				value: 10,
			},
			want: measurements.FromUSLiquidQuart(10),
		},
		{
			name: "USLiquidGallon to USLiquidQuart",
			fields: fields{
				unit:  measurements.USLiquidGallon,
				value: 2.5,
			},
			want: measurements.FromUSLiquidQuart(10),
		},
		{
			name: "ImperialFluidOunce to USLiquidQuart",
			fields: fields{
				unit:  measurements.ImperialFluidOunce,
				value: 333.07,
			},
			want: measurements.FromUSLiquidQuart(10),
		},
		{
			name: "ImperialCup to USLiquidQuart",
			fields: fields{
				unit:  measurements.ImperialCup,
				value: 33.307,
			},
			want: measurements.FromUSLiquidQuart(10),
		},
		{
			name: "ImperialPint to USLiquidQuart",
			fields: fields{
				unit:  measurements.ImperialPint,
				value: 16.6535,
			},
			want: measurements.FromUSLiquidQuart(10),
		},
		{
			name: "ImperialQuart to USLiquidQuart",
			fields: fields{
				unit:  measurements.ImperialQuart,
				value: 8.32674,
			},
			want: measurements.FromUSLiquidQuart(10),
		},
		{
			name: "ImperialGallon to USLiquidQuart",
			fields: fields{
				unit:  measurements.ImperialGallon,
				value: 2.08169,
			},
			want: measurements.FromUSLiquidQuart(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := measurements.NewVolume(tt.fields.unit, tt.fields.value)
			if got := s.ToUSLiquidQuart(); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("ToUSLiquidQuart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_volume_ToUSLiquidGallon(t *testing.T) {
	type fields struct {
		unit  measurements.VolumeType
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   measurements.Volume
	}{
		{
			name: "Milliliter to USLiquidGallon",
			fields: fields{
				unit:  measurements.Milliliter,
				value: 37854.1,
			},
			want: measurements.FromUSLiquidGallon(10),
		},
		{
			name: "Litre to USLiquidGallon",
			fields: fields{
				unit:  measurements.Litre,
				value: 37.8541,
			},
			want: measurements.FromUSLiquidGallon(10),
		},
		{
			name: "USfluidOunce to USLiquidGallon",
			fields: fields{
				unit:  measurements.USfluidOunce,
				value: 1280,
			},
			want: measurements.FromUSLiquidGallon(10),
		},
		{
			name: "USlegalCup to USLiquidGallon",
			fields: fields{
				unit:  measurements.USlegalCup,
				value: 157.725,
			},
			want: measurements.FromUSLiquidGallon(10),
		},
		{
			name: "USliquidPint to USLiquidGallon",
			fields: fields{
				unit:  measurements.USliquidPint,
				value: 80,
			},
			want: measurements.FromUSLiquidGallon(10),
		},
		{
			name: "USLiquidQuart to USLiquidGallon",
			fields: fields{
				unit:  measurements.USLiquidQuart,
				value: 40,
			},
			want: measurements.FromUSLiquidGallon(10),
		},
		{
			name: "USLiquidGallon to USLiquidGallon",
			fields: fields{
				unit:  measurements.USLiquidGallon,
				value: 10,
			},
			want: measurements.FromUSLiquidGallon(10),
		},
		{
			name: "ImperialFluidOunce to USLiquidGallon",
			fields: fields{
				unit:  measurements.ImperialFluidOunce,
				value: 1332.28,
			},
			want: measurements.FromUSLiquidGallon(10.02),
		},
		{
			name: "ImperialCup to USLiquidGallon",
			fields: fields{
				unit:  measurements.ImperialCup,
				value: 133.228,
			},
			want: measurements.FromUSLiquidGallon(10),
		},
		{
			name: "ImperialPint to USLiquidGallon",
			fields: fields{
				unit:  measurements.ImperialPint,
				value: 66.6139,
			},
			want: measurements.FromUSLiquidGallon(10),
		},
		{
			name: "ImperialQuart to USLiquidGallon",
			fields: fields{
				unit:  measurements.ImperialQuart,
				value: 33.307,
			},
			want: measurements.FromUSLiquidGallon(10),
		},
		{
			name: "ImperialGallon to USLiquidGallon",
			fields: fields{
				unit:  measurements.ImperialGallon,
				value: 8.32674,
			},
			want: measurements.FromUSLiquidGallon(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := measurements.NewVolume(tt.fields.unit, tt.fields.value)
			if got := s.ToUSLiquidGallon(); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("ToUSLiquidGallon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_volume_ToImperialFluidOunce(t *testing.T) {
	type fields struct {
		unit  measurements.VolumeType
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   measurements.Volume
	}{
		{
			name: "Milliliter to ImperialFluidOunce",
			fields: fields{
				unit:  measurements.Milliliter,
				value: 284.131,
			},
			want: measurements.FromImperialFluidOunce(10),
		},
		{
			name: "Litre to ImperialFluidOunce",
			fields: fields{
				unit:  measurements.Litre,
				value: 0.284131,
			},
			want: measurements.FromImperialFluidOunce(10),
		},
		{
			name: "USfluidOunce to ImperialFluidOunce",
			fields: fields{
				unit:  measurements.USfluidOunce,
				value: 9.6076,
			},
			want: measurements.FromImperialFluidOunce(10),
		},
		{
			name: "USlegalCup to ImperialFluidOunce",
			fields: fields{
				unit:  measurements.USlegalCup,
				value: 1.18388,
			},
			want: measurements.FromImperialFluidOunce(10),
		},
		{
			name: "USliquidPint to ImperialFluidOunce",
			fields: fields{
				unit:  measurements.USliquidPint,
				value: 0.600475,
			},
			want: measurements.FromImperialFluidOunce(10),
		},
		{
			name: "USLiquidQuart to ImperialFluidOunce",
			fields: fields{
				unit:  measurements.USLiquidQuart,
				value: 0.300237,
			},
			want: measurements.FromImperialFluidOunce(10),
		},
		{
			name: "USLiquidGallon to ImperialFluidOunce",
			fields: fields{
				unit:  measurements.USLiquidGallon,
				value: 0.0750594,
			},
			want: measurements.FromImperialFluidOunce(9.98),
		},
		{
			name: "ImperialFluidOunce to ImperialFluidOunce",
			fields: fields{
				unit:  measurements.ImperialFluidOunce,
				value: 10,
			},
			want: measurements.FromImperialFluidOunce(10),
		},
		{
			name: "ImperialCup to ImperialFluidOunce",
			fields: fields{
				unit:  measurements.ImperialCup,
				value: 1,
			},
			want: measurements.FromImperialFluidOunce(10),
		},
		{
			name: "ImperialPint to ImperialFluidOunce",
			fields: fields{
				unit:  measurements.ImperialPint,
				value: 0.5,
			},
			want: measurements.FromImperialFluidOunce(10),
		},
		{
			name: "ImperialQuart to ImperialFluidOunce",
			fields: fields{
				unit:  measurements.ImperialQuart,
				value: 0.25,
			},
			want: measurements.FromImperialFluidOunce(10),
		},
		{
			name: "ImperialGallon to ImperialFluidOunce",
			fields: fields{
				unit:  measurements.ImperialGallon,
				value: 0.0625,
			},
			want: measurements.FromImperialFluidOunce(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := measurements.NewVolume(tt.fields.unit, tt.fields.value)
			if got := s.ToImperialFluidOunce(); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("ToImperialFluidOunce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_volume_ToImperialCup(t *testing.T) {
	type fields struct {
		unit  measurements.VolumeType
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   measurements.Volume
	}{
		{
			name: "Milliliter to ImperialCup",
			fields: fields{
				unit:  measurements.Milliliter,
				value: 2841.31,
			},
			want: measurements.FromImperialCup(10),
		},
		{
			name: "Litre to ImperialCup",
			fields: fields{
				unit:  measurements.Litre,
				value: 2.84131,
			},
			want: measurements.FromImperialCup(10),
		},
		{
			name: "USfluidOunce to ImperialCup",
			fields: fields{
				unit:  measurements.USfluidOunce,
				value: 96.076,
			},
			want: measurements.FromImperialCup(10),
		},
		{
			name: "USlegalCup to ImperialCup",
			fields: fields{
				unit:  measurements.USlegalCup,
				value: 11.8388,
			},
			want: measurements.FromImperialCup(10),
		},
		{
			name: "USliquidPint to ImperialCup",
			fields: fields{
				unit:  measurements.USliquidPint,
				value: 6.00475,
			},
			want: measurements.FromImperialCup(10),
		},
		{
			name: "USLiquidQuart to ImperialCup",
			fields: fields{
				unit:  measurements.USLiquidQuart,
				value: 3.00237,
			},
			want: measurements.FromImperialCup(10),
		},
		{
			name: "USLiquidGallon to ImperialCup",
			fields: fields{
				unit:  measurements.USLiquidGallon,
				value: 0.750594,
			},
			want: measurements.FromImperialCup(10),
		},
		{
			name: "ImperialFluidOunce to ImperialCup",
			fields: fields{
				unit:  measurements.ImperialFluidOunce,
				value: 100,
			},
			want: measurements.FromImperialCup(10),
		},
		{
			name: "ImperialCup to ImperialCup",
			fields: fields{
				unit:  measurements.ImperialCup,
				value: 1,
			},
			want: measurements.FromImperialCup(1),
		},
		{
			name: "ImperialPint to ImperialCup",
			fields: fields{
				unit:  measurements.ImperialPint,
				value: 5,
			},
			want: measurements.FromImperialCup(10),
		},
		{
			name: "ImperialQuart to ImperialCup",
			fields: fields{
				unit:  measurements.ImperialQuart,
				value: 2.5,
			},
			want: measurements.FromImperialCup(10),
		},
		{
			name: "ImperialGallon to ImperialCup",
			fields: fields{
				unit:  measurements.ImperialGallon,
				value: 0.625,
			},
			want: measurements.FromImperialCup(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := measurements.NewVolume(tt.fields.unit, tt.fields.value)
			if got := s.ToImperialCup(); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("ToImperialCup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_volume_ToImperialPint(t *testing.T) {
	type fields struct {
		unit  measurements.VolumeType
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   measurements.Volume
	}{
		{
			name: "Milliliter to ImperialPint",
			fields: fields{
				unit:  measurements.Milliliter,
				value: 5682.61,
			},
			want: measurements.FromImperialPint(10),
		},
		{
			name: "Litre to ImperialPint",
			fields: fields{
				unit:  measurements.Litre,
				value: 5.68261,
			},
			want: measurements.FromImperialPint(10),
		},
		{
			name: "USfluidOunce to ImperialPint",
			fields: fields{
				unit:  measurements.USfluidOunce,
				value: 192.152,
			},
			want: measurements.FromImperialPint(10),
		},
		{
			name: "USlegalCup to ImperialPint",
			fields: fields{
				unit:  measurements.USlegalCup,
				value: 23.6776,
			},
			want: measurements.FromImperialPint(10),
		},
		{
			name: "USliquidPint to ImperialPint",
			fields: fields{
				unit:  measurements.USliquidPint,
				value: 12.0095,
			},
			want: measurements.FromImperialPint(10),
		},
		{
			name: "USLiquidQuart to ImperialPint",
			fields: fields{
				unit:  measurements.USLiquidQuart,
				value: 6.00475,
			},
			want: measurements.FromImperialPint(10),
		},
		{
			name: "USLiquidGallon to ImperialPint",
			fields: fields{
				unit:  measurements.USLiquidGallon,
				value: 1.50119,
			},
			want: measurements.FromImperialPint(10),
		},
		{
			name: "ImperialFluidOunce to ImperialPint",
			fields: fields{
				unit:  measurements.ImperialFluidOunce,
				value: 200,
			},
			want: measurements.FromImperialPint(10),
		},
		{
			name: "ImperialCup to ImperialPint",
			fields: fields{
				unit:  measurements.ImperialCup,
				value: 20,
			},
			want: measurements.FromImperialPint(10),
		},
		{
			name: "ImperialPint to ImperialPint",
			fields: fields{
				unit:  measurements.ImperialPint,
				value: 10,
			},
			want: measurements.FromImperialPint(10),
		},
		{
			name: "ImperialQuart to ImperialPint",
			fields: fields{
				unit:  measurements.ImperialQuart,
				value: 5,
			},
			want: measurements.FromImperialPint(10),
		},
		{
			name: "ImperialGallon to ImperialPint",
			fields: fields{
				unit:  measurements.ImperialGallon,
				value: 1.25,
			},
			want: measurements.FromImperialPint(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := measurements.NewVolume(tt.fields.unit, tt.fields.value)
			if got := s.ToImperialPint(); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("ToImperialPint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_volume_ToImperialQuart(t *testing.T) {
	type fields struct {
		unit  measurements.VolumeType
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   measurements.Volume
	}{
		{
			name: "Milliliter to ImperialQuart",
			fields: fields{
				unit:  measurements.Milliliter,
				value: 11365.2,
			},
			want: measurements.FromImperialQuart(10),
		},
		{
			name: "Litre to ImperialQuart",
			fields: fields{
				unit:  measurements.Litre,
				value: 11.3652,
			},
			want: measurements.FromImperialQuart(10),
		},
		{
			name: "USfluidOunce to ImperialQuart",
			fields: fields{
				unit:  measurements.USfluidOunce,
				value: 384.304,
			},
			want: measurements.FromImperialQuart(10),
		},
		{
			name: "USlegalCup to ImperialQuart",
			fields: fields{
				unit:  measurements.USlegalCup,
				value: 47.3551,
			},
			want: measurements.FromImperialQuart(10),
		},
		{
			name: "USliquidPint to ImperialQuart",
			fields: fields{
				unit:  measurements.USliquidPint,
				value: 24.019,
			},
			want: measurements.FromImperialQuart(10),
		},
		{
			name: "USLiquidQuart to ImperialQuart",
			fields: fields{
				unit:  measurements.USLiquidQuart,
				value: 12.0095,
			},
			want: measurements.FromImperialQuart(10),
		},
		{
			name: "USLiquidGallon to ImperialQuart",
			fields: fields{
				unit:  measurements.USLiquidGallon,
				value: 3.00237,
			},
			want: measurements.FromImperialQuart(10),
		},
		{
			name: "ImperialFluidOunce to ImperialQuart",
			fields: fields{
				unit:  measurements.ImperialFluidOunce,
				value: 400,
			},
			want: measurements.FromImperialQuart(10),
		},
		{
			name: "ImperialCup to ImperialQuart",
			fields: fields{
				unit:  measurements.ImperialCup,
				value: 40,
			},
			want: measurements.FromImperialQuart(10),
		},
		{
			name: "ImperialPint to ImperialQuart",
			fields: fields{
				unit:  measurements.ImperialPint,
				value: 20,
			},
			want: measurements.FromImperialQuart(10),
		},
		{
			name: "ImperialQuart to ImperialQuart",
			fields: fields{
				unit:  measurements.ImperialQuart,
				value: 10,
			},
			want: measurements.FromImperialQuart(10),
		},
		{
			name: "ImperialGallon to ImperialQuart",
			fields: fields{
				unit:  measurements.ImperialGallon,
				value: 2.5,
			},
			want: measurements.FromImperialQuart(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := measurements.NewVolume(tt.fields.unit, tt.fields.value)
			if got := s.ToImperialQuart(); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("ToImperialQuart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_volume_ToImperialGallon(t *testing.T) {
	type fields struct {
		unit  measurements.VolumeType
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   measurements.Volume
	}{
		{
			name: "Milliliter to ImperialGallon",
			fields: fields{
				unit:  measurements.Milliliter,
				value: 45460.9,
			},
			want: measurements.FromImperialGallon(10),
		},
		{
			name: "Litre to ImperialGallon",
			fields: fields{
				unit:  measurements.Litre,
				value: 45.4609,
			},
			want: measurements.FromImperialGallon(10),
		},
		{
			name: "USfluidOunce to ImperialGallon",
			fields: fields{
				unit:  measurements.USfluidOunce,
				value: 1537.22,
			},
			want: measurements.FromImperialGallon(9.98),
		},
		{
			name: "USlegalCup to ImperialGallon",
			fields: fields{
				unit:  measurements.USlegalCup,
				value: 189.42,
			},
			want: measurements.FromImperialGallon(10),
		},
		{
			name: "USliquidPint to ImperialGallon",
			fields: fields{
				unit:  measurements.USliquidPint,
				value: 96.076,
			},
			want: measurements.FromImperialGallon(10),
		},
		{
			name: "USLiquidQuart to ImperialGallon",
			fields: fields{
				unit:  measurements.USLiquidQuart,
				value: 48.038,
			},
			want: measurements.FromImperialGallon(10),
		},
		{
			name: "USLiquidGallon to ImperialGallon",
			fields: fields{
				unit:  measurements.USLiquidGallon,
				value: 12.0095,
			},
			want: measurements.FromImperialGallon(10),
		},
		{
			name: "ImperialFluidOunce to ImperialGallon",
			fields: fields{
				unit:  measurements.ImperialFluidOunce,
				value: 1600,
			},
			want: measurements.FromImperialGallon(10),
		},
		{
			name: "ImperialCup to ImperialGallon",
			fields: fields{
				unit:  measurements.ImperialCup,
				value: 160,
			},
			want: measurements.FromImperialGallon(10),
		},
		{
			name: "ImperialPint to ImperialGallon",
			fields: fields{
				unit:  measurements.ImperialPint,
				value: 80,
			},
			want: measurements.FromImperialGallon(10),
		},
		{
			name: "ImperialQuart to ImperialGallon",
			fields: fields{
				unit:  measurements.ImperialQuart,
				value: 40,
			},
			want: measurements.FromImperialGallon(10),
		},
		{
			name: "ImperialGallon to ImperialGallon",
			fields: fields{
				unit:  measurements.ImperialGallon,
				value: 10,
			},
			want: measurements.FromImperialGallon(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := measurements.NewVolume(tt.fields.unit, tt.fields.value)
			if got := s.ToImperialGallon(); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("ToImperialGallon() = %v, want %v", got, tt.want)
			}
		})
	}
}
