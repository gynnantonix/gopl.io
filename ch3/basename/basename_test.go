package basename

import "testing"

const FILENAME = "/ZErlrga3Jf5p2OPYHvRi/ITESBIIbtJ3SevhuoRbE/dLLBkO8p6HN21QHiTx68.z"

func BenchmarkBasename1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Basename1(FILENAME)
	}
}

func BenchmarkBasename2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Basename2(FILENAME)
	}
}
