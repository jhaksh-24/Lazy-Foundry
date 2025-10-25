package main

import (
	"fmt"
	"os"

	"github.com/jhaksh-24/Lazy-Foundry/internal/anvil"
	"github.com/jhaksh-24/Lazy-Foundry/internal/forge"
)

func main() {
	// Check if user provided any arguments
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	// First argument is the mode (forge or anvil)
	mode := os.Args[1]

	// Route to the correct handler based on mode
	switch mode {
	case "forge":
		handleForge()
	case "anvil":
		handleAnvil()
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Printf("Unknown mode: %s\n\n", mode)
		printUsage()
	}
}

func handleForge() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: lazyfoundry forge <command> [args...]")
		fmt.Println("\nAvailable commands:")
		fmt.Println("  build, test, init, coverage, create, script, install")
		return
	}

	command := os.Args[2]
	args := os.Args[3:]

	var err error

	// Route to the correct forge function
	switch command {
	case "build":
		err = forge.Build(args...)
	case "test":
		err = forge.Test(args...)
	case "init":
		err = forge.Init(args...)
	case "coverage":
		err = forge.Coverage(args...)
	case "install":
		if len(args) < 1 {
			err = fmt.Errorf("install requires package name")
		} else {
			err = forge.Install(args[0], args[1:]...)
		}
	case "script":
		if len(args) < 1 {
			err = fmt.Errorf("script requires script path")
		} else {
			err = forge.Script(args[0], args[1:]...)
		}
	case "create":
		if len(args) < 1 {
			err = fmt.Errorf("create requires contract name")
		} else {
			err = forge.Create(args[0], args[1:]...)
		}
	default:
		err = fmt.Errorf("unknown forge command: %s", command)
	}

	// Handle any errors that occurred
	if err != nil {
		fmt.Printf(" Error: %s\n", err)
		os.Exit(1)
	}
}

// handleAnvil processes all anvil-related commands
func handleAnvil() {
	// Initialize the anvil preset system
	anvil.Initializer()

	if len(os.Args) < 3 {
		fmt.Println("Usage: lazyfoundry anvil <command> [args...]")
		fmt.Println("\nAvailable commands:")
		fmt.Println("  add <n> <rpc-url> <chain-id> [fork-url] [private-key]")
		fmt.Println("  list")
		fmt.Println("  show <n>")
		fmt.Println("  delete <n>")
		fmt.Println("  start [preset-name]")
		return
	}

	command := os.Args[2]
	args := os.Args[3:]

	if err := anvil.Execute(command, args...); err != nil {
		fmt.Printf(" Error: %s\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`
╔═══════════════════════════════════════════════════════════════╗
║                     Lazy-Foundry                              ║
║          Simplified Foundry Workflow Tool                     ║
╚═══════════════════════════════════════════════════════════════╝

USAGE:
  lazyfoundry <mode> <command> [args...]

MODES:
  forge   Build, test, and deploy smart contracts
  anvil   Manage environment presets and run local node

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

FORGE COMMANDS:

  Build & Test:
    lazyfoundry forge build [flags...]
    lazyfoundry forge test [flags...]
    lazyfoundry forge coverage [flags...]
    lazyfoundry forge init [flags...]

  Deploy & Interact:
    lazyfoundry forge create <contract-name> [flags...]
    lazyfoundry forge script <script-path> [flags...]
    
  Dependencies:
    lazyfoundry forge install <package-name> [flags...]

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

ANVIL COMMANDS:

  Preset Management:
    lazyfoundry anvil add <n> <rpc-url> <chain-id> [fork-url] [pk]
    lazyfoundry anvil list
    lazyfoundry anvil show <n>
    lazyfoundry anvil delete <n>

  Start Local Node:
    lazyfoundry anvil start [preset-name]

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

EXAMPLES:

  # Create a local preset
  lazyfoundry anvil add local http://127.0.0.1:8545 31337

  # Start anvil with the local preset
  lazyfoundry anvil start local

  # Initialize a new Foundry project
  lazyfoundry forge init

  # Build your contracts
  lazyfoundry forge build

  # Run tests
  lazyfoundry forge test

  # Deploy a contract
  lazyfoundry forge create MyContract --rpc-url http://localhost:8545

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

For more information, visit:
https://github.com/jhaksh-24/Lazy-Foundry

`)
}
