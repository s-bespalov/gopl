package popcount_test

import (
	"testing"

	"github.com/s-bespalov/gopl/ch2/popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(uint64(i))
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PCLoop(uint64(i))
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PCShift(uint64(i))
	}
}

func BenchmarkPopCountMask(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PCMask(uint64(i))
	}
}
