package slimuuid

import (
	"testing"
)

func BenchmarkSlimUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Generate()
	}
}

func BenchmarkSlimUUIDFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenerateFast("mitocondria")  
	}
}
