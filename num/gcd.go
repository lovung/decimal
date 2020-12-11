package num

import "math/big"

// GCDEuclidean calculates GCD by Euclidian algorithm.
func GCDEuclidean(a, b big.Int) big.Int {
	na, nb := new(big.Int).Set(&a), new(big.Int).Set(&b)
	for na.Cmp(nb) != 0 {
		if na.Cmp(nb) > 0 {
			na.Sub(na, nb)
		} else {
			nb.Sub(nb, na)
		}
	}

	return *na
}

// GCDRemainderRecursive calculates GCD recursively using remainder.
func GCDRemainderRecursive(a, b big.Int) big.Int {
	na, nb := new(big.Int).Set(&a), new(big.Int).Set(&b)
	pa := gcdRemainderRecursive(na, nb)
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
	var c = new(big.Int)
	for b.Cmp(zeroInt) != 0 {
		c.Set(&b)
		pb := a.Mod(&a, &b)
		b = *pb
		a = *c
	}

	return a
}
