package install

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Gabrielfernandes7/crabe/internal/ui"
	"github.com/spf13/cobra"
)

func NewInstallCmd() *cobra.Command {
	var model string

	cmd := &cobra.Command{
		Use:   "install",
		Short: "Instala modelos no Ollama",
		Long:  `Baixa e instala modelos do Ollama (ex: llama3.2:1b, qwen2.5:7b, etc).`,
		Run: func(cmd *cobra.Command, args []string) {
			if model == "" {
				ui.Error("Você deve informar um modelo.")
				ui.Info("Exemplo: crabe install --model llama3.2:1b")
				return
			}
			RunInstall(model)
		},
	}

	cmd.Flags().StringVarP(&model, "model", "m", "", "Nome do modelo a ser instalado")
	return cmd
}

func RunInstall(model string) {
	ui.Title("Crabe Install - Instalando modelo")

	ui.Section("Verificações")
	if !isOllamaRunning() {
		ui.Error("Ollama não está rodando")
		ui.Info("Dica: Rode 'crabe start' primeiro")
		return
	}
	ui.Success("Ollama está rodando ✓")

	ui.Section("Instalando modelo")
	ui.Info(fmt.Sprintf("Baixando modelo → %s", model))
	ui.Info("⏳ Isso pode demorar dependendo da sua internet...")

	err := pullModelWithOutput(model)
	if err != nil {
		ui.Error("Falha ao instalar o modelo")
		return
	}

	ui.Success(fmt.Sprintf("✅ Modelo %s instalado com sucesso!", model))
	ui.Info("Agora você pode usar este modelo no Open WebUI ou em qualquer cliente Ollama.")
}

func isOllamaRunning() bool {
	for _, useSudo := range []bool{false, true} {
		cmd := exec.Command("docker", "ps", "--filter", "name=ollama", "--format", "{{.Status}}")
		if useSudo {
			cmd = exec.Command("sudo", "docker", "ps", "--filter", "name=ollama", "--format", "{{.Status}}")
		}

		out, err := cmd.CombinedOutput()
		if err == nil && strings.Contains(strings.ToLower(string(out)), "up") {
			return true
		}
	}
	return false
}

func pullModelWithOutput(model string) error {
	cmd := exec.Command("sudo", "docker", "exec", "ollama", "ollama", "pull", model)

	// Mostra o progresso real do download
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
