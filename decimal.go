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

// New returns the new BigDecimal from another BigDecimal
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

// NewBigDecimalFromInt64 returns a new fixed-point big decimal with int64
func NewBigDecimalFromInt64(value int64) BigDecimal {
	return BigDecimal{
		value: big.NewInt(value),
	}
}

// NewBigDecimalFromInt returns a new fixed-point big decimal with int
func NewBigDecimalFromInt(value int) BigDecimal {
	return BigDecimal{
		value: big.NewInt(int64(value)),
	}
}

// NewBigDecimalFromBigInt returns a new fixed-point big decimal with big.Int and scale
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

func (d *BigDecimal) ensureInitialized() {
	if d.value == nil {
		d.value = new(big.Int)
	}
}

func (d BigDecimal) toFractionIgnoreScale() (*big.Int, *big.Int) {
	if d.denominator == 0 {
		return new(big.Int).Set(d.value), new(big.Int).Set(oneInt)
	}
	num, dem := new(big.Int).SetUint64(d.numerator), new(big.Int).SetUint64(d.denominator)
	vMulD := new(big.Int).Set(d.value)
	vMulD = vMulD.Mul(vMulD, dem)
	num.Add(num, vMulD)
	return num, dem
}

func (d *BigDecimal) optimize() {
	if d.denominator == 0 {
		if d.numerator != 0 {
			panic("denominator is zero but numerator not")
		}
		return
	}
	if d.numerator >= d.denominator {
		d.value = d.value.Add(d.value, new(big.Int).SetUint64(d.numerator/d.denominator))
		d.numerator %= d.denominator
	}
	if d.numerator == 0 {
		d.denominator = 0
	}
}

// rescale helps to change the scale value but keep the real decimal value.
// rescale supports some operators; basically, the sum/add methods need two numbers
// have the same scale
func (d BigDecimal) rescale(scale int32) BigDecimal {
	bigDec := BigDecimal{}
	d.ensureInitialized()
	bigDec.ensureInitialized()

	if d.scale == scale {
		bigDec = BigDecimal{
			value:       new(big.Int).Set(d.value),
			scale:       d.scale,
			numerator:   d.numerator,
			denominator: d.denominator,
		}
		bigDec.optimize()
		return bigDec
	}
	diffScale := scale - d.scale
	value := new(big.Int).Set(d.value)
	bigDec = BigDecimal{value: value, scale: scale}
	if diffScale < 0 {
		expScale := new(big.Int).Exp(tenInt, big.NewInt(int64(-diffScale)), nil)
		rem := new(big.Int)
		value, rem = value.DivMod(value, expScale, rem)
		r := rem.Uint64()
		bigDec.numerator, bigDec.denominator = addFraction(
			r, expScale.Uint64(),
			d.numerator, d.denominator*expScale.Uint64(),
		)
	} else {
		expScale := new(big.Int).Exp(tenInt, big.NewInt(int64(diffScale)), nil)
		value = value.Mul(value, expScale)
		bigDec.numerator = d.numerator * expScale.Uint64()
		bigDec.denominator = d.denominator
	}
	bigDec.optimize()
	return bigDec
}

// rescalePair rescales two decimals to common exponential value (minimal exp of both decimals)
func rescalePair(d1 BigDecimal, d2 BigDecimal) (BigDecimal, BigDecimal) {
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
