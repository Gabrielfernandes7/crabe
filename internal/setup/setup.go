package setup

import (
	"github.com/Gabrielfernandes7/crabe/internal/ui"
	"github.com/spf13/cobra"
)

func NewSetupCmd() *cobra.Command {
	var force bool

	cmd := &cobra.Command{
		Use:   "setup",
		Short: "Provisiona ambiente Ollama",
		Run: func(cmd *cobra.Command, args []string) {
			RunSetup(force)
		},
	}

	cmd.Flags().BoolVarP(&force, "force", "f", false, "Força reinstalação")

	return cmd
}

func RunSetup(force bool) {
	ui.Title("Crabe Setup")

	state := RunPreflight()

	if !state.OllamaRunning {
		ui.Error("Ollama não está rodando")
		return
	}

	model, err := EnsureModel(state.Models)
	if err != nil {
		ui.Error("Erro ao baixar modelo")
		return
	}

	ui.Title("Ambiente pronto")
	ui.Success("Modelo ativo: %s", model)
}
