package anvil

import (
	"os"
)
func CheckRpcURL (){
	if os.Args[1]==isRpcURL() {
		anvilConfig.RpcURL= os.Args[1]
	} else{
		anvilConfig.RpcURL=constants.RpcURL
	}
}
func CheckChainId (){
	if os.Args[2]==isChainId(){
		anvilConfig.ChainId= os.Args[2]
	} else {
		anvilConfig.ChainId=constants.ChainId
	}
}
func CheckGasLimit(){
	if os.Args[5]==isGasLimit(){
		anvilConfig.CheckGasLimit= os.Args[5]
	} else {
		anvilConfig.GasLimit=constants.GasLimit
	}
}
func CheckGasPrice(){
	if os.Args[6]==isGasLimit(){
		anvilConfig.CheckGasPrice= os.Args[6]
	} else {
		anvilConfig.GasPrice=constants.GasPrice
	}
}
func CheckForkURLError(){
	if os.Args[3]==isForkURL(){
		anvilConfig.CheckForkURLError=os.Args[3]
	}
}
func CheckPrivateKey() error {
	if os.Args[4]==isPrivateKey(){
		anvilConfig.CheckPrivateKey=constants.PrivateKey
		return nil
	}
}
