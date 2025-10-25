Spackage tui

import (
	"strings"
)

// View renders the entire UI based on current state
func (m Model) View() string {
	// If form is active, always show form (overrides screen)
	if m.form.Active {
		return m.renderForm()
	}
	
	// Otherwise show current screen
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
	case HelpScreen:
		return m.renderHelp()
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
		"📜 Run Script",
		"📥 Install Package",
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
	b.WriteString(renderTitle("Manage Anvil Presets"))
	b.WriteString("\n\n")
	
	// Menu items
	menuItems := []string{
		"➕ Add New Preset",
		"📋 View All Presets",
		"🗑️  Delete Preset",
		"⬅️  Back to Anvil Menu",
	}
	
	for i, item := range menuItems {
		b.WriteString(renderMenuItem(item, i == m.cursor))
		b.WriteString("\n")
	}
	
	b.WriteString("\n")
	b.WriteString(renderHelp("[↑↓/jk] Navigate  [Enter] Select  [Esc] Back"))
	
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

// renderHelp renders the help/documentation screen
func (m Model) renderHelp() string {
	var b strings.Builder
	
	b.WriteString(headerStyle.Render("ℹ️  Help & Documentation"))
	b.WriteString("\n\n")
	
	// Navigation section
	b.WriteString(titleStyle.Render("🎮 Navigation"))
	b.WriteString("\n\n")
	b.WriteString(menuItemStyle.Render("  ↑/↓ or j/k  - Navigate menu items"))
	b.WriteString("\n")
	b.WriteString(menuItemStyle.Render("  Enter/Space - Select item"))
	b.WriteString("\n")
	b.WriteString(menuItemStyle.Render("  Esc         - Go back"))
	b.WriteString("\n")
	b.WriteString(menuItemStyle.Render("  q           - Quit application"))
	b.WriteString("\n\n")
	
	// Forge section
	b.WriteString(titleStyle.Render("🔨 Forge Commands"))
	b.WriteString("\n\n")
	b.WriteString(menuItemStyle.Render("  Build      - Compile your smart contracts"))
	b.WriteString("\n")
	b.WriteString(menuItemStyle.Render("  Test       - Run your test suite"))
	b.WriteString("\n")
	b.WriteString(menuItemStyle.Render("  Init       - Initialize a new Foundry project"))
	b.WriteString("\n")
	b.WriteString(menuItemStyle.Render("  Coverage   - Generate test coverage report"))
	b.WriteString("\n")
	b.WriteString(menuItemStyle.Render("  Deploy     - Deploy contracts to network"))
	b.WriteString("\n")
	b.WriteString(menuItemStyle.Render("  Script     - Run deployment scripts"))
	b.WriteString("\n")
	b.WriteString(menuItemStyle.Render("  Install    - Install dependencies"))
	b.WriteString("\n\n")
	
	// Anvil section
	b.WriteString(titleStyle.Render("⚙️  Anvil Presets"))
	b.WriteString("\n\n")
	b.WriteString(menuItemStyle.Render("  Presets let you save network configurations"))
	b.WriteString("\n")
	b.WriteString(menuItemStyle.Render("  including RPC URLs, chain IDs, and fork URLs."))
	b.WriteString("\n\n")
	b.WriteString(menuItemStyle.Render("  • Add Preset    - Create new configuration"))
	b.WriteString("\n")
	b.WriteString(menuItemStyle.Render("  • List Presets  - View all saved presets"))
	b.WriteString("\n")
	b.WriteString(menuItemStyle.Render("  • Start Anvil   - Launch local node with preset"))
	b.WriteString("\n\n")
	
	// Forms section
	b.WriteString(titleStyle.Render("📝 Using Forms"))
	b.WriteString("\n\n")
	b.WriteString(menuItemStyle.Render("  Tab/Shift+Tab - Move between fields"))
	b.WriteString("\n")
	b.WriteString(menuItemStyle.Render("  Enter         - Submit form"))
	b.WriteString("\n")
	b.WriteString(menuItemStyle.Render("  Esc           - Cancel form"))
	b.WriteString("\n")
	b.WriteString(menuItemStyle.Render("  Backspace     - Delete character"))
	b.WriteString("\n")
	b.WriteString(menuItemStyle.Render("  Ctrl+U        - Clear current field"))
	b.WriteString("\n\n")
	
	// Tips section
	b.WriteString(titleStyle.Render("💡 Tips"))
	b.WriteString("\n\n")
	b.WriteString(menuItemStyle.Render("  • Fields marked (optional) can be left empty"))
	b.WriteString("\n")
	b.WriteString(menuItemStyle.Render("  • Private keys are shown as dots for security"))
	b.WriteString("\n")
	b.WriteString(menuItemStyle.Render("  • Use CLI for advanced flags and options"))
	b.WriteString("\n")
	b.WriteString(menuItemStyle.Render("  • Check output screen after each command"))
	b.WriteString("\n\n")
	
	// Footer
	b.WriteString(dimStyle.Render("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"))
	b.WriteString("\n")
	b.WriteString(dimStyle.Render("For more info: github.com/jhaksh-24/Lazy-Foundry"))
	b.WriteString("\n\n")
	b.WriteString(renderHelp("[Esc/Enter] Back to Main Menu"))
	
	return boxStyle.Render(b.String())
}

// renderLoading renders a loading state (for future async operations)
func (m Model) renderLoading() string {
	var b strings.Builder
	
	b.WriteString(headerStyle.Render("⏳ Executing..."))
	b.WriteString("\n\n")
	b.WriteString(dimStyle.Render("Please wait..."))
	b.WriteString("\n\n")
	b.WriteString(successStyle.Render("🔄 Working..."))
	
	return boxStyle.Render(b.String())
}tring(renderHelp("[↑↓/jk] Navigate  [Enter
