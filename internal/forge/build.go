package forge

import (
	"fmt"
	"os/exec"
	"strings"
)

func Build(flags ...string) error {
	args := []string{"build"}
	
	args = append(args, flags...)
	
	cmd := exec.Command("forge", args...)
	
	output, err := cmd.CombinedOutput()
	
	if err != nil {
		return fmt.Errorf("forge build failed: %w\nOutput: %s", err, string(output))
	}
	
	fmt.Println("Forge build executed successfully!")
	
	if len(output) > 0 {
		fmt.Println(strings.TrimSpace(string(output)))
	}
	
	return nil
}
