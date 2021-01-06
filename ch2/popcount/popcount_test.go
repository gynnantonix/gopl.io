package popcount

import (
	"math/rand"
	"testing"
)

var testVal = rand.Uint64()

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(testVal)
	}
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop(testVal)
	}
}

func BenchmarkDrop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Drop(testVal)
	}
}
