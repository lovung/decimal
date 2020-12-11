package num

import "math/big"

// GCDEuclidean calculates GCD by Euclidian algorithm.
func GCDEuclidean(a, b big.Int) big.Int {
	na, nb := a, b
	pa, pb := &na, &nb
	for pa.Cmp(pb) != 0 {
		if pa.Cmp(pb) > 0 {
			pa = pa.Sub(pa, pb)
		} else {
			pb = pb.Sub(pb, pa)
		}
	}

	return na
}

// GCDRemainderRecursive calculates GCD recursively using remainder.
func GCDRemainderRecursive(a, b big.Int) big.Int {
	na, nb := a, b
	pa := gcdRemainderRecursive(&na, &nb)
	return *pa
}

func gcdRemainderRecursive(a, b *big.Int) *big.Int {
	if b.Cmp(zeroInt) == 0 {
		return a
	}
	return gcdRemainderRecursive(b, a.Mod(a, b))
}

// GCDRemainder calculates GCD iteratively using remainder.
func GCDRemainder(a, b big.Int) big.Int {
	if b.Cmp(zeroInt) != 0 {
		c := b
		pb := a.Mod(&a, &b)
		b = *pb
		a = c
	}

	return a
}
