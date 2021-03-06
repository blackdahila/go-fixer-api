package currency

import (
	"reflect"
	"testing"
)

func TestRound(t *testing.T) {
	type args struct {
		some float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Rounds up", args{3.66666}, 3.67},
		{"Rounds down", args{4.22222}, 4.22},
		{"Does not truncate when not needed", args{1.0}, 1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Round(tt.args.some); got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiply(t *testing.T) {
	rates := &Rates{
		"USD": 3.51,
		"EUR": 4.15,
		"PLN": 1.0,
	}

	multipliedRates := Rates{
		"USD": 7.02,
		"EUR": 8.3,
		"PLN": 2.0,
	}

	type args struct {
		rates  *Rates
		amount float64
	}
	tests := []struct {
		name string
		args args
		want Rates
	}{
		{"Should return proper values in map after multiplication.",
			args{rates, 2},
			multipliedRates,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.rates.Multiply(tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
