# go-mem-visualizer

![Status do Build](https://img.shields.io/badge/status-in%20progress-blue.svg)
![Licença](https://img.shields.io/badge/license-MIT-green.svg)

---

### O que é o go-mem-visualizer?

O `go-mem-visualizer` é uma ferramenta interativa e visual para análise de performance e consumo de memória de aplicações Go. Ele se conecta ao endpoint de profiling nativo do Go (`pprof`) e transforma dados brutos de alocação de memória em gráficos e relatórios detalhados, ajudando desenvolvedores a identificar rapidamente vazamentos de memória, gargalos de performance e otimizar o uso de recursos.

Diferente da interface padrão do `pprof`, esta ferramenta oferece uma experiência de usuário intuitiva e visualmente rica, permitindo que você:

- Visualize o crescimento e o comportamento da memória heap em tempo real.
- Compare perfis de memória (`diffing`) para identificar alocações que causam leaks.
- Analise flame graphs de forma mais detalhada e interativa.
- Receba sugestões de otimização baseadas em padrões comuns de alocação.

### Por que usar o go-mem-visualizer?

A otimização de memória é crucial para aplicações de alta performance, especialmente aquelas de longa duração. Embora o Go tenha um garbage collector eficiente, alocações desnecessárias podem causar latência e consumo excessivo de recursos. O `go-mem-visualizer` simplifica a tarefa de diagnóstico, tornando a análise de memória acessível e eficiente.

### Como Instalar

```bash
# Baixe a ferramenta diretamente via go get
go install [github.com/](https://github.com/)[seu-usuario]/go-mem-visualizer@latest
