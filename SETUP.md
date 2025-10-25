# 🚀 Lazy-Foundry Setup Guide

## 📋 Prerequisites

1. **Go** (version 1.21 or higher)
   ```bash
   go version
   ```

2. **Foundry** (forge, anvil, cast)
   ```bash
   curl -L https://foundry.paradigm.xyz | bash
   foundryup
   ```

## 🔧 Installation

### Step 1: Clone the Repository
```bash
git clone https://github.com/jhaksh-24/Lazy-Foundry.git
cd Lazy-Foundry
```

### Step 2: Install Dependencies
```bash
go mod download
```

### Step 3: Build the Project
```bash
go build -o lazyfoundry cmd/lazyfoundry/main.go
```

### Step 4: (Optional) Install Globally
```bash
# Move to a directory in your PATH
sudo mv lazyfoundry /usr/local/bin/

# Or add to your shell config
echo 'export PATH=$PATH:~/Lazy-Foundry' >> ~/.bashrc
source ~/.bashrc
```

## ✅ Verify Installation

```bash
# Should show the help menu
./lazyfoundry help

# Launch TUI
./lazyfoundry
```

## 🎨 TUI Features

The TUI (Terminal User Interface) provides:
- **Visual Navigation** - Use arrow keys to navigate menus
- **Forge Mode** - Build, test, and deploy contracts
- **Anvil Mode** - Manage presets and local nodes
- **Beautiful UI** - Color-coded, intuitive interface

## 🖥️ CLI Mode

You can also use CLI commands directly:

```bash
# Forge commands
lazyfoundry forge build
lazyfoundry forge test
lazyfoundry forge create MyContract --rpc-url http://localhost:8545

# Anvil commands
lazyfoundry anvil add local http://127.0.0.1:8545 31337
lazyfoundry anvil list
lazyfoundry anvil start local
```

## 🐛 Troubleshooting

### "command not found: forge"
Install Foundry first using the command in Prerequisites.

### "package not found" errors
Run `go mod download` and `go mod tidy`.

### TUI not rendering correctly
Make sure your terminal supports colors and Unicode.
Try a modern terminal like:
- **iTerm2** (macOS)
- **Windows Terminal** (Windows)
- **Alacritty** (Linux/macOS/Windows)

### Dependencies not installing
```bash
# Clean and reinstall
go clean -modcache
go mod download
```

## 📚 Next Steps

1. **Create your first preset:**
   ```bash
   lazyfoundry anvil add local http://127.0.0.1:8545 31337
   ```

2. **Launch the TUI:**
   ```bash
   lazyfoundry
   ```

3. **Try building a contract:**
   - Navigate to Forge Mode
   - Select "Build Contracts"

## 🆘 Getting Help

- Check the [README](README.md)
- Open an [issue](https://github.com/jhaksh-24/Lazy-Foundry/issues)
- Read Foundry docs: https://book.getfoundry.sh/

---

Happy building! 🧰
