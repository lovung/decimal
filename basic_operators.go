package decimal

import "math/big"

// Add returns d + d2
func (d BigDecimal) Add(d2 BigDecimal) BigDecimal {
	rd, rd2 := rescalePair(d, d2)

	d3Value := new(big.Int).Add(rd.value, rd2.value)
	nn, nd := sumFraction(rd.numerator, rd.denominator, rd2.numerator, rd2.denominator)

	newBD := BigDecimal{
		value:       d3Value,
		scale:       rd.scale,
		numerator:   nn,
		denominator: nd,
	}
	newBD.optimize()
	return newBD
}
