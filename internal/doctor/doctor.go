package doctor

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/Gabrielfernandes7/crabe/internal/llm"
	"github.com/Gabrielfernandes7/crabe/internal/ui"
)

func Run() {
	ui.Init()
	ui.Title("Crabe Doctor - Diagnóstico do ambiente")

	ui.Section("Ollama")
	
	// Default URL for now, could be configurable
	ollama := llm.NewOllamaClient("http://localhost:11434", "")
	
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if ollama.Health(ctx) == nil {
		ui.Success("Ollama está rodando e acessível")
	} else {
		ui.Error("Ollama não está rodando ou não está acessível em http://localhost:11434")
		ui.Warning("Verifique se o Ollama está instalado e em execução independentemente")
	}

	ui.Section("Porta 11434 (Ollama)")
	if isPortInUse(11434) {
		ui.Success("Porta 11434 está ativa")
	} else {
		ui.Info("Porta 11434 está inativa")
	}
}

// Funções auxiliares
func isPortInUse(port int) bool {
	// lsof não existe em todos os sistemas (ex: Windows), então usamos uma abordagem mais compatível
	cmd := exec.Command("sh", "-c", fmt.Sprintf("lsof -i:%d 2>/dev/null 2>&1 || ss -tuln 2>/dev/null | grep :%d >/dev/null || netstat -tuln 2>/dev/null | grep :%d >/dev/null", port, port, port))
	err := cmd.Run()
	return err == nil
}
