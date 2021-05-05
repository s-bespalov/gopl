package popcount_test

import (
	"math/rand"
	"testing"

	"github.com/s-bespalov/gopl/ch2/Exercise-2.3-2.5/popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(rand.Uint64())
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountLoop(rand.Uint64())
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountShift(rand.Uint64())
	}
}

func BenchmarkPopCountClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountClear(rand.Uint64())
	}
}
