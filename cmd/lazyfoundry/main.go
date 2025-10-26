package main

import (
	"fmt"
	"os"

	"github.com/jhaksh-24/Lazy-Foundry/internal/anvil"
	"github.com/jhaksh-24/Lazy-Foundry/internal/forge"
	"github.com/jhaksh-24/Lazy-Foundry/internal/server"
)

func main() {
	// If no arguments, launch web UI
	if len(os.Args) < 2 {
		fmt.Println("🌐 Starting Lazy-Foundry Web Interface...")
		s := server.New()
		if err := s.Start("3000"); err != nil {
			fmt.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
		return
	}

	// First argument is the mode
	mode := os.Args[1]

	// Route to the correct handler based on mode
	switch mode {
	case "web":
		// Explicit web server launch
		port := "3000"
		if len(os.Args) > 2 {
			port = os.Args[2]
		}
		fmt.Printf("🌐 Starting Lazy-Foundry Web Interface on port %s...\n", port)
		s := server.New()
		if err := s.Start(port); err != nil {
			fmt.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	case "tui":
		// Explicit TUI launch
		if err := tui.Run(); err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
	case "forge":
		handleForge()
	case "anvil":
		handleAnvil()
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Printf("❌ Unknown mode: %s\n\n", mode)
		printUsage()
	}
}

// handleForge processes all forge-related commands
func handleForge() {
	if len(os.Args) < 3 {
		fmt.Println("❌ Usage: lazyfoundry forge <command> [args...]")
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
		fmt.Printf("❌ Error: %s\n", err)
		os.Exit(1)
	}
}

// handleAnvil processes all anvil-related commands
func handleAnvil() {
	// Initialize the anvil preset system
	anvil.Initializer()

	if len(os.Args) < 3 {
		fmt.Println("❌ Usage: lazyfoundry anvil <command> [args...]")
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

	// Execute the anvil command
	if err := anvil.Execute(command, args...); err != nil {
		fmt.Printf("❌ Error: %s\n", err)
		os.Exit(1)
	}
}

// printUsage displays help information
func printUsage() {
	fmt.Println(`
╔═══════════════════════════════════════════════════════════════╗
║                     🧰 Lazy-Foundry                           ║
║          Simplified Foundry Workflow Tool                     ║
╚═══════════════════════════════════════════════════════════════╝

USAGE:
  lazyfoundry                    Launch Web Interface (default)
  lazyfoundry web [port]         Launch Web Interface on specified port
  lazyfoundry tui                Launch Terminal User Interface
  lazyfoundry <mode> <command>   Run CLI command

MODES:
  web     Launch web-based graphical interface
  tui     Launch terminal-based interface
  forge   Build, test, and deploy smart contracts
  anvil   Manage environment presets and run local node

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

WEB INTERFACE:

  Start Web UI:
    lazyfoundry                     (default port 3000)
    lazyfoundry web                 (default port 3000)
    lazyfoundry web 8080            (custom port)

  Access at: http://localhost:3000

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

  # Launch Web Interface (recommended)
  lazyfoundry

  # Launch Web Interface on custom port
  lazyfoundry web 8080

  # Launch Terminal UI
  lazyfoundry tui

  # Create a local preset (CLI)
  lazyfoundry anvil add local http://127.0.0.1:8545 31337

  # Build your contracts (CLI)
  lazyfoundry forge build

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

For more information, visit:
https://github.com/jhaksh-24/Lazy-Foundry

`)
}
