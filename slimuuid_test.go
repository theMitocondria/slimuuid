package slimuuid

import (
	"testing"
)

func BenchmarkSlim(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Generate()
	}
}

func BenchmarkSlimFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenerateFast("c4:75:ab:cf:66:bf")  
	}
}

func BenchmarkSlimBest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenerateBest("c4:75:ab:cf:66:bf")  
	}
}


