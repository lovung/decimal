package decimal

import "testing"

func Benchmark_gcdEuclidean(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gcdEuclidean(4, 3)
		gcdEuclidean(400, 123)
		gcdEuclidean(41265478, 399123)
		gcdEuclidean(102, 2)
		gcdEuclidean(11153, 1230)
		gcdEuclidean(5460, 999122)
	}
}

func Benchmark_gcdRemainderRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gcdRemainderRecursive(4, 3)
		gcdRemainderRecursive(400, 123)
		gcdRemainderRecursive(41265478, 399123)
		gcdRemainderRecursive(102, 2)
		gcdRemainderRecursive(11153, 1230)
		gcdRemainderRecursive(5460, 999122)
	}
}

func Benchmark_gcdRemainder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gcdRemainder(4, 3)
		gcdRemainder(400, 123)
		gcdRemainder(41265478, 399123)
		gcdRemainder(102, 2)
		gcdRemainder(11153, 1230)
		gcdRemainder(5460, 999122)
	}
}
