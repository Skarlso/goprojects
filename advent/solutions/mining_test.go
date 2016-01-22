package solutions

import "testing"

func BenchmarkMining(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mine()
	}
}
