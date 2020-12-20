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

// Div returns d / d2
// The return value will be recurring decimal if can
func (d BigDecimal) Div(d2 BigDecimal) BigDecimal {
	if d2.Cmp(Zero) == 0 {
		panic("divise by zero")
	}

	var (
		dn  = new(big.Int).Set(d.value)
		dd  = new(big.Int).Set(d2.value)
		dd1 = new(big.Int).SetUint64(d.denominator)
		dd2 = new(big.Int).SetUint64(d2.denominator)
	)

	if d.denominator != 0 {
		dn = dn.Mul(dn, dd1)
		dn = dn.Add(dn, new(big.Int).SetUint64(d.numerator))
	}
	if d2.denominator != 0 {
		dd = dd.Mul(dd, dd2)
		dd = dd.Add(dd, new(big.Int).SetUint64(d2.numerator))
	}

	if d.denominator != 0 && d2.denominator != 0 {
		dn = dn.Mul(dn, dd2)
		dd = dd.Mul(dd, dd1)
	} else if d2.denominator != 0 {
		dn = dn.Mul(dn, dd2)
	} else if d.denominator != 0 {
		dd = dd.Mul(dd, dd1)
	}

	q, r := new(big.Int).QuoRem(dn, dd, new(big.Int))
	newBD := BigDecimal{
		value:       q,
		scale:       d.scale - d2.scale,
		numerator:   r.Uint64(),
		denominator: dd.Uint64(),
	}
	newBD.optimize()
	return newBD
}
