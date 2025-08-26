package memvisualizer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"
)

// MemMetrics represents the memory metrics that will be sent as JSON.
type MemMetrics struct {
	Allocated string `json:"allocated"`
}

// indexHandler serves the HTML page with JavaScript.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	htmlContent := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Memory Monitor</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
            background-color: #f4f4f9;
            color: #333;
        }
        .container {
            text-align: center;
            padding: 2rem;
            background-color: #fff;
            border-radius: 10px;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
        }
        h1 {
            color: #007BFF;
            margin-bottom: 1rem;
        }
        #memory-display {
            font-size: 2.5em;
            font-weight: bold;
            color: #28a745;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Allocated Memory</h1>
        <div id="memory-display">Loading...</div>
    </div>

    <script>
        const memoryDisplay = document.getElementById('memory-display');

        // Function that formats the value in bytes into a readable string.
        function formatBytes(bytes, decimals = 2) {
            if (bytes === 0) return '0 Bytes';
            const k = 1024;
            const dm = decimals < 0 ? 0 : decimals;
            const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
            const i = Math.floor(Math.log(bytes) / Math.log(k));
            return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + ' ' + sizes[i];
        }

        // Function to fetch metrics from the API.
        async function fetchMetrics() {
            try {
                const response = await fetch('/metrics');
                const data = await response.json();
                memoryDisplay.textContent = data.allocated;
            } catch (error) {
                console.error('Error fetching metrics:', error);
                memoryDisplay.textContent = 'Failed to load';
            }
        }

        // Updates the metrics every 2 seconds.
        setInterval(fetchMetrics, 2000);
        
        // Makes the first call immediately.
        fetchMetrics();
    </script>
</body>
</html>
`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(htmlContent))
}

// metricsHandler retrieves the memory statistics and returns them as JSON.
func metricsHandler(w http.ResponseWriter, r *http.Request) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	// Creates an instance of our struct and formats the allocated memory.
	metrics := MemMetrics{
		Allocated: FormatMemory(memStats.HeapAlloc),
	}

	// Encodes the struct into JSON and sends it to the client.
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(metrics); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}

// GetAllocatedMemory returns the current allocated heap memory in bytes.
func GetAllocatedMemory() uint64 {
	var memStats runtime.MemStats
	// Fills memStats with the current memory statistics of the process.
	runtime.ReadMemStats(&memStats)

	// Returns the allocated heap memory in bytes.
	return memStats.HeapAlloc
}

// FormatMemory formats a byte value (uint64) into a human-readable string.
func FormatMemory(size uint64) string {
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
	)

	switch {
	case size >= GB:
		return fmt.Sprintf("%.2f GB", float64(size)/float64(GB))
	case size >= MB:
		return fmt.Sprintf("%.2f MB", float64(size)/float64(MB))
	case size >= KB:
		return fmt.Sprintf("%.2f KB", float64(size)/float64(KB))
	default:
		return fmt.Sprintf("%d Bytes", size)
	}
}

// GenerateGraphics starts the web server and the memory consumption goroutine.
func GenerateGraphics() {
	// Increases memory allocation for demonstration purposes.
	// This goroutine makes memory consumption grow over time.
	go func() {
		var data []byte
		for {
			data = append(data, make([]byte, 1024*1024)...)
			time.Sleep(1 * time.Second)
			runtime.KeepAlive(data)
		}
	}()

	// Configures the server routes.
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/metrics", metricsHandler)

	addr := "localhost:8080"
	log.Printf("Server started at http://%s", addr)

	// Starts the HTTP server.
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
