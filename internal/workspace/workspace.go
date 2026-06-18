package workspace

import (
	"fmt"
	"os"
	"strings"

	"github.com/Gabrielfernandes7/crabe/internal/ui"
	"github.com/spf13/cobra"
)

type TrustLevel string

const (
	Untrusted      TrustLevel = "untrusted"
	ReadOnly       TrustLevel = "read-only"
	WorkspaceWrite TrustLevel = "workspace-write"
)

const trustFileName = ".crabe/trust"

func NewTrustCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "trust",
		Short: "Configura o nível de confiança do workspace atual",
		Run: func(cmd *cobra.Command, args []string) {
			RunTrustPrompt()
		},
	}
}

func RunTrustPrompt() {
	ui.Title("Workspace Trust")
	
	wd, _ := os.Getwd()
	fmt.Printf("Workspace:\n%s\n\n", wd)
	fmt.Println("Você confia nesta pasta?")
	fmt.Println("[1] Sim (Escrita permitida)")
	fmt.Println("[2] Somente leitura")
	fmt.Println("[3] Cancelar")
	fmt.Print("\nEscolha: ")

	var choice string
	fmt.Scanln(&choice)

	var level TrustLevel
	switch choice {
	case "1":
		level = WorkspaceWrite
	case "2":
		level = ReadOnly
	default:
		ui.Info("Operação cancelada")
		return
	}

	if err := SetTrustLevel(level); err != nil {
		ui.Error("Erro ao salvar nível de confiança: %v", err)
		return
	}

	ui.Success("Nível de confiança definido para: %s", level)
}

func SetTrustLevel(level TrustLevel) error {
	_ = os.MkdirAll(".crabe", 0755)
	return os.WriteFile(trustFileName, []byte(level), 0644)
}

func GetTrustLevel() TrustLevel {
	data, err := os.ReadFile(trustFileName)
	if err != nil {
		return Untrusted
	}
	return TrustLevel(data)
}

func IsTrusted() bool {
	level := GetTrustLevel()
	return level == WorkspaceWrite || level == ReadOnly
}

func CanWrite() bool {
	return GetTrustLevel() == WorkspaceWrite
}

func LoadConfig() (string, error) {
	data, err := os.ReadFile(".crabe/config.yaml")
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "model: ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "model: ")), nil
		}
	}
	return "", fmt.Errorf("model not found in config")
}
