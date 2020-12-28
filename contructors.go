package decimal

import "math/big"

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
