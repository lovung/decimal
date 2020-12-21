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

	// We have:
	//		D := (v1 + n1/d1) * 10^(-s1)
	//		D2 := (v2 + n2/d2) * 10^(-s2)
	// So:
	//   D      v1 + n1/d1
	// ---- = ----------------  * 10^-(s1-s2)					(1)
	//  D2      v2 + n2/d2
	//
	//       D      (v1 * d1 + n1) * d2
	// <=> ---- = -----------------------  * 10^-(s1-s2)		(2)
	//      D2      (v2 * d2 + n2) * d1
	var (
		// Set dn = v1, dd = v2
		dn = new(big.Int).Set(d.value)  // Decimal numerator
		dd = new(big.Int).Set(d2.value) // Decimal denominator
		// Set dd1 = d1, dd2 = d2
		dd1 = new(big.Int).SetUint64(d.denominator)  // Decimal of d1
		dd2 = new(big.Int).SetUint64(d2.denominator) // Decimal of d2
	)

	// If d1 != 0, we can calculate dn = v1 * d1 + n1			(3)
	if d.denominator != 0 {
		dn = dn.Mul(dn, dd1)
		dn = dn.Add(dn, new(big.Int).SetUint64(d.numerator))
	}
	// If d2 != 0, we can calculate dd = v2 * d2 + n2			(4)
	if d2.denominator != 0 {
		dd = dd.Mul(dd, dd2)
		dd = dd.Add(dd, new(big.Int).SetUint64(d2.numerator))
	}

	if d.denominator != 0 && d2.denominator != 0 {
		// This case d1 and d2 != 0 both
		// We calculate dn and dd base on the (2) calculation above
		dn = dn.Mul(dn, dd2)
		dd = dd.Mul(dd, dd1)
	} else if d2.denominator != 0 {
		// In case, d2 != 0 and d1 == 0
		//   D        v1 * d2
		// ---- = --------------								(5)
		//  D2     v2 * d2 + n2
		// => So we keep the dd as (4) above
		dn = dn.Mul(dn, dd2)
	} else if d.denominator != 0 {
		// In case, d1 != 0 and d2 == 0
		//   D      v1 * d1 + n1
		// ---- = --------------								(6)
		//  D2      v2 * d1
		// => So we keep the dn as (3) above
		dd = dd.Mul(dd, dd1)
	}

	// q and r are quotient and remainder
	// helps to calculate the D/D2 based on (2) or (5) or (6)
	// q will be saved as value
	// r will be saved as numerator of result
	q, r := new(big.Int).QuoRem(dn, dd, new(big.Int))
	newBD := BigDecimal{
		value:       q,
		scale:       d.scale - d2.scale, // based on (2)
		numerator:   r.Uint64(),
		denominator: dd.Uint64(),
	}
	newBD.optimize()
	return newBD
}
