package decimal

import "math/big"

// IsZero returns true if d == 0
func (d BigDecimal) IsZero() bool {
	d.optimize()
	if d.value.Cmp(zeroInt) != 0 {
		return false
	}
	if d.numerator != 0 {
		return false
	}
	return true
}

// IsPositive returns true if d > 0
func (d BigDecimal) IsPositive() bool {
	d.optimize()
	switch d.value.Cmp(zeroInt) {
	case 1:
		return true
	case -1:
		return false
	}
	if d.numerator != 0 && d.denominator != 0 {
		return true
	}
	return false
}

// IsNegative returns true if d > 0
func (d BigDecimal) IsNegative() bool {
	d.optimize()
	switch d.value.Cmp(zeroInt) {
	case 1:
		return false
	case -1:
		return true
	default:
		return false
	}
}

// Neg returns -d
// Simple cases
// Neg(-2 + 1/3) = Neg(-5/3) = 5/3 = 1 + 2/3
// Neg(2 + 1/3) = Neg(7/3) = -7/3 = -3 + 2/3
func (d BigDecimal) Neg() BigDecimal {
	if d.denominator == 0 {
		return BigDecimal{
			value: new(big.Int).Neg(d.value),
			scale: d.scale,
		}
	}

	d.optimize()
	// Break it down.
	//		D = (v1 + n1/d1)			* 10^(-s1)
	// <=> -D = (-v1 - n1/d1)			* 10^(-s1)
	// <=> -D = (-v1 - 1 + 1 - n1/d1)	* 10^(-s1)
	// <=> -D = (-(v1+1) + (d1-n1)/d1)	* 10^(-s1)
	vA1 := new(big.Int).Add(d.value, oneInt)
	return BigDecimal{
		value:       new(big.Int).Neg(vA1),
		scale:       d.scale,
		numerator:   d.denominator - d.numerator,
		denominator: d.denominator,
	}
}

// Abs returns |d|
// d.Abs() == d if d >= 0
// d.Abs() == -d if d < 0
func (d BigDecimal) Abs() BigDecimal {
	if d.IsNegative() {
		return d.Neg()
	}
	return d
}

// Sign returns the sign of decimal
// -1 for negative
// 0 for zero
// 1 for positive
func (d BigDecimal) Sign() int {
	if d.IsPositive() {
		return 1
	}
	if d.IsNegative() {
		return -1
	}
	return 0
}
