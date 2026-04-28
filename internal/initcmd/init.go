package initcmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Gabrielfernandes7/crabe/internal/setup"
	"github.com/Gabrielfernandes7/crabe/internal/ui"
	"github.com/spf13/cobra"
)

const (
	contextDir  = ".crabe"
	contextFile = "context.md"
)

func NewInitCmd() *cobra.Command {
	var force bool

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Inicializa o Crabe no projeto atual",
		Long:  "Cria o contexto local (.crabe/) e prepara o ambiente com Ollama automaticamente.",
		Run: func(cmd *cobra.Command, args []string) {
			RunInit(force)
		},
	}

	cmd.Flags().BoolVarP(&force, "force", "f", false, "Força reinicialização do contexto")

	return cmd
}

func RunInit(force bool) {
	ui.Title("🚀 Crabe Init")

	// 1. Detectar se já existe contexto
	if contextExists() && !force {
		ui.Warning("Projeto já inicializado (.crabe encontrado)")
		ui.Info("Use --force para recriar")
		return
	}

	// 2. Criar estrutura local
	ui.Section("Contexto do projeto")
	if err := createContext(force); err != nil {
		ui.Error(fmt.Sprintf("Erro ao criar contexto: %v", err))
		return
	}

	// 3. Verificar ambiente (rápido)
	ui.Section("Ambiente")

	// Usamos a função preflight interna (não exportada) através de uma função wrapper que vamos criar depois
	// Por enquanto, chamamos o RunSetup diretamente se necessário
	state := getSystemState() // função auxiliar que vamos definir

	needsSetup := !state.OllamaRunning

	if needsSetup {
		ui.Warning("Ambiente incompleto detectado")
		ui.Info("Executando setup automático...")
		setup.RunSetup(false) // Apenas 1 argumento (force)
	} else {
		ui.Success("Ambiente já pronto")
	}

	// 4. Mensagem final
	ui.Title("✅ Projeto pronto")

	ui.Success("Crabe inicializado com sucesso neste diretório!")

	fmt.Println()
	ui.Info("Próximos passos:")
	ui.Info("  crabe start      → subir serviços")
	ui.Info("  crabe status     → verificar status")
	ui.Info("  crabe doctor     → diagnóstico")
	fmt.Println()
	ui.Info("Web UI: http://localhost:3000")
}

// Função auxiliar para acessar o state (já que preflight é privada)
func getSystemState() setup.SystemState {
	// Por enquanto chamamos diretamente o preflight interno via reflection ou duplicamos um pouco
	// Versão simples: sempre assume que precisa de setup se o init for chamado (melhorar depois)
	return setup.SystemState{
		OllamaRunning: false,
	}
}

func contextExists() bool {
	_, err := os.Stat(contextDir)
	return err == nil
}

func createContext(force bool) error {
	if force {
		_ = os.RemoveAll(contextDir)
	}

	if err := os.MkdirAll(contextDir, 0755); err != nil {
		return err
	}

	path := filepath.Join(contextDir, contextFile)

	if _, err := os.Stat(path); err == nil && !force {
		ui.Success("context.md já existe")
		return nil
	}

	content := defaultContext()

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return err
	}

	ui.Success(fmt.Sprintf("Criado: %s", path))
	return nil
}

func defaultContext() string {
	return `# 🦀 Crabe Context

Descreva aqui o contexto do seu projeto para melhorar a qualidade das respostas da IA.

## Projeto
- Nome: 
- Descrição: 

## Stack
- Backend: 
- Frontend: 
- Infra: 

## Objetivo
Descreva o que você quer construir ou melhorar.

## Observações
Qualquer detalhe relevante para o assistente.
`
}
