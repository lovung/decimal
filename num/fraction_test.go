package num

import (
	"math/big"
	"reflect"
	"testing"
)

func TestNewFractionFromStr(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name    string
		args    args
		want    Fraction
		wantErr bool
	}{
		{
			args:    args{"0.25"},
			want:    Fraction{*oneInt, *fourInt},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFractionFromStr(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFractionFromStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFractionFromStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFraction_Reduce(t *testing.T) {
	one := new(big.Int).SetInt64(1)
	eight := new(big.Int).SetInt64(8)
	tests := []struct {
		name  string
		input Fraction
		want  Fraction
	}{
		{
			input: MustNewFractionFromStr("0.125"),
			want: Fraction{
				Numerator:   *one,
				Denominator: *eight,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Fraction{
				Numerator:   tt.input.Numerator,
				Denominator: tt.input.Denominator,
			}
			if got := f.Reduce(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fraction.Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_Reduce(b *testing.B) {
	// run the Fib function b.N times
	first := new(big.Int).SetInt64(1000)
	second := new(big.Int).SetInt64(285)
	for n := 0; n < b.N; n++ {
		NewFraction(*first, *second).Reduce()
	}
}
