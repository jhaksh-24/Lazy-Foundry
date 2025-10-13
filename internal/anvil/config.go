package anvil

type AnvilConfig struct {
	RpcURL string
	ChainID int
	ForkURL string
	PrivateKey string
	GasLimit uint64
	GasFee uint64
	OutputDir string
}

var anvilConfig AnvilConfig

func Initializer() {
	anvilConfig = AnvilConfig{
		RpcURL:    DefaultRPCURL,
		ChainID:   DefaultChainID,
		GasLimit:  DefaultGasLimit,
		GasFee:  DefaultGasFee,
		OutputDir: ConfigDirName,
	}
}
