package forge

import (
	"fmt"
	"os/exec"
	"strings"
)

func Install(packageName string, flags ...string) error {
	if packageName == "" {
		return fmt.Errorf("package name is required!")
	}

	args := []string{"install", packageName}
	args = append(args, flags...)

	cmd := exec.Command("forge", args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("forge install failed: %w\nOutput: %s", err, string(output))
	}

	fmt.Println("Forge install executed Successfully!")

	if len(output) > 0 {
		fmt.Println(strings.TrimSpace(string(output)))
	}

	return nil
}
