package measurements_test

import (
	"reflect"
	"testing"

	"github.com/RossMerr/herms/src/controller/measurements"
)

func Test_mass_ToKilogram(t *testing.T) {
	type fields struct {
		unit  measurements.MassUnit
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   measurements.Mass
	}{
		{
			name: "Kilogram to Kilogram",
			fields: fields{
				unit: measurements.Kilogram,
				value: 10,
			},
			want: measurements.FromKilogram(10),
		},
		{
			name: "Gram to Kilogram",
			fields: fields{
				unit: measurements.Gram,
				value: 10000,
			},
			want: measurements.FromKilogram(10),
		},
		{
			name: "Pound to Kilogram",
			fields: fields{
				unit: measurements.Pound,
				value: 22.0462,
			},
			want: measurements.FromKilogram(10),
		},
		{
			name: "Ounce to Kilogram",
			fields: fields{
				unit: measurements.Ounce,
				value: 352.74,
			},
			want: measurements.FromKilogram(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := measurements.NewMass(tt.fields.unit, tt.fields.value)
			if got := s.ToKilogram(); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("ToKilogram() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_mass_ToGram(t *testing.T) {
	type fields struct {
		unit  measurements.MassUnit
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   measurements.Mass
	}{
		{
			name: "Kilogram to Gram",
			fields: fields{
				unit: measurements.Kilogram,
				value: 0.01,
			},
			want: measurements.FromGram(10),
		},
		{
			name: "Gram to Gram",
			fields: fields{
				unit: measurements.Gram,
				value: 10,
			},
			want: measurements.FromGram(10),
		},
		{
			name: "Pound to Gram",
			fields: fields{
				unit: measurements.Pound,
				value: 0.0220462,
			},
			want: measurements.FromGram(10),
		},
		{
			name: "Ounce to Gram",
			fields: fields{
				unit: measurements.Ounce,
				value: 0.35274,
			},
			want: measurements.FromGram(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := measurements.NewMass(tt.fields.unit, tt.fields.value)
			if got := s.ToGram(); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("ToGram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mass_ToPound(t *testing.T) {
	type fields struct {
		unit  measurements.MassUnit
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   measurements.Mass
	}{
		{
			name: "Kilogram to Pound",
			fields: fields{
				unit: measurements.Kilogram,
				value: 4.53592,
			},
			want: measurements.FromPound(10),
		},
		{
			name: "Gram to Pound",
			fields: fields{
				unit: measurements.Gram,
				value: 4535.92,
			},
			want: measurements.FromPound(10),
		},
		{
			name: "Pound to Pound",
			fields: fields{
				unit: measurements.Pound,
				value: 10,
			},
			want: measurements.FromPound(10),
		},
		{
			name: "Ounce to Pound",
			fields: fields{
				unit: measurements.Ounce,
				value: 160,
			},
			want: measurements.FromPound(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := measurements.NewMass(tt.fields.unit, tt.fields.value)
			if got := s.ToPound(); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("ToPound() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_mass_ToOunce(t *testing.T) {
	type fields struct {
		unit  measurements.MassUnit
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   measurements.Mass
	}{
		{
			name: "Kilogram to Ounce",
			fields: fields{
				unit: measurements.Kilogram,
				value: 0.283495,
			},
			want: measurements.FromOunce(10),
		},
		{
			name: "Gram to Ounce",
			fields: fields{
				unit: measurements.Gram,
				value: 283.495,
			},
			want: measurements.FromOunce(10),
		},
		{
			name: "Pound to Ounce",
			fields: fields{
				unit: measurements.Pound,
				value: 0.625,
			},
			want: measurements.FromOunce(10),
		},
		{
			name: "Ounce to Ounce",
			fields: fields{
				unit: measurements.Ounce,
				value: 10,
			},
			want: measurements.FromOunce(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := measurements.NewMass(tt.fields.unit, tt.fields.value)
			if got := s.ToOunce(); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("ToOunce() = %v, want %v", got, tt.want)
			}
		})
	}
}