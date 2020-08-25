package code

import (
	"testing"
)

func BenchmarkHasAlternatingBits(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		HasAlternatingBits(i)
	}
}

func BenchmarkHasAlternatingBits2(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		HasAlternatingBits2(i)
	}
}
