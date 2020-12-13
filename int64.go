package decimal

// signInt64 returns the sign of an int64 number
//		1 if ref > 0
//		0 if ref = 0
//	   -1 if ref < 0
func signInt64(ref int64) int {
	if ref == 0 {
		return 0
	}
	if ref < 0 {
		return -1
	}
	return 1
}

func absInt64(ref int64) int64 {
	if ref < 0 {
		return -ref
	}
	return ref
}
