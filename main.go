package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/VieiraGabrielAlexandre/go-mem-visualizer/memvisualizer"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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
		// Calls the function from the new package
		allocatedMemory, err := memvisualizer.GetAllocatedMemory(*targetPprofURL)
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
