package uninstall

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/Gabrielfernandes7/crabe/internal/ui"
	"github.com/spf13/cobra"
)

func NewUninstallCmd() *cobra.Command {
	var yes bool

	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Remove completamente o Crabe do sistema (MacOS, Linux, Windows)",
		Run: func(cmd *cobra.Command, args []string) {
			runUninstall(yes)
		},
	}

	cmd.Flags().BoolVarP(&yes, "yes", "y", false, "Confirmação automática")

	return cmd
}

func runUninstall(yes bool) {
	ui.Title("🗑️  Crabe Global Uninstall")

	if !yes {
		ui.Warning("⚠️  Atenção: Esta ação removerá as configurações globais do Crabe.")
		ui.Info("Isso inclui arquivos em ~/.crabe ou AppData/Roaming/crabe.")
		fmt.Print("\nDeseja continuar? (s/N): ")
		var confirm string
		fmt.Scanln(&confirm)
		if confirm != "s" && confirm != "S" {
			ui.Info("Operação cancelada.")
			return
		}
	}

	ui.Section("Removendo arquivos de configuração")
	cleanGlobalConfig()

	ui.Section("Instruções de remoção do binário")
	showBinaryInfo()

	ui.Title("✅ Processo concluído")
	ui.Success("Configurações globais removidas com sucesso.")
}

func cleanGlobalConfig() {
	home, _ := os.UserHomeDir()
	var configPath string

	switch runtime.GOOS {
	case "windows":
		appData := os.Getenv("APPDATA")
		if appData == "" {
			appData = filepath.Join(home, "AppData", "Roaming")
		}
		configPath = filepath.Join(appData, "crabe")
	default:
		// MacOS e Linux
		configPath = filepath.Join(home, ".crabe")
	}

	if _, err := os.Stat(configPath); err == nil {
		if err := os.RemoveAll(configPath); err == nil {
			ui.Success("Removido: %s", configPath)
		} else {
			ui.Error("Erro ao remover %s: %v", configPath, err)
		}
	} else {
		ui.Info("Diretório de configuração não encontrado: %s", configPath)
	}
}

func showBinaryInfo() {
	exePath, err := os.Executable()
	if err != nil {
		ui.Warning("Não foi possível localizar o binário atual.")
		return
	}

	ui.Info("O binário atual está em: %s", exePath)
	
	switch runtime.GOOS {
	case "windows":
		ui.Info("Para remover completamente, apague o arquivo .exe acima.")
	case "darwin", "linux":
		ui.Info("Para remover o binário, você pode rodar:")
		ui.Highlight("  sudo rm %s", exePath)
	}
}
