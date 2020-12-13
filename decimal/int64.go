package decimal

import "math"

func tenPow(exponent uint64) int64 {
	return int64(math.Pow(10, float64(exponent)))
}

func sumFraction(n1, d1, n2, d2 int64) (int64, int64) {
	if d1 == 0 || d2 == 0 {
		panic(ErrZeroDenominator)
	}
	demGCD := gcd(d1, d2)
	return (n1*d2 + n2*d1) / demGCD, d1 * d2 / demGCD
}

var gcd = gcdEuclidean

// gcdEuclidean calculates GCD by Euclidian algorithm.
func gcdEuclidean(a, b int64) int64 {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

// gcdRemainderRecursive calculates GCD recursively using remainder.
func gcdRemainderRecursive(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcdRemainderRecursive(b, a%b)
}

// gcdRemainder calculates GCD iteratively using remainder.
func gcdRemainder(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

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
