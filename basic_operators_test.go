package decimal

import (
	"math/big"
	"reflect"
	"testing"
)

func TestBigDecimal_Add(t *testing.T) {
	tests := []struct {
		name   string
		fields BigDecimal
		args   BigDecimal
		want   BigDecimal
	}{
		{
			fields: One,
			args:   One,
			want:   Two,
		},
		{
			fields: Zero,
			args:   One,
			want:   One,
		},
		{
			fields: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       -1,
				numerator:   1,
				denominator: 3,
			},
			args: One,
			want: BigDecimal{
				value:       new(big.Int).SetInt64(14),
				scale:       0,
				numerator:   1,
				denominator: 3,
			},
		},
		{
			fields: One,
			args: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       -1,
				numerator:   1,
				denominator: 3,
			},
			want: BigDecimal{
				value:       new(big.Int).SetInt64(14),
				scale:       0,
				numerator:   1,
				denominator: 3,
			},
		},

		{
			fields: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       -1,
				numerator:   1,
				denominator: 3,
			},
			args: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       -1,
				numerator:   2,
				denominator: 3,
			},
			want: BigDecimal{
				value:       new(big.Int).SetInt64(3),
				scale:       -1,
				numerator:   0,
				denominator: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.Add(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BigDecimal.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBigDecimal_Sub(t *testing.T) {
	tests := []struct {
		name   string
		fields BigDecimal
		args   BigDecimal
		want   BigDecimal
	}{
		{
			fields: One,
			args:   One,
			want:   Zero,
		},
		{
			fields: One,
			args:   Zero,
			want:   One,
		},
		{
			fields: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       -1,
				numerator:   1,
				denominator: 3,
			},
			args: One,
			want: BigDecimal{
				value:       new(big.Int).SetInt64(12),
				scale:       0,
				numerator:   1,
				denominator: 3,
			},
		},
		{
			fields: One,
			args: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       -1,
				numerator:   1,
				denominator: 3,
			},
			want: BigDecimal{
				value:       new(big.Int).SetInt64(-13),
				scale:       0,
				numerator:   2,
				denominator: 3,
			},
		},
		{
			fields: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       -1,
				numerator:   1,
				denominator: 3,
			},
			args: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       -1,
				numerator:   2,
				denominator: 3,
			},
			want: BigDecimal{
				value:       new(big.Int).SetInt64(-1),
				scale:       -1,
				numerator:   2,
				denominator: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			got := d.Sub(tt.args)
			if got.Cmp(tt.want) != 0 {
				t.Errorf("BigDecimal.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBigDecimal_Mul(t *testing.T) {
	tests := []struct {
		name   string
		fields BigDecimal
		args   BigDecimal
		want   BigDecimal
	}{
		{
			fields: One,
			args:   One,
			want:   One,
		},
		{
			fields: One,
			args:   Zero,
			want:   Zero,
		},
		{
			fields: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       0,
				numerator:   1,
				denominator: 3,
			},
			args: One,
			want: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       0,
				numerator:   1,
				denominator: 3,
			},
		},
		{
			fields: One,
			args: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       0,
				numerator:   1,
				denominator: 3,
			},
			want: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       0,
				numerator:   1,
				denominator: 3,
			},
		},
		{
			fields: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       0,
				numerator:   1,
				denominator: 3,
			},
			args: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       0,
				numerator:   1,
				denominator: 3,
			},
			want: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       0,
				numerator:   7,
				denominator: 9,
			},
		},
		{
			fields: BigDecimal{
				value:       new(big.Int).SetInt64(-2),
				scale:       0,
				numerator:   1,
				denominator: 3,
			},
			args: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       0,
				numerator:   1,
				denominator: 3,
			},
			want: BigDecimal{
				value:       new(big.Int).SetInt64(-3),
				scale:       0,
				numerator:   7,
				denominator: 9,
			},
		},
		{
			fields: BigDecimal{
				value:       new(big.Int).SetInt64(-2),
				scale:       0,
				numerator:   1,
				denominator: 3,
			},
			args: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       -1,
				numerator:   1,
				denominator: 3,
			},
			want: BigDecimal{
				value:       new(big.Int).SetInt64(-3),
				scale:       -1,
				numerator:   7,
				denominator: 9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.Mul(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BigDecimal.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}
