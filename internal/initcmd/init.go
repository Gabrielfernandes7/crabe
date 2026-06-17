package initcmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Gabrielfernandes7/crabe/internal/ui"
	"github.com/spf13/cobra"
)

const (
	contextDir  = ".crabe"
	contextFile = "CRABE.md"
	configFile  = "config.yaml"
)

var subDirs = []string{"skills", "outputs", "memory"}

func NewInitCmd() *cobra.Command {
	var force bool

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Inicializa o Crabe Business Agent no projeto atual",
		Long:  "Cria a estrutura do Business Agent (.crabe/) e prepara o ambiente local.",
		Run: func(cmd *cobra.Command, args []string) {
			RunInit(force)
		},
	}

	cmd.Flags().BoolVarP(&force, "force", "f", false, "Força reinicialização do workspace")

	return cmd
}

func RunInit(force bool) {
	ui.Title("🚀 Crabe Workspace Init")

	// 1. Detectar se já existe contexto
	if contextExists() && !force {
		ui.Warning("Workspace já inicializado (.crabe encontrado)")
		ui.Info("Use --force para recriar")
		return
	}

	// 2. Criar estrutura local
	ui.Section("Estrutura do Workspace")
	if err := createStructure(force); err != nil {
		ui.Error(fmt.Sprintf("Erro ao criar estrutura: %v", err))
		return
	}

	// 3. Verificar ambiente
	ui.Section("Ambiente")
	// Por agora, apenas informa que o ambiente deve ser verificado com doctor
	ui.Info("Verifique seu ambiente local com: crabe doctor")
	ui.Success("Estrutura de pastas pronta")

	// 4. Mensagem final
	ui.Title("✅ Workspace Pronto")

	ui.Success("Crabe Business Agent inicializado com sucesso!")

	fmt.Println()
	ui.Info("Próximos passos:")
	ui.Info("  crabe trust      → autorizar escrita no workspace")
	ui.Info("  crabe            → iniciar agente interativo")
	fmt.Println()
}

func contextExists() bool {
	_, err := os.Stat(contextDir)
	return err == nil
}

func createStructure(force bool) error {
	if force {
		_ = os.RemoveAll(contextDir)
	}

	// Criar diretório principal
	if err := os.MkdirAll(contextDir, 0755); err != nil {
		return err
	}

	// Criar subdiretórios
	for _, dir := range subDirs {
		path := filepath.Join(contextDir, dir)
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
		ui.Success("Criado: %s/", path)
	}

	// Criar CRABE.md
	crabePath := filepath.Join(contextDir, contextFile)
	if err := os.WriteFile(crabePath, []byte(defaultCRABE()), 0644); err != nil {
		return err
	}
	ui.Success("Criado: %s", crabePath)

	// Criar config.yaml
	configPath := filepath.Join(contextDir, configFile)
	if err := os.WriteFile(configPath, []byte(defaultConfig()), 0644); err != nil {
		return err
	}
	ui.Success("Criado: %s", configPath)

	return nil
}

func defaultCRABE() string {
	return `# 🦀 Business Context: {{Project Name}}

Este arquivo define o contexto do negócio para o Crabe Agent.

## Visão Geral
- **Nome do Projeto:** 
- **Setor:** 
- **Público-Alvo:** 

## Objetivos de Negócio
1. 
2. 

## Diferenciais
- 

## Desafios Atuais
- 
`
}

func defaultConfig() string {
	return `model: qwen3:8b
ollama_url: http://localhost:11434
temperature: 0.7
max_tokens: 4096
`
}
