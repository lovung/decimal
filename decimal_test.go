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
		numerator   uint64
		denominator uint64
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
				new(big.Int).SetInt64(123),
				-1, 1, 3, "",
			},
			scale: -1,
			want: BigDecimal{
				new(big.Int).SetInt64(123),
				-1, 1, 3, "",
			},
		},
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
				new(big.Int).SetInt64(12333),
				2, 1, 3, "",
			},
		},
		{
			fields: fields{
				new(big.Int).SetInt64(-12),
				0, 1, 3, "",
			},
			scale: -1,
			want: BigDecimal{
				new(big.Int).SetInt64(-2),
				-1, 25, 30, "",
			},
		},
		{
			fields: fields{
				new(big.Int).SetInt64(-12),
				0, 5, 3, "",
			},
			scale: 0,
			want: BigDecimal{
				new(big.Int).SetInt64(-11),
				0, 2, 3, "",
			},
		},
		{
			fields: fields{
				new(big.Int).SetInt64(12),
				0, 5, 3, "",
			},
			scale: 0,
			want: BigDecimal{
				new(big.Int).SetInt64(13),
				0, 2, 3, "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := BigDecimal{
				value:       tt.fields.value,
				scale:       tt.fields.scale,
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
				strCache:    tt.fields.strCache,
			}
			if got := d.rescale(tt.scale); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BigDecimal.rescale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBigDecimal_reduce(t *testing.T) {
	type fields struct {
		value       *big.Int
		scale       int32
		numerator   uint64
		denominator uint64
		strCache    string
	}
	tests := []struct {
		name   string
		fields fields
		want   *BigDecimal
	}{
		{
			fields: fields{
				new(big.Int).SetInt64(1),
				0, 0, 0, "",
			},
			want: &BigDecimal{
				new(big.Int).SetInt64(1),
				0, 0, 0, "",
			},
		},
		{
			fields: fields{
				new(big.Int).SetInt64(1),
				0, 0, 3, "",
			},
			want: &BigDecimal{
				new(big.Int).SetInt64(1),
				0, 0, 0, "",
			},
		},
		{
			fields: fields{
				new(big.Int).SetInt64(1),
				0, 9, 27, "",
			},
			want: &BigDecimal{
				new(big.Int).SetInt64(1),
				0, 1, 3, "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &BigDecimal{
				value:       tt.fields.value,
				scale:       tt.fields.scale,
				numerator:   tt.fields.numerator,
				denominator: tt.fields.denominator,
				strCache:    tt.fields.strCache,
			}
			d.reduce()
			if !reflect.DeepEqual(d, tt.want) {
				t.Errorf("BigDecimal.rescale() = %v, want %v", d, tt.want)
			}
		})
	}
	shouldPanic(t, func() {
		d := BigDecimal{
			new(big.Int).SetInt64(1),
			0, 1, 0, "",
		}
		d.reduce()
	})
}

func shouldPanic(t *testing.T, f func()) {
	defer func() { recover() }()
	f()
	t.Errorf("should have panicked")
}
