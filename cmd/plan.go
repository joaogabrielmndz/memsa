package cmd

import (
	"errors"
	"fmt"
	"strings"

	"memsa/internal/memory"
	"memsa/internal/runner"

	"github.com/spf13/cobra"
)

var planCmd = &cobra.Command{
	Use:   "plan [objetivo]",
	Short: "Gera tarefas atômicas no .ai-memory/PLAN.md",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if !memory.WorkspaceExists() {
			return errors.New("workspace não encontrado: execute 'memsa init' primeiro")
		}

		goal := strings.Join(args, " ")
		prompt := fmt.Sprintf(
			"OBJETIVO: %s. Leia estrutura projeto. Escreva tarefas atomicas checklist '- [ ] tarefa' em .ai-memory/PLAN.md. Estilo Caveman. Sem explicar.",
			goal,
		)

		return runner.ExecuteOpenCode(prompt)
	},
}

func init() {
	rootCmd.AddCommand(planCmd)
}