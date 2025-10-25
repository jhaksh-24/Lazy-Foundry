package tui

import (
	"bytes"
	"fmt"
	"os/exec"
	"github.com/jhaksh-24/Lazy-Foundry/internal/anvil"
	"github.com/jhaksh-24/Lazy-Foundry/internal/forge"
)

func executeForge(command string, args ...string) (string, error) {
	var output bytes.Buffer
	var errOutput bytes.Buffer
	
	var err error
	
	switch command {
	case "build":
		err = forge.Build(args...)
	case "test":
		err = forge.Test(args...)
	case "init":
		err = forge.Init(args...)
	case "coverage":
		err = forge.Coverage(args...)
	case "create":
		if len(args) > 0 {
			err = forge.Create(args[0], args[1:]...)
		} else {
			err = fmt.Errorf("contract name required")
		}
	case "script":
		if len(args) > 0 {
			err = forge.Script(args[0], args[1:]...)
		} else {
			err = fmt.Errorf("script path required")
		}
	case "install":
		if len(args) > 0 {
			err = forge.Install(args[0], args[1:]...)
		} else {
			err = fmt.Errorf("package name required")
		}
	default:
		return "", fmt.Errorf("unknown forge command: %s", command)
	}
	
	if err != nil {
		return errOutput.String(), err
	}
	
	return output.String(), nil
}

func executeAnvil(command string, args ...string) (string, error) {
	anvil.Initializer()
	
	switch command {
	case "list":
		names := anvil.ListPresets()
		if len(names) == 0 {
			return "No presets found.\nCreate one with CLI: lazyfoundry anvil add <n> <rpc> <chain-id>", nil
		}
		
		output := "Available Presets:\n\n"
		for _, name := range names {
			preset, _ := anvil.GetPreset(name)
			output += fmt.Sprintf("  • %s\n", name)
			output += fmt.Sprintf("    RPC: %s\n", preset.RpcURL)
			output += fmt.Sprintf("    Chain ID: %d\n\n", preset.ChainID)
		}
		return output, nil
		
	case "show":
		if len(args) < 1 {
			return "", fmt.Errorf("preset name required")
		}
		preset, err := anvil.GetPreset(args[0])
		if err != nil {
			return "", err
		}
		
		output := fmt.Sprintf("Preset: %s\n\n", args[0])
		output += fmt.Sprintf("  RPC URL:    %s\n", preset.RpcURL)
		output += fmt.Sprintf("  Chain ID:   %d\n", preset.ChainID)
		output += fmt.Sprintf("  Gas Limit:  %d\n", preset.GasLimit)
		output += fmt.Sprintf("  Gas Fee:    %d\n", preset.GasFee)
		if preset.ForkURL != "" {
			output += fmt.Sprintf("  Fork URL:   %s\n", preset.ForkURL)
		}
		return output, nil
		
	case "start":
		presetName := "local"
		if len(args) > 0 {
			presetName = args[0]
		}
		return fmt.Sprintf("To start Anvil with preset '%s', run:\n\nlazyfoundry anvil start %s\n\n(Anvil must run in a separate terminal)", presetName, presetName), nil
		
	default:
		return "", fmt.Errorf("unknown anvil command: %s", command)
	}
}
