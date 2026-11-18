package cmd

import (
	"fmt"

	"memsa/internal/memory"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Inicializa .ai-memory/ e AGENTS.md no repositório atual",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := memory.InitWorkspace(); err != nil {
			return fmt.Errorf("erro ao inicializar workspace: %w", err)
		}
		fmt.Println("Workspace inicializado (.ai-memory/ e AGENTS.md criados).")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}