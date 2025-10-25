package forge

import (
	"fmt"
	"os"
	"os/exec"
)

func Coverage(flags ...string) error {
	args := []string{"coverage"}
	args = append(args, flags...)
	
	cmd := exec.Command("forge", args...)
	
	// Stream output directly to terminal so user sees coverage report
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	// Run the command
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("forge coverage failed: %w", err)
	}
	
	fmt.Println("\n✅ Forge coverage executed successfully!")
	
	return nil
}
