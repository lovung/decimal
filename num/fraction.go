package num

import (
	"math/big"
)

// Fraction represents ratio.
type Fraction struct {
	Numerator   big.Int
	Denominator big.Int
}

// TODO: Care if the denominator is 0
func NewFraction(
	numerator big.Int,
	denominator big.Int,
) Fraction {
	return Fraction{
		numerator,
		denominator,
	}
}

func NewFractionFromDecimal(d Decimal) Fraction {
	numer := *d.value
	exp := new(big.Int).SetInt64(int64(-d.exp))
	var denom big.Int
	denom.Exp(tenInt, exp, nil)

	return Fraction{
		numer,
		denom,
	}.Reduce()
}

func NewFractionFromStr(s string) (Fraction, error) {
	d, err := NewDecimalFromString(s)
	if err != nil {
		return Fraction{}, err
	}
	return NewFractionFromDecimal(d), nil
}

func MustNewFractionFromStr(s string) Fraction {
	d, err := NewDecimalFromString(s)
	if err != nil {
		panic(err)
	}

	return NewFractionFromDecimal(d)
}

func (f Fraction) Decimal() Decimal {
	n := NewDecimalFromBigInt(&f.Numerator, 0)
	d := NewDecimalFromBigInt(&f.Denominator, 0)
	return n.Div(d)
}

// func (f Fraction) String() string {
// 	return fmt.Sprintf("%d/%d", f.Numerator, f.Denominator)
// }

// Reduce makes given fraction reduced.
func (f Fraction) Reduce() Fraction {
	gcd := new(big.Int).GCD(nil, nil, &f.Numerator, &f.Denominator)
	f.Numerator.Div(&f.Numerator, gcd)
	f.Denominator.Div(&f.Denominator, gcd)
	return f
}

// MulNum multiplies fraction by given number.
func (f Fraction) MulNum(m big.Int) Fraction {
	f.Numerator.Mul(&f.Numerator, &m)

	return f
}

// Mul multiplies fraction by given fraction.
func (f Fraction) Mul(m Fraction) Fraction {
	f.Numerator.Mul(&f.Numerator, &m.Numerator)
	f.Denominator.Mul(&f.Denominator, &m.Denominator)

	return f
}

// Div devices fraction by given fraction.
func (f Fraction) Div(m Fraction) Fraction {
	f.Numerator.Mul(&f.Numerator, &m.Denominator)
	f.Denominator.Mul(&f.Denominator, &m.Numerator)

	return f
}

func (f Fraction) Inverse() Fraction {
	return Fraction{
		f.Denominator,
		f.Numerator,
	}
}
