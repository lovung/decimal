package decimal

import (
	"math/big"
)

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

	numerator   uint64
	denominator uint64

	strCache string
}

var (
	Zero = BigDecimal{
		big.NewInt(0),
		0, 0, 0, "0",
	}
	One = BigDecimal{
		big.NewInt(1),
		0, 0, 0, "1",
	}
	Ten = BigDecimal{
		big.NewInt(10),
		0, 0, 0, "1",
	}
)

func New(ref BigDecimal) BigDecimal {
	ref.ensureInitialized()
	return BigDecimal{
		value:       new(big.Int).Set(ref.value),
		scale:       ref.scale,
		numerator:   ref.numerator,
		denominator: ref.denominator,
	}
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

func (bd BigDecimal) toFractionIgnoreScale() (*big.Int, *big.Int) {
	if bd.denominator == 0 {
		return new(big.Int).Set(bd.value), new(big.Int).Set(oneInt)
	}
	num, dem := new(big.Int).SetUint64(bd.numerator), new(big.Int).SetUint64(bd.denominator)
	vMulD := new(big.Int).Set(bd.value)
	vMulD = vMulD.Mul(vMulD, dem)
	num.Add(num, vMulD)
	return num, dem
}

func (bd *BigDecimal) optimize() {
	if bd.numerator > bd.denominator {
		bd.value = bd.value.Add(bd.value, new(big.Int).SetUint64(bd.numerator/bd.denominator))
		bd.numerator %= bd.denominator
	}
}

// rescale helps to change the scale value but keep the real decimal value.
// rescale supports some operators; basically, the sum/add methods need two numbers
// have the same scale
func (bd BigDecimal) rescale(scale int32) BigDecimal {
	bigDec := BigDecimal{}
	bd.ensureInitialized()
	bigDec.ensureInitialized()

	if bd.scale == scale {
		bigDec = BigDecimal{
			value:       new(big.Int).Set(bd.value),
			scale:       bd.scale,
			numerator:   bd.numerator,
			denominator: bd.denominator,
		}
		bigDec.optimize()
		return bigDec
	}
	diffScale := scale - bd.scale
	value := new(big.Int).Set(bd.value)
	bigDec = BigDecimal{value: value, scale: scale}
	if diffScale < 0 {
		expScale := new(big.Int).Exp(tenInt, big.NewInt(int64(-diffScale)), nil)
		rem := new(big.Int)
		value, rem = value.DivMod(value, expScale, rem)
		r := rem.Uint64()
		bigDec.numerator, bigDec.denominator = sumFraction(
			r, expScale.Uint64(),
			bd.numerator, bd.denominator*expScale.Uint64(),
		)
	} else {
		expScale := new(big.Int).Exp(tenInt, big.NewInt(int64(diffScale)), nil)
		value = value.Mul(value, expScale)
		bigDec.numerator = bd.numerator * expScale.Uint64()
		bigDec.denominator = bd.denominator
	}
	bigDec.optimize()
	return bigDec
}

// RescalePair rescales two decimals to common exponential value (minimal exp of both decimals)
func RescalePair(d1 BigDecimal, d2 BigDecimal) (BigDecimal, BigDecimal) {
	d1.ensureInitialized()
	d2.ensureInitialized()

	if d1.scale == d2.scale {
		return d1, d2
	}

	baseScale := maxInt32(d1.scale, d2.scale)
	if baseScale != d1.scale {
		return d1.rescale(baseScale), d2
	}
	return d1, d2.rescale(baseScale)
}
