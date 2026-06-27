package install

import (
	"context"
	"time"

	"github.com/Gabrielfernandes7/crabe/internal/llm"
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
	ollamaClient := llm.NewOllamaClient("http://localhost:11434", "")

	ui.Section("Verificações")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if ollamaClient.Health(ctx) != nil {
		ui.Error("Ollama não está rodando ou não está acessível")
		ui.Info("Dica: Certifique-se de que o Ollama está rodando na porta 11434")
		return
	}
	ui.Success("Ollama está rodando ✓")

	ui.Section("Instalando modelo")
	ui.Info("Baixando modelo → %s", model)
	ui.Info("⏳ Isso pode demorar dependendo da sua internet...")

	pullCtx, pullCancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer pullCancel()

	err := ollamaClient.PullModel(pullCtx, model)
	if err != nil {
		ui.Error("Falha ao instalar o modelo: %v", err)
		return
	}

	ui.Success("✅ Modelo %s instalado com sucesso!", model)
	ui.Info("Agora você pode usar este modelo.")
}

