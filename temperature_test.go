package measurements_test

import (
	"reflect"
	"testing"

	"github.com/RossMerr/herms/src/controller/measurements"
)

func Test_temperature_ToCelsius(t *testing.T) {
	type fields struct {
		unit  measurements.TemperatureUnit
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   measurements.Temperature
	}{
		{
			name: "Celsius to Celsius",
			fields: fields{
				unit:  measurements.Celsius,
				value: 10,
			},
			want: measurements.FromCelsius(10),
		},
		{
			name: "Fahrenheit to Celsius",
			fields: fields{
				unit:  measurements.Fahrenheit,
				value: 50,
			},
			want: measurements.FromCelsius(10),
		},
		{
			name: "Kelvin to Celsius",
			fields: fields{
				unit:  measurements.Kelvin,
				value: 283.15,
			},
			want: measurements.FromCelsius(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := measurements.NewTemperature(tt.fields.unit, tt.fields.value)
			if got := s.To(measurements.Celsius); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("ToCelsius() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_temperature_ToFahrenheit(t *testing.T) {
	type fields struct {
		unit  measurements.TemperatureUnit
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   measurements.Temperature
	}{
		{
			name: "Celsius to Fahrenheit",
			fields: fields{
				unit:  measurements.Celsius,
				value: -12.22,
			},
			want: measurements.FromFahrenheit(10),
		},
		{
			name: "Fahrenheit to Fahrenheit",
			fields: fields{
				unit:  measurements.Fahrenheit,
				value: 10,
			},
			want: measurements.FromFahrenheit(10),
		},
		{
			name: "Kelvin to Fahrenheit",
			fields: fields{
				unit:  measurements.Kelvin,
				value: 260.928,
			},
			want: measurements.FromFahrenheit(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := measurements.NewTemperature(tt.fields.unit, tt.fields.value)
			if got := s.To(measurements.Fahrenheit); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("ToFahrenheit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_temperature_ToKelvin(t *testing.T) {
	type fields struct {
		unit  measurements.TemperatureUnit
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   measurements.Temperature
	}{
		{
			name: "Celsius to Kelvin",
			fields: fields{
				unit:  measurements.Celsius,
				value: -263.15,
			},
			want: measurements.FromKelvin(10),
		},
		{
			name: "Fahrenheit to Kelvin",
			fields: fields{
				unit:  measurements.Fahrenheit,
				value: -441.67,
			},
			want: measurements.FromKelvin(10),
		},
		{
			name: "Kelvin to Kelvin",
			fields: fields{
				unit:  measurements.Kelvin,
				value: 10,
			},
			want: measurements.FromKelvin(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := measurements.NewTemperature(tt.fields.unit, tt.fields.value)
			if got := s.To(measurements.Kelvin); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("ToKelvin() = %v, want %v", got, tt.want)
			}
		})
	}
}
