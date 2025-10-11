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

func Initializer() {
	anvilConfig AnvilConfig = new AnvilConfig()
}
