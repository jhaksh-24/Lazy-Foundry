package main

import (
	"fmt"
	"os"

	"yourmodule/internal/anvil"
)

func main() {
	// 1️⃣ Show defaults before parsing
	anvil.Initializer()
	fmt.Println("=== Defaults before parsing ===")
	fmt.Printf("%+v\n\n", anvil.AnvilConfigInstance)

	// 2️⃣ Simulate command-line args if not running via CLI
	if len(os.Args) < 8 {
		os.Args = []string{
			"main",                                 // placeholder
			"http://127.0.0.1:8545",               // RpcURL
			"31337",                                // ChainID
			"https://somefork.com",                 // ForkURL
			"0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890", // PrivateKey
			"30000000",                             // GasLimit
			"1000000000",                           // GasPrice
			"./output",                             // OutputDir
		}
	}

	// 3️⃣ Parse user input & override defaults
	err := anvil.Parser()
	if err != nil {
		fmt.Println("Parser error:", err)
		return
	}

	// 4️⃣ Show config after parsing
	fmt.Println("=== Config after parsing ===")
	fmt.Printf("%+v\n", anvil.AnvilConfigInstance)
}

