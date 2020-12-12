package num_test

import (
	"testing"

	"github.com/lovung/gomath/num"
)

func Benchmark_2(b *testing.B) {
	fee, _ := num.NewDecimalFromString("0.0295")
	feeMin, _ := num.NewDecimalFromString("4")
	feeAmt, _ := num.NewDecimalFromString("2.95")
	for n := 0; n < b.N; n++ {
		feeMin.Div(feeAmt).Mul(fee)
	}
}

func Benchmark_3(b *testing.B) {
	fee, _ := num.NewDecimalFromString("0.0295")
	feeMin, _ := num.NewDecimalFromString("4")
	feeAmt, _ := num.NewDecimalFromString("2.95")
	for n := 0; n < b.N; n++ {
		feeMin.Mul(fee.Div(feeAmt))
	}
}

func Benchmark_4(b *testing.B) {
	fee, _ := num.NewDecimalFromString("0.0295")
	feeMin, _ := num.NewDecimalFromString("4")
	feeAmt, _ := num.NewDecimalFromString("2.95")
	feeFrac := num.NewFractionFromDecimal(fee)
	feeMinFrac := num.NewFractionFromDecimal(feeMin)
	feeAmtFrac := num.NewFractionFromDecimal(feeAmt)
	for n := 0; n < b.N; n++ {
		feeMinFrac.Div(feeAmtFrac).Mul(feeFrac).Decimal()
	}
}
