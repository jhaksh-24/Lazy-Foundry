package tui

import (
	"fmt"
	"strings"
)

// View renders the entire UI based on current state
// This is called every time something changes
func (m Model) View() string {
	// Build the view based on current screen
	switch m.currentScreen {
	case HomeScreen:
		return m.renderHome()
	case ForgeScreen:
		return m.renderForge()
	case AnvilScreen:
		return m.renderAnvil()
	case PresetsScreen:
		return m.renderPresets()
	case ExecutingScreen:
		return m.renderExecuting()
	default:
		return "Unknown screen"
	}
}

// renderHome renders the main menu
func (m Model) renderHome() string {
	var b strings.Builder
	
	// Header
	b.WriteString(headerStyle.Render("🧰 Lazy Foundry"))
	b.WriteString("\n")
	b.WriteString(dimStyle.Render("Simplified Foundry Workflow Tool"))
	b.WriteString("\n\n")
	
	// Menu title
	b.WriteString(renderTitle("Select Mode"))
	b.WriteString("\n\n")
	
	// Menu items
	menuItems := []string{
		"🔨 Forge - Build & Deploy Contracts",
		"⚙️  Anvil - Manage Presets & Local Node",
		"ℹ️  Help & Documentation",
		"👋 Exit",
	}
	
	for i, item := range menuItems {
		b.WriteString(renderMenuItem(item, i == m.cursor))
		b.WriteString("\n")
	}
	
	b.WriteString("\n")
	b.WriteString(renderHelp("[↑↓/jk] Navigate  [Enter] Select  [q] Quit"))
	
	return boxStyle.Render(b.String())
}

// renderForge renders the forge menu
func (m Model) renderForge() string {
	var b strings.Builder
	
	// Header
	b.WriteString(headerStyle.Render("🔨 Forge Mode"))
	b.WriteString("\n\n")
	b.WriteString(renderTitle("What would you like to do?"))
	b.WriteString("\n\n")
	
	// Menu items
	menuItems := []string{
		"🏗️  Build Contracts",
		"🧪 Run Tests",
		"📦 Initialize New Project",
		"📊 Generate Coverage Report",
		"🚀 Deploy Contract",
		"⬅️  Back to Main Menu",
	}
	
	for i, item := range menuItems {
		b.WriteString(renderMenuItem(item, i == m.cursor))
		b.WriteString("\n")
	}
	
	b.WriteString("\n")
	b.WriteString(renderHelp("[↑↓/jk] Navigate  [Enter] Select  [Esc] Back  [q] Quit"))
	
	return boxStyle.Render(b.String())
}

// renderAnvil renders the anvil menu
func (m Model) renderAnvil() string {
	var b strings.Builder
	
	// Header
	b.WriteString(headerStyle.Render("⚙️  Anvil Mode"))
	b.WriteString("\n\n")
	b.WriteString(renderTitle("Environment & Presets"))
	b.WriteString("\n\n")
	
	// Menu items
	menuItems := []string{
		"📋 List All Presets",
		"🚀 Start Local Anvil Node",
		"⚡ Manage Presets",
		"⬅️  Back to Main Menu",
	}
	
	for i, item := range menuItems {
		b.WriteString(renderMenuItem(item, i == m.cursor))
		b.WriteString("\n")
	}
	
	b.WriteString("\n")
	b.WriteString(renderHelp("[↑↓/jk] Navigate  [Enter] Select  [Esc] Back  [q] Quit"))
	
	return boxStyle.Render(b.String())
}

// renderPresets renders the preset management screen
func (m Model) renderPresets() string {
	var b strings.Builder
	
	b.WriteString(headerStyle.Render("⚡ Preset Management"))
	b.WriteString("\n\n")
	b.WriteString(dimStyle.Render("Use CLI for now to manage presets:"))
	b.WriteString("\n\n")
	b.WriteString(menuItemStyle.Render("  • lazyfoundry anvil add <n> <rpc> <chain-id>"))
	b.WriteString("\n")
	b.WriteString(menuItemStyle.Render("  • lazyfoundry anvil show <n>"))
	b.WriteString("\n")
	b.WriteString(menuItemStyle.Render("  • lazyfoundry anvil delete <n>"))
	b.WriteString("\n\n")
	b.WriteString(renderHelp("[Esc] Back"))
	
	return boxStyle.Render(b.String())
}

// renderExecuting renders command execution and output
func (m Model) renderExecuting() string {
	var b strings.Builder
	
	// Header
	if m.lastError != nil {
		b.WriteString(headerStyle.Render("❌ Execution Failed"))
		b.WriteString("\n\n")
		b.WriteString(renderError(m.lastError.Error()))
	} else {
		b.WriteString(headerStyle.Render("✓ Execution Complete"))
		b.WriteString("\n\n")
	}
	
	// Output
	if m.output != "" {
		b.WriteString("\n")
		b.WriteString(outputStyle.Render(m.output))
	}
	
	b.WriteString("\n\n")
	b.WriteString(renderHelp("[Enter] Continue  [q] Quit"))
	
	return b.String()
}

// renderLoading renders a loading state
func (m Model) renderLoading() string {
	var b strings.Builder
	
	b.WriteString(headerStyle.Render("⏳ Executing..."))
	b.WriteString("\n\n")
	b.WriteString(dimStyle.Render("Please wait..."))
	
	return boxStyle.Render(b.String())
}

// Helper function to format key bindings
func formatKeyBinding(key, description string) string {
	return fmt.Sprintf("%s %s", 
		successStyle.Render("["+key+"]"),
		dimStyle.Render(description))
}
