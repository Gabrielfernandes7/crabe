package setup

import (
	"os/exec"
	"strings"

	"github.com/Gabrielfernandes7/crabe/internal/ui"
)

func RunPreflight() SystemState {
	ui.Section("Preflight")

	state := SystemState{}

	// Docker
	if _, err := exec.LookPath("docker"); err == nil {
		ui.Success("Docker instalado")
		state.DockerAvailable = true
	} else {
		ui.Error("Docker não encontrado")
		return state
	}

	// Docker engine
	if err := exec.Command("docker", "info").Run(); err == nil {
		ui.Success("Docker está rodando")
		state.DockerRunning = true
	} else {
		ui.Error("Docker não está rodando")
	}

	// Ollama container
	state.OllamaRunning = isOllamaRunning()
	if state.OllamaRunning {
		ui.Success("Ollama container ativo")
	} else {
		ui.Warning("Ollama não está rodando")
	}

	state.Models = listModels()

	return state
}

func isOllamaRunning() bool {
	out, err := exec.Command("docker", "ps", "--format", "{{.Names}} {{.Image}}").Output()
	if err != nil {
		return false
	}
	return strings.Contains(strings.ToLower(string(out)), "ollama")
}
