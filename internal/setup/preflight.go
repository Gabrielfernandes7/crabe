package setup

import (
	"context"
	"time"

	"github.com/Gabrielfernandes7/crabe/internal/llm"
	"github.com/Gabrielfernandes7/crabe/internal/ui"
)

func RunPreflight() SystemState {
	ui.Section("Preflight")

	state := SystemState{}
	
	ollama := llm.NewOllamaClient("http://localhost:11434", "")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Ollama running check
	state.OllamaRunning = ollama.Health(ctx) == nil
	if state.OllamaRunning {
		ui.Success("Ollama ativo")
	} else {
		ui.Warning("Ollama não está rodando")
	}

	state.Models = listModels()

	return state
}
