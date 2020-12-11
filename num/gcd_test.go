package num

import (
	"math/big"
	"testing"
)

// Arrange
var cases = []struct {
	expected, a, b *big.Int
}{
	{new(big.Int).Set(oneInt), new(big.Int).Set(fourInt), new(big.Int).Set(threeInt)},
	{new(big.Int).Set(tenInt), new(big.Int).Set(fourtyInt), new(big.Int).Set(thirdtyInt)},
}

var invalidCases = []struct {
	expected, a, b *big.Int
}{
	{new(big.Int).Set(twoInt), new(big.Int).Set(fourInt), new(big.Int).Set(threeInt)},
	{new(big.Int).Set(fiveInt), new(big.Int).Set(fourtyInt), new(big.Int).Set(thirdtyInt)},
}

func actAndAssert(t *testing.T, fp func(big.Int, big.Int) big.Int) {
	for _, c := range cases {
		// Act
		actual := fp(*c.a, *c.b)

		// Assert
		if actual.Cmp(c.expected) != 0 {
			t.Errorf("GCD(%s, %s) == %s, expected %s", c.a.String(), c.b.String(), actual.String(), c.expected.String())
		}
	}
}

func actAndAssertWithInvalidData(t *testing.T, fp func(big.Int, big.Int) big.Int) {
	for _, c := range invalidCases {
		// Act
		actual := fp(*c.a, *c.b)

		// Assert
		if actual.Cmp(c.expected) == 0 {
			t.Errorf("GCD(%s, %s) == %s, expected %s", c.a.String(), c.b.String(), actual.String(), c.expected.String())
		}
	}
}

func BenchmarkGCDEuclidean(b *testing.B) {
	// run the Fib function b.N times
	first := new(big.Int).SetInt64(400)
	second := new(big.Int).SetInt64(284)
	for n := 0; n < b.N; n++ {
		GCDEuclidean(*first, *second)
	}
}

func BenchmarkGCDRemainderRecursive(b *testing.B) {
	// run the Fib function b.N times
	first := new(big.Int).SetInt64(400)
	second := new(big.Int).SetInt64(284)
	for n := 0; n < b.N; n++ {
		GCDRemainderRecursive(*first, *second)
	}
}

func BenchmarkGCDRemainder(b *testing.B) {
	// run the Fib function b.N times
	first := new(big.Int).SetInt64(400)
	second := new(big.Int).SetInt64(284)
	for n := 0; n < b.N; n++ {
		GCDRemainder(*first, *second)
	}
}

func TestGCDEuclidean(t *testing.T) {
	actAndAssert(t, GCDEuclidean)
}

func TestGCDRemainderRecursive(t *testing.T) {
	actAndAssert(t, GCDRemainderRecursive)
}

func TestGCDRemainder(t *testing.T) {
	actAndAssert(t, GCDRemainder)
}

func TestGCDEuclideanWithInvalidData(t *testing.T) {
	actAndAssertWithInvalidData(t, GCDEuclidean)
}

func TestGCDRemainderRecursiveWithInvalidData(t *testing.T) {
	actAndAssertWithInvalidData(t, GCDRemainderRecursive)
}

func TestGCDRemainderWithInvalidData(t *testing.T) {
	actAndAssertWithInvalidData(t, GCDRemainder)
}
