package num

import (
	"testing"
)

func Benchmark_2(b *testing.B) {
	fee, _ := NewDecimalFromString("0.0295")
	feeMin, _ := NewDecimalFromString("4")
	feeAmt, _ := NewDecimalFromString("2.95")
	for n := 0; n < b.N; n++ {
		feeMin.Div(feeAmt).Mul(fee)
	}
}

func Benchmark_3(b *testing.B) {
	fee, _ := NewDecimalFromString("0.0295")
	feeMin, _ := NewDecimalFromString("4")
	feeAmt, _ := NewDecimalFromString("2.95")
	for n := 0; n < b.N; n++ {
		feeMin.Mul(fee.Div(feeAmt))
	}
}

func Benchmark_4(b *testing.B) {
	fee, _ := NewDecimalFromString("0.0295")
	feeMin, _ := NewDecimalFromString("4")
	feeAmt, _ := NewDecimalFromString("2.95")
	feeFrac := NewFractionFromDecimal(fee)
	feeMinFrac := NewFractionFromDecimal(feeMin)
	feeAmtFrac := NewFractionFromDecimal(feeAmt)
	for n := 0; n < b.N; n++ {
		feeMinFrac.Div(feeAmtFrac).Mul(feeFrac).Decimal()
	}
}
