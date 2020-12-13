package decimal

// Cmp compares the bd with ref number
// returns: 1 if bd > ref
//			0 if bd = ref
//		   -1 if bd < ref
func (bd BigDecimal) Cmp(ref BigDecimal) int {
	bd.ensureInitialized()
	ref.ensureInitialized()

	if bd.scale != ref.scale {
		rbd, rref := RescalePair(bd, ref)
		return rbd.cmpSameScale(rref)
	}
	return bd.cmpSameScale(ref)
}

func (bd BigDecimal) cmpSameScale(ref BigDecimal) int {
	valueCmp := bd.value.Cmp(ref.value)
	if valueCmp != 0 {
		return valueCmp
	}
	if bd.denominator == 0 && ref.denominator == 0 {
		return 0
	}
	if bd.denominator == 0 {
		return -1
	}
	if ref.denominator == 0 {
		return 1
	}
	return signInt64(int64(bd.numerator*ref.denominator) - int64(bd.denominator*ref.numerator))

}
