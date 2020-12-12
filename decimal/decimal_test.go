package decimal

import (
	"math/big"
	"reflect"
	"testing"
)

func TestBigDecimal_rescale(t *testing.T) {
	type fields struct {
		value       *big.Int
		scale       int32
		numerator   int64
		denominator int64
		strCache    string
	}
	tests := []struct {
		name   string
		fields fields
		scale  int32
		want   BigDecimal
	}{
		{
			fields: fields{
				new(big.Int).SetInt64(12),
				0, 1, 3, "",
			},
			scale: -1,
			want: BigDecimal{
				new(big.Int).SetInt64(1),
				-1, 7, 30, "",
			},
		},
		{
			fields: fields{
				new(big.Int).SetInt64(123),
				0, 1, 3, "",
			},
			scale: -2,
			want: BigDecimal{
				new(big.Int).SetInt64(1),
				-2, 70, 300, "",
			},
		},
		{
			fields: fields{
				new(big.Int).SetInt64(123),
				0, 1, 3, "",
			},
			scale: 2,
			want: BigDecimal{
				new(big.Int).SetInt64(12300),
				2, 100, 3, "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bd := BigDecimal{
				value:       tt.fields.value,
				scale:       tt.fields.scale,
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
				strCache:    tt.fields.strCache,
			}
			if got := bd.rescale(tt.scale); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BigDecimal.rescale() = %v, want %v", got, tt.want)
			}
		})
	}
}