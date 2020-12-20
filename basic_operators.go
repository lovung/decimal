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

// Mul returns d * d2
func (d BigDecimal) Mul(d2 BigDecimal) BigDecimal {
	var (
		nn int64
		nd uint64
	)

	d3Value := new(big.Int).Mul(d.value, d2.value)
	if d.denominator == 0 && d2.denominator != 0 {
		nn = d.value.Int64() * int64(d2.numerator)
		nd = d2.denominator
	}
	if d.denominator != 0 && d2.denominator == 0 {
		nn = d2.value.Int64() * int64(d.numerator)
		nd = d.denominator
	}
	if d.denominator != 0 && d2.denominator != 0 {
		nn = d.value.Int64() * int64(d.denominator) * int64(d2.numerator)
		nn += d2.value.Int64() * int64(d2.denominator) * int64(d.numerator)
		nn += int64(d.numerator) * int64(d2.numerator)
		nd = d.denominator * d.denominator
	}
	if nn < 0 {
		d3Value.Sub(d3Value, new(big.Int).SetInt64(nn/int64(nd)))
		nn %= int64(nd)
		if nn < 0 {
			d3Value.Sub(d3Value, oneInt)
			nn += int64(nd)
		}
	}

	newBD := BigDecimal{
		value:       d3Value,
		scale:       d.scale + d2.scale,
		numerator:   uint64(nn),
		denominator: nd,
	}
	newBD.optimize()
	return newBD
}
