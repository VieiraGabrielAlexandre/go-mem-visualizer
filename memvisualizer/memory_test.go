package memvisualizer_test

import (
	"runtime"
	"testing"
	"time"

	// Import your library package for testing.
	"github.com/VieiraGabrielAlexandre/go-mem-visualizer/memvisualizer"
)

// TestGetAllocatedMemory checks if the function from your library
// returns an allocated memory value greater than zero.
func TestGetAllocatedMemory(t *testing.T) {
	// Get the initial allocated memory value before the test.
	var initialMem runtime.MemStats
	runtime.ReadMemStats(&initialMem)

	// Define an allocation size for the test (10 MB).
	const allocationSize = 10 * 1024 * 1024

	// Allocate memory.
	testData := make([]byte, allocationSize)

	// Ensure the memory is not garbage collected immediately.
	runtime.KeepAlive(testData)
	time.Sleep(100 * time.Millisecond) // Small delay to allow the system to update metrics.

	// Call the function from your library to get the allocated memory.
	currentAllocated := memvisualizer.GetAllocatedMemory()

	// Check if the current allocation is greater than the initial one.
	if currentAllocated <= initialMem.HeapAlloc {
		t.Errorf("Expected allocated memory to be greater than %d bytes, but got %d bytes.", initialMem.HeapAlloc, currentAllocated)
	}
}

// BenchmarkGetAllocatedMemory measures the performance of the GetAllocatedMemory function.
// To run, use: go test -bench=.
func BenchmarkGetAllocatedMemory(b *testing.B) {
	// Define an allocation size for the test (10 MB).
	const allocationSize = 10 * 1024 * 1024
	testData := make([]byte, allocationSize)
	runtime.KeepAlive(testData)

	// The b.N loop is automatically generated and optimizes the number of iterations
	// to obtain an accurate measurement.
	for i := 0; i < b.N; i++ {
		// This is the line we're measuring performance for.
		_ = memvisualizer.GetAllocatedMemory()
	}
}
