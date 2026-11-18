package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "memsa",
	Short:   "Orquestrador multi-agente com memória em disco e compressão de tokens",
	Long:    "Gerencia ciclos isolados do OpenCode usando arquivos em .ai-memory/ e protocolo Caveman.",
	Version: "0.1.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}