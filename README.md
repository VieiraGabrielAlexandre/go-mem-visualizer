# go-mem-visualizer

![Build Status](https://img.shields.io/badge/status-stable-blue.svg)
![License](https://img.shields.io/badge/license-MIT-green.svg)

---

### 📖 What is `go-mem-visualizer`?

`go-mem-visualizer` is a lightweight Go library for **inspecting and visualizing memory usage** in your applications.
It eliminates the need to configure `pprof` servers or manually parse memory stats, giving you both **real-time metrics** and an **optional web dashboard**.

With it, you can:

* 📊 Retrieve real-time memory allocation with a single line of code.
* 🌐 Launch a local **web interface** to visualize heap memory usage in real-time.
* 🔎 Automatically format memory values into human-readable units (`KB`, `MB`, `GB`).

---

### ⚡ Why use it?

Memory optimization is crucial for Go applications, especially long-running or resource-intensive ones.
While Go’s garbage collector is efficient, unnecessary allocations can lead to **higher latency** and **excessive memory consumption**.

`go-mem-visualizer` makes diagnosis simple by exposing memory metrics in a **developer-friendly way**: either as raw values or through a **live dashboard**.

---

### 📦 Installation

```bash
go get github.com/VieiraGabrielAlexandre/go-mem-visualizer
```

---

### 🛠 Usage

#### 1. Get and Format Memory Metrics

Retrieve allocated memory and format it into a human-readable string:

```go
package main

import (
    "fmt"
    "time"

    "github.com/VieiraGabrielAlexandre/go-mem-visualizer/memvisualizer"
)

func main() {
    var data []byte

    for i := 0; i < 5; i++ {
        // Simulate memory consumption
        data = append(data, make([]byte, 1024*1024)...)

        allocated := memvisualizer.GetAllocatedMemory()
        formatted := memvisualizer.FormatMemory(allocated)

        fmt.Printf("Allocated memory: %s\n", formatted)

        time.Sleep(2 * time.Second)
    }
}
```

---

#### 2. Start the Monitoring Web Dashboard

Spin up a lightweight local server with an interactive dashboard:

```go
package main

import (
	"github.com/VieiraGabrielAlexandre/go-mem-visualizer/memvisualizer"
)

func main() {
	// Start the web server and visualization
	memvisualizer.GenerateGraphics()
}
```

Open [http://localhost:8080](http://localhost:8080) to see heap memory usage in real time.

<img width="887" height="442" alt="image" src="https://github.com/user-attachments/assets/baf9497c-d599-41c8-ad48-8c0f51389509" />


---

### 🧪 Running Tests & Benchmarks

```bash
# Run unit tests
go test ./...

# Run benchmarks
go test -bench=. ./...
```

---

### 🤝 Contributing

Contributions are welcome!
Feel free to **open an issue** or **submit a pull request** on the [GitHub repository](https://github.com/VieiraGabrielAlexandre/go-mem-visualizer).

---

### 📜 License

This project is licensed under the [MIT License](./LICENSE).
