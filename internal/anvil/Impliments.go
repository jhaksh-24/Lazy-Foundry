package anvil

import (
	"fmt"
	"os"
)

func ImplementRpcURL() {
	if isRpcURL(os.Args[1]) {
		anvilConfig.RpcURL= os.Args[1]
	} else{
		anvilConfig.RpcURL=constants.RpcURL
	}
}

func ImplementChainId() {
	if isChainId(os.Args[2]){
		anvilConfig.ChainId= os.Args[2]
	} else {
		anvilConfig.ChainId=constants.ChainId
	}
}

func ImplementGasLimit() {
	if isGasLimit(os.Args[5]) {
		anvilConfig.GasLimit= os.Args[5]
	} else {
		anvilConfig.GasLimit=constants.GasLimit
	}
}

func ImplementGasPrice() {
	if isGasLimit(os.Args[6]) {
		anvilConfig.GasPrice= os.Args[6]
	} else {
		anvilConfig.GasPrice=constants.GasPrice
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
