package anvil

import (
	"fmt"
)

func ImplementRpcURL(s string) {
	if s == "" {
        return
    }

	if isRpcURL(s) {
		anvilConfig.RpcURL = s
	}
}

func ImplementChainID(s string) {
	if s == "" {
        return
    }

	checkPass, id := isChainID(s)
	if checkPass {
		 anvilConfig.ChainID = id
	} 
}

func ImplementGasLimit(s string) {
	if s == "" {
		return
	}

	checkPass, gl := isGasLimit(s)
	if checkPass {
		anvilConfig.GasLimit = gl
	}
}

func ImplementGasFee(s string) {
	if s == "" {
        return
    }

	checkPass, gf := isGasFee(s)
	if checkPass {
		anvilConfig.GasFee= gf
	}
}

func ImplementForkURL(s string) {
	if s == "" {
        return
    }

	if isForkURL(s) {
		anvilConfig.ForkURL = s
	}
}

func ImplementPrivateKey(s string) error {
	if isPrivateKey(s) {
		anvilConfig.PrivateKey = s
		return nil
	}
	return fmt.Errorf("Private key was not entered")
}
