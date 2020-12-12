package decimal

import "testing"

// Arrange
var cases = []struct {
	expected, a, b int64
}{
	{1, 4, 3},
	{10, 40, 30},
}

var invalidCases = []struct {
	expected, a, b int64
}{
	{2, 4, 3},
	{5, 40, 30},
}

func actAndAssert(t *testing.T, fp func(int64, int64) int64) {
	for _, c := range cases {
		// Act
		actual := fp(c.a, c.b)

		// Assert
		if actual != c.expected {
			t.Errorf("gcd(%d, %d) == %d, expected %d", c.a, c.b, actual, c.expected)
		}
	}
}

func actAndAssertWithInvalidData(t *testing.T, fp func(int64, int64) int64) {
	for _, c := range invalidCases {
		// Act
		actual := fp(c.a, c.b)

		// Assert
		if actual == c.expected {
			t.Errorf("gcd(%d, %d) == %d, expected %d", c.a, c.b, actual, c.expected)
		}
	}
}

func TestgcdEuclidean(t *testing.T) {
	actAndAssert(t, gcdEuclidean)
}

func TestgcdRemainderRecursive(t *testing.T) {
	actAndAssert(t, gcdRemainderRecursive)
}

func TestgcdRemainder(t *testing.T) {
	actAndAssert(t, gcdRemainder)
}

func TestgcdEuclideanWithInvalidData(t *testing.T) {
	actAndAssertWithInvalidData(t, gcdEuclidean)
}

func TestgcdRemainderRecursiveWithInvalidData(t *testing.T) {
	actAndAssertWithInvalidData(t, gcdRemainderRecursive)
}

func TestgcdRemainderWithInvalidData(t *testing.T) {
	actAndAssertWithInvalidData(t, gcdRemainder)
}
