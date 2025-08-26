package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// getAllocatedMemory makes a request to the pprof URL and returns the allocated memory value.
func getAllocatedMemory(pprofURL string) (uint64, error) {
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

func main() {
	// Define the flag for the target application's pprof URL
	targetPprofURL := flag.String("pprof-url", "http://localhost:6060", "Go application URL to be analyzed (with pprof enabled)")
	flag.Parse()

	// Creates a Gin instance with default middlewares
	r := gin.Default()

	// CORS Middleware configuration to allow any origin
	r.Use(cors.Default())

	// Route for the metrics API
	r.GET("/api/metrics", func(c *gin.Context) {
		// Calls the function to get the memory metric
		allocatedMemory, err := getAllocatedMemory(*targetPprofURL)
		if err != nil {
			log.Println(err)
			// Returns the error based on its type
			if strings.Contains(err.Error(), "could not find") {
				c.JSON(http.StatusNotFound, gin.H{"error": "Memory metric not found in pprof response"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get memory metrics"})
			}
			return
		}

		// Returns the data in JSON format
		c.JSON(http.StatusOK, gin.H{
			"timestamp":       time.Now().UnixMilli(),
			"allocatedMemory": allocatedMemory,
		})
	})

	// Serves static files (CSS, JS) from the "web" directory at the /static route
	r.Static("/static", "./web")

	// Route for the home page (serving index.html)
	r.GET("/", func(c *gin.Context) {
		indexPath := filepath.Join("web", "index.html")
		c.File(indexPath)
	})

	log.Printf("go-mem-visualizer server started at http://localhost:8080")
	log.Printf("Connecting to pprof URL: %s", *targetPprofURL)

	// Starts the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
