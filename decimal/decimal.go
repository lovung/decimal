package decimal

import "math/big"

// BigDecimal real value is
//		d = (value + numerator / denominator) * 10 ^ (-scale)
// We have:
//		denominator = 0 as initial => numerator / 0 = 0
//	it means numerator / denominator only valid if denominator != 0
// Simple type of BigDecimal:
//		- The numerator shouble be less than denominator (or both are zeros)
type BigDecimal struct {
	value *big.Int
	scale int32

	numerator   int64
	denominator int64

	strCache string
}

// NewBigDecimal returns a new fixed-point big decimal.
// Decimal = value * 10 ^ (-scale)
//
// Example: 0.19  ---> value = 19, scale = 2
func NewBigDecimal(value int64, scale int32) BigDecimal {
	return BigDecimal{
		value: big.NewInt(value),
		scale: scale,
	}
}

func NewBigDecimalFromInt64(value int64) BigDecimal {
	return BigDecimal{
		value: big.NewInt(value),
	}
}

func NewBigDecimalFromInt(value int) BigDecimal {
	return BigDecimal{
		value: big.NewInt(int64(value)),
	}
}

func NewBigDecimalFromBigInt(value *big.Int, scale int32) BigDecimal {
	return BigDecimal{
		value: new(big.Int).Set(value),
		scale: scale,
	}
}

// NewBigDecimalFromString returns a new BigDecimal from a string.
// Alway want to have the zero-nearest positive-scale
// Support:
//		- Sign: NewBigDecimalFromString("-1")
//				---> value: -1, scale: 0
//		- Dot: NewBigDecimalFromString(".123")
//				---> value: 123, scale: 3
//		- E/e: NewBigDecimalFromString("1e9") / NewBigDecimalFromString(1E9)
//				---> value: 1, scale: -9
//		- Trim trailing zero: NewBigDecimalFromString("1.2300000")
//				---> value: 123, scale: 2
//		- Repeating/recurring decimal:
//			NewBigDecimalFromString("0.(3)")
//				---> value: 0, scale: 0, numerator: 1, denominator: 3
//			NewBigDecimalFromString("1.23(3)")
//				---> value: 12, scale: 1, numerator: 1, denominator: 3
//			NewBigDecimalFromString("1.(428571)")
//				---> value: 1, scale: 0, numerator: 3, denominator: 7
//				-x-> value: 0, scale: -1, numerator: 1, denominator: 7
func NewBigDecimalFromString(value string) (BigDecimal, error) {
	// TODO: implement
	return BigDecimal{}, nil
}

func (bd *BigDecimal) ensureInitialized() {
	if bd.value == nil {
		bd.value = new(big.Int)
	}
}

func (bd BigDecimal) Cmp(ref BigDecimal) int {
	// TODO: implement
	return 0
}

func (bd BigDecimal) rescale(scale int32) BigDecimal {
	bd.ensureInitialized()

	if bd.scale == scale {
		return BigDecimal{
			value:       new(big.Int).Set(bd.value),
			scale:       bd.scale,
			numerator:   bd.numerator,
			denominator: bd.denominator,
		}
	}
	diffScale := scale - bd.scale
	value := new(big.Int).Set(bd.value)
	bigDec := BigDecimal{value: value, scale: scale}
	if diffScale < 0 {
		expScale := new(big.Int).Exp(tenInt, big.NewInt(int64(-diffScale)), nil)
		rem := new(big.Int)
		value, rem = value.QuoRem(value, expScale, rem)
		r := rem.Int64()
		bigDec.numerator, bigDec.denominator = sumFraction(
			r, expScale.Int64(),
			bd.numerator, bd.denominator*expScale.Int64(),
		)
	} else {
		expScale := new(big.Int).Exp(tenInt, big.NewInt(int64(diffScale)), nil)
		value = value.Mul(value, expScale)
		bigDec.numerator = bd.numerator * expScale.Int64()
		bigDec.denominator = bd.denominator
	}

	return bigDec
}
