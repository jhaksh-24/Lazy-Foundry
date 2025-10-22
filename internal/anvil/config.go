package anvil

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type AnvilConfig struct {
	RpcURL     string `json:"rpc_url"` 
	ChainID    int    `json:"chain_id"`
	ForkURL    string `json:"fork_url,omitempty"`
	PrivateKey string `json:"private_key, omitempty"`
	GasLimit   uint64 `json:"gas_limit"`
	GasFee     uint64 `json:"gas_fee"`
	OutputDir  string `json:"output_dir"`
}

type ConfigStore struct {
	Presets map[string] AnvilConfig `json:"presets"`
}

var (
	anvilConfig AnvilConfig
	configStrore ConfigStore
	configPath string
)

func Initializer() {
	anvilConfig = AnvilConfig{
		RpcURL:    DefaultRPCURL,
		ChainID:   DefaultChainID,
		GasLimit:  DefaultGasLimit,
		GasFee:    DefaultGasFee,
		OutputDir: ConfigDirName,
	}


	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Warning: Could not get home directory")
		return
	}

	configDir := filepath.Join(home, ConfigDirName)
	configPath = filepath.Join(configDir,ConfigFileName)

	os.MkdirAll(configDir, 0755)

	configStore.Presets = make(map[string]AnvilConfig)

	err := LoadAllPresets()
	if err != nil {
		CreateDefaultPresets()
	}
}

func CreateDefaultPresets() {
	configStore.Presets["local"] = AnvilConfig{
		RpcURL:    DefaultRPCURL,
		ChainID:   DefaultChainID,
		GasLimit:  DefaultGasLimit,
		GasFee:    DefaultGasFee,
		OutputDir: ConfigDirName,
	}

	configStore.Presets["sepolia"] = AnvilConfig{
		RpcURL:    "https://sepolia.infura.io/v3/SOME_KEY",
		ChainID:   11155111,
		GasLimit:  DefaultGasLimit,
		GasFee:    DefaultGasFee,
		OutputDir: ConfigDirName,
	}

	SaveAllPresets()
}

func SavePreset(name string) error {
	if name == "" {
		return fmt.Errorf("preset has to be given a name")
	}

	configStore.Presets[name] = anvilConfig
	return SaveAllPresets()
}

func LoadPreset(nameString) error {
	preset, exists := configStore.Presets[name]
	if !exists {
		return fmt.Errorf("failed to fetch preset %s", name)
	}

	anvilConfig = preset
	return nil
}
