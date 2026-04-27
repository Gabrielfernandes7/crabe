package inspect

import "github.com/spf13/cobra"

func NewInspectCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "inspect",
		Short: "Analisa hardware e recomenda configurações de IA",
		Run: func(cmd *cobra.Command, args []string) {
			Run()
		},
	}
}