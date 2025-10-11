package anvil

import (
	"fmt"
)

func mainAnvil() {
	Initializer()
	err := Parser()
	if err != nil {
		fmt.Println("Error:" ,err)
	}
}
