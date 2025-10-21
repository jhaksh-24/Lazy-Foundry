package forge

import (
	"fmt"
	"os/exec"
	"strings"
)

func Coverage(flags ...string) error {
	args := []string{"coverage"}
	args = append(args, flags...)

	cmd := exec.Command("forge", args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("forge coverage failed: %w\nOutput is: %s", err, string(output))
	}

	fmt.Println("Forge coverage executed Successfully!")

	if len(output) > 0 {
		fmt.Println(strings.TrimSpace(string(output)))
	}

	return nil
}
