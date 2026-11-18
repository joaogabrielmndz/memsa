package runner

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecuteOpenCode(prompt string) error {
	cmd := exec.Command("opencode", "run", prompt)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("falha ao executar opencode: %w", err)
	}
	return nil
}