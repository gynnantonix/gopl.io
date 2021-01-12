package basename

import "testing"

func BenchmarkBasename(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Basename1("/ZErlrga3Jf5p2OPYHvRi/ITESBIIbtJ3SevhuoRbE/dLLBkO8p6HN21QHiTx68.z")
	}
}
