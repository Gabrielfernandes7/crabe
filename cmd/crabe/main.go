// cmd/crabe/main.go
package main

import (
	"os"

	"github.com/Gabrielfernandes7/crabe/internal/doctor"
	"github.com/Gabrielfernandes7/crabe/internal/initcmd"
	"github.com/Gabrielfernandes7/crabe/internal/inspect"
	"github.com/Gabrielfernandes7/crabe/internal/install"
	"github.com/Gabrielfernandes7/crabe/internal/setup"
	"github.com/Gabrielfernandes7/crabe/internal/ui"
	"github.com/Gabrielfernandes7/crabe/internal/ui/terminal"
	"github.com/Gabrielfernandes7/crabe/internal/uninstall"
	"github.com/Gabrielfernandes7/crabe/internal/workspace"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crabe",
	Short: "🦀 Crabe - Business Agent Runtime",
	Long:  `Crabe é um Agente de IA local especializado em negócios, validação de ideias e pesquisa de mercado.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		ui.Init()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			model, err := workspace.LoadConfig()
			if err != nil {
				model = "llama3.1:8b" // default fallback
			}
			terminal.Run(model)
			return
		}
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(doctor.NewDoctorCmd())
	rootCmd.AddCommand(initcmd.NewInitCmd())
	rootCmd.AddCommand(inspect.NewInspectCmd())
	rootCmd.AddCommand(install.NewInstallCmd())
	rootCmd.AddCommand(setup.NewSetupCmd())
	rootCmd.AddCommand(workspace.NewTrustCmd())
	rootCmd.AddCommand(uninstall.NewUninstallCmd())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		ui.Error("Erro ao executar comando: %v", err)
		os.Exit(1)
	}
}
