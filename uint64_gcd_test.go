package decimal

import "testing"

// Arrange
var cases = []struct {
	expected, a, b uint64
}{
	{1, 4, 3},
	{10, 40, 30},
}

var invalidCases = []struct {
	expected, a, b uint64
}{
	{2, 4, 3},
	{5, 40, 30},
}

func actAndAssert(t *testing.T, fp func(uint64, uint64) uint64) {
	for _, c := range cases {
		// Act
		actual := fp(c.a, c.b)

		// Assert
		if actual != c.expected {
			t.Errorf("gcd(%d, %d) == %d, expected %d", c.a, c.b, actual, c.expected)
		}
	}
}

func actAndAssertWithInvalidData(t *testing.T, fp func(uint64, uint64) uint64) {
	for _, c := range invalidCases {
		// Act
		actual := fp(c.a, c.b)

		// Assert
		if actual == c.expected {
			t.Errorf("gcd(%d, %d) == %d, expected %d", c.a, c.b, actual, c.expected)
		}
	}
}

func Test_gcdEuclidean(t *testing.T) {
	actAndAssert(t, gcdEuclidean)
}

func Test_gcdRemainderRecursive(t *testing.T) {
	actAndAssert(t, gcdRemainderRecursive)
}

func Test_gcdRemainder(t *testing.T) {
	actAndAssert(t, gcdRemainder)
}

func Test_gcdEuclideanWithInvalidData(t *testing.T) {
	actAndAssertWithInvalidData(t, gcdEuclidean)
}

func Test_gcdRemainderRecursiveWithInvalidData(t *testing.T) {
	actAndAssertWithInvalidData(t, gcdRemainderRecursive)
}

func Test_gcdRemainderWithInvalidData(t *testing.T) {
	actAndAssertWithInvalidData(t, gcdRemainder)
}
