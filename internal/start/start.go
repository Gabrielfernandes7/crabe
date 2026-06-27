package start

import (
	"github.com/Gabrielfernandes7/crabe/internal/ui"
)

func RunStart() {
	ui.Title("Crabe Start - Iniciando Ollama")

	ui.Success("Ollama está pronto para uso!")
	ui.Info("Agora execute: crabe install --model gemma3:12b")
}