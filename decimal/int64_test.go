package decimal

import (
	"testing"
)

func Benchmark_tenPow(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tenPow(30)
		tenPow(20)
		tenPow(0)
		tenPow(1000)
		tenPow(100)
	}
}

func Test_tenPow(t *testing.T) {
	type args struct {
		exponent uint64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{args: args{0}, want: 1},
		{args: args{1}, want: 10},
		{args: args{2}, want: 100},
		{args: args{3}, want: 1000},
		{args: args{4}, want: 10000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tenPow(tt.args.exponent); got != tt.want {
				t.Errorf("tenPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
