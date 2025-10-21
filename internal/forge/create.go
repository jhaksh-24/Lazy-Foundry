package forge

import (
	"fmt"
	"os/exec"
	"strings"
)

func Create(contractName string, flags ...string) error {
	if contractName == "" {
		return fmt.Errorf("Contract Name has to be given")
	}

	args := []string{"create", contractName}
	args = append(args, flags...)

	cmd := exec.Command("forge", args...)

	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("forge create failed: %w\noutput: %s", err, string(output))
	}

	fmt.Println("Forge create executed Successfully!")

	if len(output) > 0 {
		fmt.Println(strings.TrimSpace(string(output)))
	}

	return nil
}
