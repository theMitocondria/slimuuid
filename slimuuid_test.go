package slimuuid

import (
	"fmt"
	"sync"
	"testing"
	"time"
	"github.com/google/uuid"
	"github.com/matoous/go-nanoid/v2"
)

const testMacID = "c4:75:ab:cf:66:bf"
const testDate = "2025-02-01"
const testSeed = uint32(12345)
const testCharacters = "0123456789abcdefghijklmnopqrstuvwxyz_-ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Benchmarks
func BenchmarkGenerateBest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_,_ = GenerateBest(testMacID)
	}
}

func BenchmarkUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_,_= uuid.NewV7()
	}
}

func BenchmarkNanoid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_,_= gonanoid.New()
	}
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_,_= Generate()
	}
}

func BenchmarkGenerateWithDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_,_= GenerateWithDate(testDate)
	}
}

func BenchmarkGenerateWithCharacters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_,_= GenerateWithCharacters(testCharacters)
	}
}

func BenchmarkGenerateWithCharactersAndDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_,_= GenerateWithCharactersAndDate(testCharacters, testDate)
	}
}

func BenchmarkGenerateBestWithDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_,_= GenerateBestWithDate(testMacID, testDate)
	}
}

func BenchmarkGenerateBestWithCharacters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_,_= GenerateBestWithCharacters(testMacID, testCharacters)
	}
}

func BenchmarkGenerateBestWithCharactersAndDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_,_= GenerateBestWithCharactersAndDate(testMacID, testCharacters, testDate)
	}
}

func TestParallelCollisionProbability(t *testing.T) {
	const (
		numGoroutines = 5000    // Number of parallel requests
		idsPerRoutine = 5000    // IDs generated per routine
	)

	var (
		wg sync.WaitGroup
		mu sync.Mutex
		seen = make(map[string]bool)
		conflicts = 0
	)

	startTime := time.Now()

	// Launch goroutines
	for r := 0; r < numGoroutines; r++ {
		wg.Add(1)
		go func(routineNum int) {
			defer wg.Done()
			
			// Local map to store IDs before checking globally
			localIds := make([]string, 0, idsPerRoutine*2) // *2 because we generate 2 IDs per iteration
			
			// Generate IDs
			for i := 0; i < idsPerRoutine; i++ {
				id1,_ := GenerateBest(fmt.Sprintf("test1_%s_%d", testMacID, i))
				id2,_ := GenerateBest(fmt.Sprintf("test2_%s_%d", testMacID, i))
				localIds = append(localIds, id1, id2)
			}

			// Check for collisions in global map
			mu.Lock()
			for _, id := range localIds {
				if seen[id] {
					conflicts++
					t.Errorf("Collision detected: %s", id)
				}
				seen[id] = true
			}
			mu.Unlock()
		}(r)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	
	duration := time.Since(startTime)
	totalIDs := numGoroutines * idsPerRoutine * 2 // *2 because we generate 2 IDs per iteration
	idsPerSecond := float64(totalIDs) / duration.Seconds()

	t.Logf("Test completed in %v", duration)
	t.Logf("Total IDs generated: %d", totalIDs)
	t.Logf("IDs per second: %.2f", idsPerSecond)
	t.Logf("Conflicts detected: %d", conflicts)
	t.Logf("Unique ratio: %.4f%%", 100*(1-float64(conflicts)/float64(totalIDs)))
}
















