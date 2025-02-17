package slimuuid

import (
	"testing"	
)

const testMacID = "c4:75:ab:cf:66:bf"
const testDate = "2025-02-01"
const testSeed = uint32(12345)
const testCharacters = "0123456789abcdefghijklmnopqrstuvwxyz_-ABCDEFGHIJKLMNOPQRSTUVWXYZ"

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


func BenchmarkGenerateWithCharacters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateWithCharacters(testCharacters)
	}
}

func BenchmarkGenerateWithCharactersAndSeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateWithCharactersAndSeed(testCharacters, testSeed)
	}
}

func BenchmarkGenerateWithCharactersAndDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateWithCharactersAndDate(testCharacters, testDate)
	}
}

func BenchmarkGenerateWithCharactersAndDateAndSeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateWithCharactersAndDateAndSeed(testCharacters, testDate, testSeed)
	}
}

func BenchmarkGenerateBestWithCharacters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateBestWithCharacters(testMacID, testCharacters)
	}
}

func BenchmarkGenerateBestWithCharactersAndSeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateBestWithCharactersAndSeed(testMacID, testCharacters, testSeed)
	}
}

func BenchmarkGenerateBestWithCharactersAndDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateBestWithCharactersAndDate(testMacID, testCharacters, testDate)
	}
}

func BenchmarkGenerateBestWithCharactersAndDateAndSeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateBestWithCharactersAndDateAndSeed(testMacID, testCharacters, testDate, testSeed)
	}
}	
















