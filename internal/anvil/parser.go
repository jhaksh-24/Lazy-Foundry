package anvil

import (
	"fmt"
	"os"
)

func Parser() error {
	if len(os.Args) > 1 {
		CheckRpcURL()
		CheckChainId()
		anvilConfig.ForkURL := os.Args[3]
		anvilConfig.PrivateKey := os.Args[4]
		// Private Keys will be encrypted when we will be dealing withCast
		anvilConfig.GasLimit = os.Args[5]
		anvilConfig.GasPrice = os.Args[6]
		anvilConfig.OutputDir = os.Args[7]
		return nil
	}
}
