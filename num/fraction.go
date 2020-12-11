package num

import (
	"math/big"
)

// Fraction represents ratio.
type Fraction struct {
	Numerator   big.Int
	Denominator big.Int
}

func NewFraction(
	numerator big.Int,
	denominator big.Int,
) Fraction {
	return Fraction{
		numerator,
		denominator,
	}
}

func NewFractionFromStr(s string) (Fraction, error) {
	d, err := NewDecimalFromString(s)
	if err != nil {
		return Fraction{}, err
	}

	numer := *d.value
	exp := new(big.Int).SetInt64(int64(-d.exp))
	var denom big.Int
	denom.Exp(tenInt, exp, nil)

	return Fraction{
		numer,
		denom,
	}.Reduce(), nil
}

func MustNewFractionFromStr(s string) Fraction {
	d, err := NewDecimalFromString(s)
	if err != nil {
		panic(err)
	}

	numer := *d.value
	exp := new(big.Int).SetInt64(int64(-d.exp))
	var denom big.Int
	denom.Exp(tenInt, exp, nil)

	return Fraction{
		numer,
		denom,
	}.Reduce()
}

// func (f Fraction) String() string {
// 	return fmt.Sprintf("%d/%d", f.Numerator, f.Denominator)
// }

// Reduce makes given fraction reduced.
func (f Fraction) Reduce() Fraction {
	gcd := GCDRemainder(f.Numerator, f.Denominator)
	f.Numerator.Div(&f.Numerator, &gcd)
	f.Denominator.Div(&f.Denominator, &gcd)
	return f
}

// MultiplyByNumber multiplies fraction by given number.
func (f Fraction) MultiplyByNumber(m big.Int) Fraction {
	f.Numerator.Mul(&f.Numerator, &m)

	return f
}

// MultiplyByFraction multiplies fraction by given fraction.
func (f Fraction) MultiplyByFraction(m Fraction) Fraction {
	f.Numerator.Mul(&f.Numerator, &m.Numerator)
	f.Denominator.Mul(&f.Denominator, &m.Denominator)

	return f
}
