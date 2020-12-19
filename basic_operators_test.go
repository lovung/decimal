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
