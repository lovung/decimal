package decimal

import "math/big"

// Add returns d + d2
func (d BigDecimal) Add(d2 BigDecimal) BigDecimal {
	rd, rd2 := rescalePair(d, d2)

	d3Value := new(big.Int).Add(rd.value, rd2.value)
	nn, nd := addFraction(rd.numerator, rd.denominator, rd2.numerator, rd2.denominator)

	newBD := BigDecimal{
		value:       d3Value,
		scale:       rd.scale,
		numerator:   nn,
		denominator: nd,
	}
	newBD.optimize()
	return newBD
}

// Sub returns d - d2
func (d BigDecimal) Sub(d2 BigDecimal) BigDecimal {
	rd, rd2 := rescalePair(d, d2)

	d3Value := new(big.Int).Sub(rd.value, rd2.value)
	nn, nd := subFraction(rd.numerator, rd.denominator, rd2.numerator, rd2.denominator)

	if nn < 0 {
		d3Value.Sub(d3Value, oneInt)
		nn += int64(nd)
	}

	newBD := BigDecimal{
		value:       d3Value,
		scale:       rd.scale,
		numerator:   uint64(nn),
		denominator: nd,
	}
	newBD.optimize()
	return newBD
}
