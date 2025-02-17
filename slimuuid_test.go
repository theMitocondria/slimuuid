package slimuuid

import (
	"testing"
	"github.com/google/uuid"
	"github.com/matoous/go-nanoid/v2"
	
)

const testMacID = "c4:75:ab:cf:66:bf"
const testDate = "2025-02-01"
const testSeed = uint32(12345)

// Benchmarks
func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Generate()
	}
}

func BenchmarkGenerateWithSeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateWithSeed(testSeed)
	}
}

func BenchmarkGenerateWithDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateWithDate(testDate)
	}
}

func BenchmarkGenerateWithDateAndSeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateWithDateAndSeed(testDate, testSeed)
	}
}

func BenchmarkGenerateBest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenerateBest(testMacID)
	}
}

func BenchmarkGenerateBestWithSeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateBestWithSeed(testMacID, testSeed)
	}
}

func BenchmarkGenerateBestWithDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateBestWithDate(testMacID, testDate)
	}
}

func BenchmarkGenerateBestWithDateAndSeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateBestWithDateAndSeed(testMacID, testDate, testSeed)
	}
}

func BenchmarkUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = uuid.New()
	}
}

func BenchmarkNanoID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = gonanoid.New(21)
	}
}



