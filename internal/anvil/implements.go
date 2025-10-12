package anvil

import (
	"fmt"
	"os"
)

func ImplementRpcURL() {
	if isRpcURL(os.Args[1]) {
		anvilConfig.RpcURL= os.Args[1]
	} 
}

func ImplementChainID() {
	if isChainID(os.Args[2]){
		anvilConfig.ChainID= os.Args[2]
	} 
}

func ImplementGasLimit() {
	if isGasLimit(os.Args[5]) {
		anvilConfig.GasLimit= os.Args[5]
	}
}

func ImplementGasFee() {
	if isGasFee(os.Args[6]) {
		anvilConfig.GasFee= os.Args[6]
	}
}

func ImplementForkURL() {
	if isForkURL(os.Args[3]) {
		anvilConfig.ForkURL=os.Args[3]
	}
}
func ImplementPrivateKey() error {
	if isPrivateKey(os.Args[4]) {
		anvilConfig.CheckPrivateKey = os.Args[4]
		return nil
	}
	return fmt.Errorf("Private key was not entered")
}
