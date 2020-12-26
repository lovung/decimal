package decimal

import (
	"math/big"
	"sort"
	"testing"
)

type BigDecimalSlice []BigDecimal

func (p BigDecimalSlice) Len() int           { return len(p) }
func (p BigDecimalSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p BigDecimalSlice) Less(i, j int) bool { return p[i].Cmp(p[j]) < 0 }

func TestBigDecimal_Cmp(t *testing.T) {
	tests := []struct {
		name   string
		fields BigDecimal
		args   BigDecimal
		want   int
	}{
		{
			fields: Zero,
			args:   One,
			want:   -1,
		},
		{
			fields: One,
			args:   Zero,
			want:   1,
		},
		{
			fields: Ten,
			args:   Ten,
			want:   0,
		},
		{
			fields: BigDecimal{
				value: new(big.Int).Set(oneInt),
				scale: -1,
			},
			args: Ten,
			want: 0,
		},
		{
			fields: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       -1,
				numerator:   1,
				denominator: 3,
			},
			args: Ten,
			want: 1,
		},
		{
			fields: Ten,
			args: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       -1,
				numerator:   1,
				denominator: 3,
			},
			want: -1,
		},
		{
			fields: BigDecimal{
				value:       new(big.Int).Set(twelveInt),
				scale:       0,
				numerator:   1,
				denominator: 3,
			},
			args: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       -1,
				numerator:   7,
				denominator: 30,
			},
			want: 0,
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
				numerator:   0,
				denominator: 0,
			},
			want: 1,
		},
		{
			fields: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       0,
				numerator:   0,
				denominator: 0,
			},
			args: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       0,
				numerator:   1,
				denominator: 3,
			},
			want: -1,
		},
		{
			fields: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       0,
				numerator:   0,
				denominator: 0,
			},
			args: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       0,
				numerator:   0,
				denominator: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bd := tt.fields
			if got := bd.Cmp(tt.args); got != tt.want {
				t.Errorf("BigDecimal.Cmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_Cmp(b *testing.B) {
	decimals := BigDecimalSlice([]BigDecimal{})
	for i := 0; i < 1000000; i++ {
		decimals = append(decimals, NewBigDecimal(int64(i), 0))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Sort(decimals)
	}
}

func TestBigDecimal_Equal(t *testing.T) {
	type fields struct {
		value       *big.Int
		scale       int32
		numerator   uint64
		denominator uint64
		strCache    string
	}
	type args struct {
		ref BigDecimal
	}
	tests := []struct {
		name   string
		fields BigDecimal
		args   BigDecimal
		want   bool
	}{
		{
			fields: Zero,
			args:   One,
			want:   false,
		},
		{
			fields: One,
			args:   Zero,
			want:   false,
		},
		{
			fields: Ten,
			args:   Ten,
			want:   true,
		},
		{
			fields: BigDecimal{
				value: new(big.Int).Set(oneInt),
				scale: -1,
			},
			args: Ten,
			want: true,
		},
		{
			fields: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       -1,
				numerator:   1,
				denominator: 3,
			},
			args: Ten,
			want: false,
		},
		{
			fields: Ten,
			args: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       -1,
				numerator:   1,
				denominator: 3,
			},
			want: false,
		},
		{
			fields: BigDecimal{
				value:       new(big.Int).Set(twelveInt),
				scale:       0,
				numerator:   1,
				denominator: 3,
			},
			args: BigDecimal{
				value:       new(big.Int).Set(oneInt),
				scale:       -1,
				numerator:   7,
				denominator: 30,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields
			if got := d.Equal(tt.args); got != tt.want {
				t.Errorf("BigDecimal.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
