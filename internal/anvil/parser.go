package anvil

import (
	"fmt"
	"os"
)

func Parser() error {
	if len(os.Args) > 1 {
		ImplementRpcURL()
		ImplementChainId()
		ImplementForkURL()
		err = ImplementPrivateKey()
		if err != nil {
			return err
		}
		// Private Keys will be encrypted when we will be dealing withCast
		ImplementGasLimit()
		ImplementGasPrice()
		anvilConfig.OutputDir = os.Args[7]
		return nil
	}
}
