package anvil

import (
	"fmt"
)

type AnvilConfig struct {
	RpcURL string
	ChainID int
	ForkURL string
	PrivateKey string
	GasLimit int
	GasPrice int
	OutputDir string
}

var anvilConfig AnvilConfig

func Initializer() {
	anvilConfig = AnvilConfig{
		RpcURL:    DefaultRPCURL,
		ChainID:   DefaultChainID,
		GasLimit:  DefaultGasLimit,
		GasPrice:  DefaultGasPrice,
		OutputDir: ConfigDirName,
	}
}
