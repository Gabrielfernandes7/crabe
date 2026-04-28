package doctor

import (
	"github.com/spf13/cobra"
)

func NewDoctorCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "doctor",
		Short: "Executa diagnóstico do sistema (Docker, Ollama, portas, etc.)",
		Long:  `Verifica se tudo necessário para rodar o Crabe está funcionando corretamente.`,
		Run: func(cmd *cobra.Command, args []string) {
			Run()
		},
	}
}
