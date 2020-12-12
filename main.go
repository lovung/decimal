package main

import (
	"fmt"

	"github.com/lovung/gomath/num"
)

func main() {
	fee, _ := num.NewDecimalFromString("0.0295")
	feeMin, _ := num.NewDecimalFromString("4")
	feeAmt, _ := num.NewDecimalFromString("2.95")

	fmt.Println((4 / 2.95) * 0.0295)
	fmt.Println(feeMin.Div(feeAmt).Mul(fee).String())
	fmt.Println(feeMin.Mul(fee.Div(feeAmt)).String())

	feeFrac := num.NewFractionFromDecimal(fee)
	feeMinFrac := num.NewFractionFromDecimal(feeMin)
	feeAmtFrac := num.NewFractionFromDecimal(feeAmt)
	fmt.Println(feeMinFrac.Div(feeAmtFrac).Mul(feeFrac).Decimal())
}
