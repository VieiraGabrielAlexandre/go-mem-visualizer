package main

import (
	"github.com/VieiraGabrielAlexandre/go-mem-visualizer/memvisualizer"
)

func main() {
	// Starts the web application for memory visualization from your library.
	println(memvisualizer.FormatMemory(memvisualizer.GetAllocatedMemory()))
	memvisualizer.GenerateGraphics()
}
