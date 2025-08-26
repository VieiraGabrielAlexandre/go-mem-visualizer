package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestGetAllocatedMemorySuccess tests the getAllocatedMemory function for a successful case.
func TestGetAllocatedMemorySuccess(t *testing.T) {
	// Prepare the expected pprof output for the test.
	pprofOutput := `# runtime.MemStats
# HeapInuse = 105848832`

	expectedMemory := uint64(105848832)

	// Create a test HTTP server (mock server) that responds with the pprof output.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/debug/pprof/heap" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(pprofOutput))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	// Call the function to be tested using the test server's URL.
	gotMemory, err := getAllocatedMemory(server.URL)
	if err != nil {
		t.Fatalf("getAllocatedMemory() returned an unexpected error: %v", err)
	}

	// Check if the returned value is the same as the expected value.
	if gotMemory != expectedMemory {
		t.Errorf("getAllocatedMemory() returned %d; expected %d", gotMemory, expectedMemory)
	}
}

// TestGetAllocatedMemoryFail tests the function when the metric is not found.
func TestGetAllocatedMemoryFail(t *testing.T) {
	// Prepare a pprof output that does not contain the expected metric.
	pprofOutput := `
# runtime.MemStats
# HeapReleased = 2146304
`

	// Create a test HTTP server (mock server).
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/debug/pprof/heap" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(pprofOutput))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	// Call the function and expect an error.
	_, err := getAllocatedMemory(server.URL)
	if err == nil {
		t.Errorf("getAllocatedMemory() did not return an error, but it should have failed")
	}

	// Check if the error message contains the expected substring.
	if !strings.Contains(err.Error(), "could not find memory metric") {
		t.Errorf("unexpected error message: %v", err)
	}
}
