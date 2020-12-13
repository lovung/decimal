package decimal

func minInt32(x, y int32) int32 {
	if x >= y {
		return y
	}
	return x
}

func maxInt32(x, y int32) int32 {
	if x <= y {
		return y
	}
	return x
}
