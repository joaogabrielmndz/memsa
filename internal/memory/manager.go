package memory

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"memsa/internal/templates"
)

const (
	MemoryDir    = ".ai-memory"
	PlanFile     = ".ai-memory/PLAN.md"
	ProgressFile = ".ai-memory/PROGRESS.md"
	AgentsFile   = "AGENTS.md"
)

func InitWorkspace() error {
	if err := os.MkdirAll(MemoryDir, 0755); err != nil {
		return err
	}

	files := map[string]string{
		AgentsFile:   templates.AgentsMD,
		PlanFile:     templates.PlanMD,
		ProgressFile: "# Log de Progresso\n",
	}

	for path, content := range files {
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			if err := os.WriteFile(path, []byte(content), 0644); err != nil {
				return err
			}
		}
	}
	return nil
}

func NextTask() (string, error) {
	data, err := os.ReadFile(PlanFile)
	if err != nil {
		return "", err
	}

	for _, line := range strings.Split(string(data), "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "- [ ] ") {
			return strings.TrimPrefix(trimmed, "- [ ] "), nil
		}
	}
	return "", nil
}

func CompleteTask(task string) error {
	data, err := os.ReadFile(PlanFile)
	if err != nil {
		return err
	}

	oldLine := "- [ ] " + task
	newLine := "- [x] " + task
	updated := strings.Replace(string(data), oldLine, newLine, 1)

	return os.WriteFile(PlanFile, []byte(updated), 0644)
}

func WorkspaceExists() bool {
	_, err := os.Stat(filepath.Clean(PlanFile))
	return err == nil
}