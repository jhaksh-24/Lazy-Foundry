package main

import (
//	"fmt"
//	"os"
//	"github.com/jhaksh-24/Lazy-Foundry/internal/forge"
	"github.com/jhaksh-24/Lazy-Foundry/internal/anvil"
)

func main() {
/*	if len(os.Args) < 2 {
		fmt.Println("Command not provided")
		return
	}

	command := os.Args[1]
	err := forge.Execute(command)

	if err != nil {
		fmt.Println("Error:", err)
	}*/

	anvil.StartAnvilUI()
}
