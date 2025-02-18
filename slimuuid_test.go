package slimuuid

import (
	"fmt"
	"sync"
	"testing"
	"time"
	"github.com/google/uuid"
)

const testMacID = "c4:75:ab:cf:66:bf"
const testDate = "2025-02-01"
const testSeed = uint32(12345)
const testCharacters = "0123456789abcdefghijklmnopqrstuvwxyz_-ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Benchmarks
func BenchmarkUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ ,_= uuid.NewV7()
	}
}

func BenchmarkGenerateBest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenerateBest(testMacID)
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
				id1 := GenerateBest(fmt.Sprintf("test1_%d_%d", routineNum, i))
				id2 := GenerateBest(fmt.Sprintf("test2_%d_%d", routineNum, i))
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

// Keep the original sequential test for comparison
func TestSequentialCollisionProbability(t *testing.T) {
	seen := make(map[string]bool)
	iterations := 1000000
	
	startTime := time.Now()
	
	for i := 0; i < iterations; i++ {
		id1 := GenerateBest("test1")
		id2 := GenerateBest("test2")
		
		if seen[id1] || seen[id2] {
			t.Errorf("Collision detected at iteration %d", i)
		}
		seen[id1] = true
		seen[id2] = true
	}

	duration := time.Since(startTime)
	totalIDs := iterations * 2
	idsPerSecond := float64(totalIDs) / duration.Seconds()

	t.Logf("Sequential test completed in %v", duration)
	t.Logf("Total IDs generated: %d", totalIDs)
	t.Logf("IDs per second: %.2f", idsPerSecond)
}
















