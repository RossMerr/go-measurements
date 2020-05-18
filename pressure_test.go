package measurements_test

import (
	"reflect"
	"testing"

	"github.com/RossMerr/go-measurements"
)

func Test_pressure_ToTorr(t *testing.T) {
	type fields struct {
		unit  measurements.PressureUnit
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   measurements.Pressure
	}{
		{
			name: "Torr to Torr",
			fields: fields{
				unit:  measurements.Torr,
				value: 10,
			},
			want: measurements.FromTorr(10),
		},
		{
			name: "Bar to Torr",
			fields: fields{
				unit:  measurements.Bar,
				value: 0.0133322,
			},
			want: measurements.FromTorr(10),
		},
		{
			name: "Pascal to Torr",
			fields: fields{
				unit:  measurements.Pascal,
				value: 1333.22,
			},
			want: measurements.FromTorr(10.02),
		},
		{
			name: "PoundForcePerSquareInch to Torr",
			fields: fields{
				unit:  measurements.PoundForcePerSquareInch,
				value: 0.193368,
			},
			want: measurements.FromTorr(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := measurements.NewPressure(tt.fields.unit, tt.fields.value)
			if got := s.ToTorr(); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("ToTorr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pressure_ToBar(t *testing.T) {
	type fields struct {
		unit  measurements.PressureUnit
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   measurements.Pressure
	}{
		{
			name: "Torr to Bar",
			fields: fields{
				unit:  measurements.Torr,
				value: 7500.62,
			},
			want: measurements.FromBar(10),
		},
		{
			name: "Bar to Bar",
			fields: fields{
				unit:  measurements.Bar,
				value: 10,
			},
			want: measurements.FromBar(10),
		},
		{
			name: "Pascal to Bar",
			fields: fields{
				unit:  measurements.Pascal,
				value: 1e+6,
			},
			want: measurements.FromBar(10),
		},
		{
			name: "PoundForcePerSquareInch to Bar",
			fields: fields{
				unit:  measurements.PoundForcePerSquareInch,
				value: 145.038,
			},
			want: measurements.FromBar(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := measurements.NewPressure(tt.fields.unit, tt.fields.value)
			if got := s.ToBar(); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("ToBar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pressure_ToPascal(t *testing.T) {
	type fields struct {
		unit  measurements.PressureUnit
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   measurements.Pressure
	}{
		{
			name: "Torr to Pascal",
			fields: fields{
				unit:  measurements.Torr,
				value: 0.0750062,
			},
			want: measurements.FromPascal(9.98),
		},
		{
			name: "Bar to Pascal",
			fields: fields{
				unit:  measurements.Bar,
				value: 1e-4,
			},
			want: measurements.FromPascal(10),
		},
		{
			name: "Pascal to Pascal",
			fields: fields{
				unit:  measurements.Pascal,
				value: 10,
			},
			want: measurements.FromPascal(10),
		},
		{
			name: "PoundForcePerSquareInch to Pascal",
			fields: fields{
				unit:  measurements.PoundForcePerSquareInch,
				value: 0.00145038,
			},
			want: measurements.FromPascal(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := measurements.NewPressure(tt.fields.unit, tt.fields.value)
			if got := s.ToPascal(); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("ToPascal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pressure_ToPoundForcePerSquareInch(t *testing.T) {
	type fields struct {
		unit  measurements.PressureUnit
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   measurements.Pressure
	}{
		{
			name: "Torr to PoundForcePerSquareInch",
			fields: fields{
				unit:  measurements.Torr,
				value: 517.149,
			},
			want: measurements.FromPoundForcePerSquareInch(10),
		},
		{
			name: "Bar to PoundForcePerSquareInch",
			fields: fields{
				unit:  measurements.Bar,
				value: 0.689476,
			},
			want: measurements.FromPoundForcePerSquareInch(10),
		},
		{
			name: "Pascal to PoundForcePerSquareInch",
			fields: fields{
				unit:  measurements.Pascal,
				value: 68947.6,
			},
			want: measurements.FromPoundForcePerSquareInch(10),
		},
		{
			name: "PoundForcePerSquareInch to PoundForcePerSquareInch",
			fields: fields{
				unit:  measurements.PoundForcePerSquareInch,
				value: 10,
			},
			want: measurements.FromPoundForcePerSquareInch(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := measurements.NewPressure(tt.fields.unit, tt.fields.value)
			if got := s.ToPoundForcePerSquareInch(); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("ToPoundForcePerSquareInch() = %v, want %v", got, tt.want)
			}
		})
	}
}
