package forge

import (
	"fmt"
	"os"
	"os/exec"
)

func Create(contractName string, flags ...string) error {
	if contractName == "" {
		return fmt.Errorf("contract name is required")
	}
	
	args := []string{"create", contractName}
	args = append(args, flags...)
	
	cmd := exec.Command("forge", args...)
	
	// Stream output so user sees deployment details and contract address
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("forge create failed: %w", err)
	}
	
	fmt.Println("\n✅ Contract deployed successfully!")
	
	return nil
}
