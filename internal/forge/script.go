package forge

import (
	"fmt"
	"os/exec"
	"strings"
)

func Script(scriptPath string, flags ...string) error {
	if scriptPath == "" {
		return fmt.Errorf("Script Path is required")
	}

	args := []string{"script", scriptPath}
	args = append(args, flags...)

	cmd := exec.Command("forge", args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("forge script failed: %w\nOutput: %s", err, string(output))
	}

	fmt.Println("Forge script executed Successfully!")

	if len(output) > 0 {
		fmt.Println(strings.TrimSpace(string(output)))
	}
	return nil
}
