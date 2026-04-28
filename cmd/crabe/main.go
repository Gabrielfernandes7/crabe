// cmd/crabe/main.go
package main

import (
	"os"

	"github.com/Gabrielfernandes7/crabe/internal/doctor"
	"github.com/Gabrielfernandes7/crabe/internal/initcmd"
	"github.com/Gabrielfernandes7/crabe/internal/inspect"
	"github.com/Gabrielfernandes7/crabe/internal/install"
	"github.com/Gabrielfernandes7/crabe/internal/setup"
	"github.com/Gabrielfernandes7/crabe/internal/start"
	"github.com/Gabrielfernandes7/crabe/internal/ui"
	"github.com/Gabrielfernandes7/crabe/internal/uninstall"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crabe",
	Short: "🦀 Crabe CLI - Ambiente de IA local com Ollama",
	Long:  `Ferramenta para facilitar o uso de Ollama 100% local no contexto do seu projeto.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		ui.Init()
	},
}

func init() {
	rootCmd.AddCommand(doctor.NewDoctorCmd())
	rootCmd.AddCommand(initcmd.NewInitCmd())
	rootCmd.AddCommand(inspect.NewInspectCmd())
	rootCmd.AddCommand(install.NewInstallCmd())
	rootCmd.AddCommand(setup.NewSetupCmd())
	rootCmd.AddCommand(uninstall.NewUninstallCmd())
	rootCmd.AddCommand(start.NewStartCmd())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		ui.Error("Erro ao executar comando: %v", err)
		os.Exit(1)
	}
}
