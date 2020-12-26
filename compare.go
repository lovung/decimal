package decimal

// Cmp compares the bd with ref number
// returns: 1 if bd > ref
//			0 if bd = ref
//		   -1 if bd < ref
func (d BigDecimal) Cmp(ref BigDecimal) int {
	d.ensureInitialized()
	ref.ensureInitialized()

	if d.scale != ref.scale {
		rbd, rref := rescalePair(d, ref)
		return rbd.cmpSameScale(rref)
	}
	return d.cmpSameScale(ref)
}

func (d BigDecimal) cmpSameScale(ref BigDecimal) int {
	valueCmp := d.value.Cmp(ref.value)
	if valueCmp != 0 {
		return valueCmp
	}
	if d.denominator == 0 && ref.denominator == 0 {
		return 0
	}
	if d.denominator == 0 {
		return -1
	}
	if ref.denominator == 0 {
		return 1
	}
	return signInt64(int64(d.numerator*ref.denominator) - int64(d.denominator*ref.numerator))

}

// Equal returns true if d == ref
// and returns false if d != ref
func (d BigDecimal) Equal(ref BigDecimal) bool {
	if d.Cmp(ref) != 0 {
		return false
	}
	return true
}
