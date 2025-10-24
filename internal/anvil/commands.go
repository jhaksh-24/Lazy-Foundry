package anvil

func Execute(command string, args ...string) error {
	switch command {
		case "add":
			return AddPresetCLI(args)
		case "list":
			return ListPresetsCLI()
		case "show":
			return ShowPresetCLI(args)
		case "delete":
			return DeletePresetCLI(args)
		case "start":
			return StartAnvil(args...)
		default:
			return fmt.Errorf("Unknown command: %s", command)
	}
}

func AddPresetCLI(args []string) error {
	if len(args) < 3 {
		return fmt.Errorf("usage: anvil add <name> <rpc-url> <chain-id> [fork-url] [private-key]")
	}

	name := args[0]

	Initializer()

	ImplementRpcURL(args[1])
	ImplementChainID(args[2])

	if len(args) > 3 && args[3] != "" {
		ImplementForkURL(args[3])
	}

	if len(args) > 4 && args[4] != "" {
		if err := ImplementPrivateKey(args[4]); err != nil {
			return fmt.Errorf("invalid private key: %w", err)
		}
	}

	if err := SavePreset(name); err != nil {
		return fmt.Errorf("failed to save preset: %w", err)
	}

	fmt.Println("Preset '%s' created successfully!\n", name)
	return nil
}

func ListPresetsCLI() error {
	names := ListPresets()

	if len(names) == 0 {
		fmt.Println("No presets found.")
		fmt.Println("Create one with: lazyfoundry anvil add <name> <rpc-url> <chain-id>")
		return nil
	}

	fmt.Println("\nAvailable Presets:")
	fmt.Println("─────────────────────────────────────────")

	for _, name := range names {
		preset, err := GetPreset(name)
		if err != nil {
			continue
		}

		fmt.Printf("\n  - %s\n", name)
		fmt.Printf("     RPC: %s\n", preset.RpcURL)
		fmt.Printf("     Chain ID: %d\n", preset.ChainID)
		
		if preset.ForkURL != "" {
			fmt.Printf("     Fork: %s\n", preset.ForkURL)
		}
	}

	fmt.Println("\n─────────────────────────────────────────")
	return nil
}

func DeletePresetCLI(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: anvil delete <preset-name>")
	}

	name := args[0]

	_, err := GetPreset(name)
	if err != nil {
		return err
	}

	err := DeletePreset(name)
	if err != nil {
		return fmt.Errorf("failed to delete preset: %w", err)
	}

	fmt.Printf("Preset '%s' deleted successfully!\n", name)
	return nil
}

func StartAnvil(args ...string) error {
	presetName := "local"
	if len(args) > 0 && args[0] != "" {
		presetName = args[0]
	}

	if err := LoadPreset(presetName); err != nil {
		return fmt.Errorf("failed to load preset '%s': %w", presetName, err)
	}

	fmt.Printf("Starting Anvil with preset '%s'...\n\n", presetName)

	anvilArgs := []string{
		"--chain-id", fmt.Sprintf("%d", anvilConfig.ChainID),
		"--gas-limit", fmt.Sprintf("%d", anvilConfig.GasLimit),
		"--gas-price", fmt.Sprintf("%d", anvilConfig.GasFee),
	}

	if anvilConfig.ForkURL != "" {
		anvilArgs = append(anvilArgs, "--fork-url", anvilConfig.ForkURL)
		fmt.Printf("Forking from: %s\n", anvilConfig.ForkURL)
	}

	fmt.Printf("Chain ID: %d\n", anvilConfig.ChainID)
	fmt.Printf("Gas Limit: %d\n", anvilConfig.GasLimit)
	fmt.Println("\n─────────────────────────────────────────")
	fmt.Println("Running: anvil " + strings.Join(anvilArgs, " "))
	fmt.Println("─────────────────────────────────────────\n")

	cmd := exec.Command("anvil", anvilArgs...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("anvil failed to start: %w", err)
	}

	return nil
}
