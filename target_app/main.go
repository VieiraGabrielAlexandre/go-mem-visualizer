package main

import (
	"log"
	"net/http"
	_ "net/http/pprof" // Importa o pprof para expor os endpoints
	"time"
)

// A simulated memory "leak"
// O "vazamento" de memória simulado
var memorySlices [][]byte

func main() {
	// Start a goroutine to allocate memory
	// Inicia uma goroutine para alocar memória
	go func() {
		for i := 0; i < 100; i++ {
			// Allocate 1 MB of memory
			// Aloca 1 MB de memória
			slice := make([]byte, 1024*1024)
			memorySlices = append(memorySlices, slice)
			log.Printf("Alocando memória, total: %d MB", len(memorySlices))
			time.Sleep(1 * time.Second)
		}
	}()

	// Start the HTTP server to expose pprof endpoints
	// Inicia o servidor HTTP para expor os endpoints do pprof
	log.Println("Servidor de teste iniciado em :6060")
	log.Fatal(http.ListenAndServe(":6060", nil))
}
