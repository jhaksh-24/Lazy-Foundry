package forge

import (
	"fmt"
	"os/exec"
	"strings"
)

func Test(flags ...string) error {
	args := []string{"test"}
	args = append(args, flags...)

	cmd := exec.Command("forge", args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("forge test failed: %w\nOutput: %s", err, string(output))
	}

	fmt.Println("Forge test executed successfully!")

	if len(output) > 0 {
		fmt.Println(strings.TrimSpace(string(output)))
	}

	return nil
} 
