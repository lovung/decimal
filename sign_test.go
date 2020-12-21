package decimal

import (
	"math/big"
	"reflect"
	"testing"
)

func TestBigDecimal_IsZero(t *testing.T) {
	tests := []struct {
		name   string
		fields BigDecimal
		want   bool
	}{
		{"#1", One, false},
		{"#2", Zero, true},
		{"#3", BigDecimal{
			value: new(big.Int).SetInt64(-1),
		}, false},
		{"#4", BigDecimal{
			value:       new(big.Int).SetInt64(0),
			scale:       -1,
			numerator:   1,
			denominator: 3,
		}, false},
		{"#5", BigDecimal{
			value:       new(big.Int).SetInt64(-1),
			scale:       -1,
			numerator:   1,
			denominator: 3,
		}, false},
		{"#6", BigDecimal{
			value:       new(big.Int).SetInt64(-1),
			scale:       -1,
			numerator:   3,
			denominator: 3,
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.IsZero(); got != tt.want {
				t.Errorf("BigDecimal.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBigDecimal_IsPositive(t *testing.T) {
	tests := []struct {
		name   string
		fields BigDecimal
		want   bool
	}{
		{"#1", One, true},
		{"#2", Zero, false},
		{"#3", BigDecimal{
			value: new(big.Int).SetInt64(-1),
		}, false},
		{"#4", BigDecimal{
			value:       new(big.Int).SetInt64(0),
			scale:       -1,
			numerator:   1,
			denominator: 3,
		}, true},
		{"#5", BigDecimal{
			value:       new(big.Int).SetInt64(-1),
			scale:       -1,
			numerator:   1,
			denominator: 3,
		}, false},
		{"#6", BigDecimal{
			value:       new(big.Int).SetInt64(-1),
			scale:       -1,
			numerator:   3,
			denominator: 3,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.IsPositive(); got != tt.want {
				t.Errorf("BigDecimal.IsPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBigDecimal_IsNegative(t *testing.T) {
	tests := []struct {
		name   string
		fields BigDecimal
		want   bool
	}{
		{"#1", One, false},
		{"#2", Zero, false},
		{"#3", BigDecimal{
			value: new(big.Int).SetInt64(-1),
		}, true},
		{"#4", BigDecimal{
			value:       new(big.Int).SetInt64(0),
			scale:       -1,
			numerator:   1,
			denominator: 3,
		}, false},
		{"#5", BigDecimal{
			value:       new(big.Int).SetInt64(-1),
			scale:       -1,
			numerator:   1,
			denominator: 3,
		}, true},
		{"#6", BigDecimal{
			value:       new(big.Int).SetInt64(-1),
			scale:       -1,
			numerator:   3,
			denominator: 3,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.IsNegative(); got != tt.want {
				t.Errorf("BigDecimal.IsNegative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBigDecimal_Neg(t *testing.T) {
	tests := []struct {
		name   string
		fields BigDecimal
		want   BigDecimal
	}{
		{"#1", One, BigDecimal{value: new(big.Int).SetInt64(-1)}},
		{"#2", Zero, Zero},
		{"#3", BigDecimal{value: new(big.Int).SetInt64(-1)}, One},
		{"#4", BigDecimal{
			value:       new(big.Int).SetInt64(1),
			numerator:   1,
			denominator: 3,
		}, BigDecimal{
			value:       new(big.Int).SetInt64(-2),
			numerator:   2,
			denominator: 3,
		}},
		{"#5", BigDecimal{
			value:       new(big.Int).SetInt64(-1),
			numerator:   1,
			denominator: 3,
		}, BigDecimal{
			value:       new(big.Int).SetInt64(0),
			numerator:   2,
			denominator: 3,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.Neg(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BigDecimal.Neg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBigDecimal_Abs(t *testing.T) {
	tests := []struct {
		name   string
		fields BigDecimal
		want   BigDecimal
	}{
		{"#1", One, One},
		{"#2", Zero, Zero},
		{"#3", BigDecimal{value: new(big.Int).SetInt64(-1)}, One},
		{"#4", BigDecimal{
			value:       new(big.Int).SetInt64(1),
			numerator:   1,
			denominator: 3,
		}, BigDecimal{
			value:       new(big.Int).SetInt64(1),
			numerator:   1,
			denominator: 3,
		}},
		{"#5", BigDecimal{
			value:       new(big.Int).SetInt64(-1),
			numerator:   1,
			denominator: 3,
		}, BigDecimal{
			value:       new(big.Int).SetInt64(0),
			numerator:   2,
			denominator: 3,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.Abs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BigDecimal.Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
