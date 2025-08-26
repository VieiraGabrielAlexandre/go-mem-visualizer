package memvisualizer_test

import (
	"github.com/VieiraGabrielAlexandre/go-mem-visualizer/memvisualizer"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestGetAllocatedMemorySuccess testa a função getAllocatedMemory para um caso de sucesso.
func TestGetAllocatedMemorySuccess(t *testing.T) {
	// Prepara a saída esperada do pprof para o teste.
	pprofOutput := `# runtime.MemStats
# HeapInuse = 105848832`

	expectedMemory := uint64(105848832)

	// Cria um servidor HTTP de teste (mock server) que responde com a saída do pprof.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/debug/pprof/heap" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(pprofOutput))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	// Chama a função a ser testada usando a URL do servidor de teste.
	gotMemory, err := memvisualizer.GetAllocatedMemory(server.URL)
	if err != nil {
		t.Fatalf("getAllocatedMemory() retornou um erro inesperado: %v", err)
	}

	// Verifica se o valor retornado é o mesmo que o esperado.
	if gotMemory != expectedMemory {
		t.Errorf("getAllocatedMemory() retornou %d; esperava %d", gotMemory, expectedMemory)
	}
}

// TestGetAllocatedMemoryFail testa a função quando a métrica não é encontrada.
func TestGetAllocatedMemoryFail(t *testing.T) {
	// Prepara uma saída de pprof que não contém a métrica esperada.
	pprofOutput := `
# runtime.MemStats
# HeapReleased = 2146304
`

	// Cria um servidor HTTP de teste (mock server).
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/debug/pprof/heap" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(pprofOutput))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	// Chama a função e espera um erro.
	_, err := memvisualizer.GetAllocatedMemory(server.URL)
	if err == nil {
		t.Errorf("getAllocatedMemory() did not return an error, but it should have failed")
	}

	// Verifica se a mensagem de erro contém a substring esperada.
	if !strings.Contains(err.Error(), "could not find memory metric") {
		t.Errorf("unexpected error message: %v", err)
	}
}
