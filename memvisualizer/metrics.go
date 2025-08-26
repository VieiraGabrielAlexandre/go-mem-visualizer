package memvisualizer

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// GetAllocatedMemory makes a request to the pprof URL and returns the allocated memory value.
func GetAllocatedMemory(pprofURL string) (uint64, error) {
	// Makes an HTTP request to the target application's pprof/heap endpoint.
	resp, err := http.Get(pprofURL + "/debug/pprof/heap?debug=1")
	if err != nil {
		return 0, fmt.Errorf("error connecting to the target application: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected response from the target application: %d", resp.StatusCode)
	}

	// Reads the response body and looks for the allocated memory value.
	scanner := bufio.NewScanner(resp.Body)
	var allocatedMemory uint64

	for scanner.Scan() {
		line := scanner.Text()

		// Remove the comment prefix if it exists
		line = strings.TrimPrefix(line, "# ")

		// Looks for the HeapInuse metric
		if strings.Contains(line, "HeapInuse") {
			fields := strings.Fields(line)
			if len(fields) >= 3 && fields[0] == "HeapInuse" && fields[1] == "=" {
				// Extracts the numeric value using strconv.ParseUint
				if val, err := strconv.ParseUint(fields[2], 10, 64); err == nil {
					allocatedMemory = val
					return allocatedMemory, nil
				}
			}
		}
	}

	return 0, fmt.Errorf("could not find memory metric in pprof response")
}
