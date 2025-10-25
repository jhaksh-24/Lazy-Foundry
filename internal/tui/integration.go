package tui

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	
	"github.com/jhaksh-24/Lazy-Foundry/internal/anvil"
)

// executeForge runs forge commands and captures their output
// This REPLACES the broken version that couldn't capture output
func executeForge(command string, args ...string) (string, error) {
	// Build the full command arguments
	cmdArgs := []string{command}
	cmdArgs = append(cmdArgs, args...)
	
	// Create the command
	cmd := exec.Command("forge", cmdArgs...)
	
	// Capture both stdout and stderr
	var outputBuf bytes.Buffer
	cmd.Stdout = &outputBuf
	cmd.Stderr = &outputBuf
	
	// Run the command
	err := cmd.Run()
	
	output := outputBuf.String()
	
	// If there's an error, return both the error and any output
	if err != nil {
		if output == "" {
			output = fmt.Sprintf("Command failed: forge %s", strings.Join(cmdArgs, " "))
		}
		return output, err
	}
	
	// If successful but no output, provide a success message
	if output == "" {
		output = fmt.Sprintf("✓ forge %s completed successfully", command)
	}
	
	return output, nil
}

// executeAnvil runs anvil commands and returns their output
func executeAnvil(command string, args ...string) (string, error) {
	// Initialize anvil system
	anvil.Initializer()
	
	switch command {
	case "list":
		return listPresets()
		
	case "show":
		if len(args) < 1 {
			return "", fmt.Errorf("preset name required")
		}
		return showPreset(args[0])
		
	case "start":
		presetName := "local"
		if len(args) > 0 {
			presetName = args[0]
		}
		return startAnvilInfo(presetName)
		
	case "add":
		if len(args) < 3 {
			return "", fmt.Errorf("usage: add <name> <rpc-url> <chain-id> [fork-url] [private-key]")
		}
		return addPreset(args)
		
	case "delete":
		if len(args) < 1 {
			return "", fmt.Errorf("preset name required")
		}
		return deletePreset(args[0])
		
	default:
		return "", fmt.Errorf("unknown anvil command: %s", command)
	}
}

// listPresets formats and returns the list of available presets
func listPresets() (string, error) {
	names := anvil.ListPresets()
	
	if len(names) == 0 {
		return "No presets found.\n\nCreate one using the 'Manage Presets' menu or CLI:\nlazyfoundry anvil add <name> <rpc-url> <chain-id>", nil
	}
	
	var output strings.Builder
	output.WriteString("Available Presets:\n")
	output.WriteString("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n\n")
	
	for _, name := range names {
		preset, err := anvil.GetPreset(name)
		if err != nil {
			continue
		}
		
		output.WriteString(fmt.Sprintf("  📦 %s\n", name))
		output.WriteString(fmt.Sprintf("     RPC: %s\n", preset.RpcURL))
		output.WriteString(fmt.Sprintf("     Chain ID: %d\n", preset.ChainID))
		
		if preset.ForkURL != "" {
			output.WriteString(fmt.Sprintf("     Fork: %s\n", preset.ForkURL))
		}
		output.WriteString("\n")
	}
	
	output.WriteString("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	
	return output.String(), nil
}

// showPreset displays detailed information about a specific preset
func showPreset(name string) (string, error) {
	preset, err := anvil.GetPreset(name)
	if err != nil {
		return "", fmt.Errorf("preset '%s' not found", name)
	}
	
	var output strings.Builder
	output.WriteString(fmt.Sprintf("📦 Preset: %s\n", name))
	output.WriteString("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n\n")
	output.WriteString(fmt.Sprintf("  RPC URL:      %s\n", preset.RpcURL))
	output.WriteString(fmt.Sprintf("  Chain ID:     %d\n", preset.ChainID))
	output.WriteString(fmt.Sprintf("  Gas Limit:    %d\n", preset.GasLimit))
	output.WriteString(fmt.Sprintf("  Gas Fee:      %d\n", preset.GasFee))
	output.WriteString(fmt.Sprintf("  Output Dir:   %s\n", preset.OutputDir))
	
	if preset.ForkURL != "" {
		output.WriteString(fmt.Sprintf("  Fork URL:     %s\n", preset.ForkURL))
	}
	
	if preset.PrivateKey != "" {
		output.WriteString(fmt.Sprintf("  Private Key:  %s... (hidden)\n", preset.PrivateKey[:min(10, len(preset.PrivateKey))]))
	}
	
	output.WriteString("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	
	return output.String(), nil
}

// startAnvilInfo provides info about starting anvil (can't run in TUI directly)
func startAnvilInfo(presetName string) (string, error) {
	// Verify preset exists
	preset, err := anvil.GetPreset(presetName)
	if err != nil {
		return "", fmt.Errorf("preset '%s' not found", presetName)
	}
	
	var output strings.Builder
	output.WriteString(fmt.Sprintf("🚀 Starting Anvil with preset '%s'\n\n", presetName))
	output.WriteString("Configuration:\n")
	output.WriteString(fmt.Sprintf("  Chain ID:  %d\n", preset.ChainID))
	output.WriteString(fmt.Sprintf("  Gas Limit: %d\n", preset.GasLimit))
	
	if preset.ForkURL != "" {
		output.WriteString(fmt.Sprintf("  Forking:   %s\n", preset.ForkURL))
	}
	
	output.WriteString("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n\n")
	output.WriteString("Note: Anvil must run in a separate terminal.\n\n")
	output.WriteString("To start, open a new terminal and run:\n")
	output.WriteString(fmt.Sprintf("  lazyfoundry anvil start %s\n\n", presetName))
	output.WriteString("Or start it in the background:\n")
	output.WriteString(fmt.Sprintf("  lazyfoundry anvil start %s &\n", presetName))
	
	return output.String(), nil
}

// addPreset creates a new preset
func addPreset(args []string) (string, error) {
	if len(args) < 3 {
		return "", fmt.Errorf("usage: add <name> <rpc-url> <chain-id> [fork-url] [private-key]")
	}
	
	name := args[0]
	
	// Use the anvil package's implementation functions
	anvil.ImplementRpcURL(args[1])
	anvil.ImplementChainID(args[2])
	
	if len(args) > 3 && args[3] != "" {
		anvil.ImplementForkURL(args[3])
	}
	
	if len(args) > 4 && args[4] != "" {
		if err := anvil.ImplementPrivateKey(args[4]); err != nil {
			return "", fmt.Errorf("invalid private key: %w", err)
		}
	}
	
	if err := anvil.SavePreset(name); err != nil {
		return "", fmt.Errorf("failed to save preset: %w", err)
	}
	
	return fmt.Sprintf("✓ Preset '%s' created successfully!\n\nUse 'Show Preset' to view details.", name), nil
}

// deletePreset removes a preset
func deletePreset(name string) (string, error) {
	// Check if preset exists
	_, err := anvil.GetPreset(name)
	if err != nil {
		return "", fmt.Errorf("preset '%s' not found", name)
	}
	
	// Delete it
	err = anvil.DeletePreset(name)
	if err != nil {
		return "", fmt.Errorf("failed to delete preset: %w", err)
	}
	
	return fmt.Sprintf("✓ Preset '%s' deleted successfully!", name), nil
}

// Helper function for min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// captureCommandOutput is a generic helper for running any command
// and capturing its output (kept for potential future use)
func captureCommandOutput(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	
	var outputBuf bytes.Buffer
	cmd.Stdout = &outputBuf
	cmd.Stderr = &outputBuf
	
	err := cmd.Run()
	
	return outputBuf.String(), err
}
