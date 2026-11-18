package cmd

import (
	"errors"
	"fmt"

	"memsa/internal/memory"
	"memsa/internal/runner"

	"github.com/spf13/cobra"
)

var maxTasks int

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Executa cada tarefa pendente do PLAN.md em contexto isolado",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !memory.WorkspaceExists() {
			return errors.New("workspace não encontrado: execute 'memsa init' primeiro")
		}

		executed := 0
		for {
			if maxTasks > 0 && executed >= maxTasks {
				fmt.Printf("Limite de %d tarefa(s) atingido.\n", maxTasks)
				break
			}

			task, err := memory.NextTask()
			if err != nil {
				return fmt.Errorf("erro ao ler PLAN.md: %w", err)
			}
			if task == "" {
				fmt.Println("Todas as tarefas de .ai-memory/PLAN.md foram concluídas.")
				break
			}

			fmt.Printf("\n=== Executando tarefa (contexto limpo): %s ===\n", task)
			prompt := fmt.Sprintf(
				"Siga AGENTS.md. Leia .ai-memory/PROGRESS.md. Execute APENAS tarefa: '%s'. Valide codigo, atualize PROGRESS.md (max 3 linhas) e encerre.",
				task,
			)

			if err := runner.ExecuteOpenCode(prompt); err != nil {
				return fmt.Errorf("falha na tarefa '%s': %w", task, err)
			}

			if err := memory.CompleteTask(task); err != nil {
				return fmt.Errorf("erro ao atualizar PLAN.md: %w", err)
			}

			executed++
		}
		return nil
	},
}

func init() {
	runCmd.Flags().IntVarP(&maxTasks, "limit", "l", 0, "Número máximo de tarefas para executar neste ciclo (0 = todas)")
	rootCmd.AddCommand(runCmd)
}