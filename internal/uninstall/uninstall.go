package uninstall

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Gabrielfernandes7/crabe/internal/ui"
	"github.com/spf13/cobra"
)

func NewUninstallCmd() *cobra.Command {
	var all bool
	var yes bool

	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Remove completamente Ollama, Docker e configurações do Crabe",
		Run: func(cmd *cobra.Command, args []string) {
			runUninstall(all, yes)
		},
	}

	cmd.Flags().BoolVarP(&all, "all", "a", false, "Remove também Ollama, volumes Docker e todos os dados")
	cmd.Flags().BoolVarP(&yes, "yes", "y", false, "Não pede confirmação (modo não-interativo)")

	return cmd
}

func runUninstall(all, yes bool) {
	ui.Title("🗑️  Crabe Uninstall")

	if !yes {
		ui.Warning("⚠️  Esta ação é irreversível!")
		ui.Info("Vai limpar configurações locais do Crabe e, se --all, remover Ollama e volumes Docker.")
		ui.Info("Execute com --yes para confirmar.")
		return
	}

	// 1. Ollama + Docker (se --all)
	if all {
		ui.Section("Removendo Ollama + Docker")
		removeOllama()
	} else {
		ui.Success("Ollama mantido (use --all para remover também)")
	}

	// 2. Limpeza final do Crabe
	ui.Section("Limpando configurações do Crabe")
	cleanCrabeDirectories()

	ui.Title("✅ Desinstalação concluída!")
	ui.Success("Limpeza do Crabe finalizada.")
}

func removeOllama() {
	ui.Info("Parando e removendo Ollama...")
	_ = exec.Command("docker", "compose", "down", "--volumes", "--remove-orphans").Run()
	_ = exec.Command("docker", "rm", "-f", "ollama").Run()
	_ = exec.Command("docker", "volume", "prune", "-f").Run()
	ui.Success("Ollama removido")
}

func cleanCrabeDirectories() {
	home, _ := os.UserHomeDir()
	paths := []string{
		filepath.Join(home, ".crabe"),
	}

	for _, p := range paths {
		if err := os.RemoveAll(p); err == nil {
			ui.Success(fmt.Sprintf("Removido: %s", p))
		} else if !os.IsNotExist(err) {
			ui.Warning(fmt.Sprintf("Não conseguiu remover: %s", p))
		}
	}
}
