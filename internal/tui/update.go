package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	
	case tea.KeyMsg:
		// If in form, handle form input
		if m.form.Active {
			return m.handleFormInput(msg)
		}
		
		// Global shortcuts
		switch msg.String() {
		case "ctrl+c", "q":
			if m.currentScreen == HomeScreen {
				return m, tea.Quit
			}
			// Ask for confirmation on other screens
		}
		
		return m.handleScreenInput(msg)
	}
	
	return m, nil
}

func (m Model) handleScreenInput(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch m.currentScreen {
	case HomeScreen:
		return m.handleHomeInput(msg)
	case ForgeScreen:
		return m.handleForgeInput(msg)
	case AnvilScreen:
		return m.handleAnvilInput(msg)
	case PresetsScreen:
		return m.handlePresetsInput(msg)
	case ExecutingScreen:
		return m.handleExecutingInput(msg)
	case HelpScreen:
		return m.handleHelpInput(msg)
	default:
		return m, nil
	}
}

func (m Model) handleHomeInput(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < 3 {
			m.cursor++
		}
	case "enter", " ":
		switch m.cursor {
		case 0: // Forge
			m.currentScreen = ForgeScreen
			m.cursor = 0
		case 1: // Anvil
			m.currentScreen = AnvilScreen
			m.cursor = 0
		case 2: // Help
			m.currentScreen = HelpScreen
			m.cursor = 0
		case 3: // Exit
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) handleForgeInput(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < 7 { // Updated count
			m.cursor++
		}
	case "enter", " ":
		switch m.cursor {
		case 0: // Build
			return m.executeForgeCommand("build")
		case 1: // Test
			return m.executeForgeCommand("test")
		case 2: // Init
			return m.executeForgeCommand("init")
		case 3: // Coverage
			return m.executeForgeCommand("coverage")
		case 4: // Deploy - OPEN FORM
			m.createDeployForm()
		case 5: // Script - OPEN FORM
			m.createScriptForm()
		case 6: // Install - OPEN FORM
			m.createInstallForm()
		case 7: // Back
			m.currentScreen = HomeScreen
			m.cursor = 0
		}
	case "esc":
		m.currentScreen = HomeScreen
		m.cursor = 0
	}
	return m, nil
}

func (m Model) handleAnvilInput(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < 3 {
			m.cursor++
		}
	case "enter", " ":
		switch m.cursor {
		case 0: // List presets
			return m.executeAnvilCommand("list")
		case 1: // Start anvil
			return m.executeAnvilCommand("start", "local")
		case 2: // Manage presets
			m.currentScreen = PresetsScreen
			m.cursor = 0
		case 3: // Back
			m.currentScreen = HomeScreen
			m.cursor = 0
		}
	case "esc":
		m.currentScreen = HomeScreen
		m.cursor = 0
	}
	return m, nil
}

func (m Model) handlePresetsInput(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < 3 {
			m.cursor++
		}
	case "enter", " ":
		switch m.cursor {
		case 0: // Add preset - OPEN FORM
			m.createAddPresetForm()
		case 1: // View presets
			return m.executeAnvilCommand("list")
		case 2: // Delete preset (would need another form for selection)
			m.output = "Select a preset to delete (feature coming soon)"
			m.currentScreen = ExecutingScreen
		case 3: // Back
			m.currentScreen = AnvilScreen
			m.cursor = 0
		}
	case "esc":
		m.currentScreen = AnvilScreen
		m.cursor = 0
	}
	return m, nil
}

func (m Model) handleExecutingInput(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "enter", "esc", " ":
		// Return to the previous screen
		if m.form.ReturnScreen != 0 {
			m.currentScreen = m.form.ReturnScreen
		} else {
			m.currentScreen = ForgeScreen
		}
		m.cursor = 0
		m.output = ""
		m.lastError = nil
	}
	return m, nil
}

func (m Model) handleHelpInput(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "esc", "q", "enter":
		m.currentScreen = HomeScreen
		m.cursor = 0
	}
	return m, nil
}

// handleFormInput processes keyboard input when a form is active
func (m Model) handleFormInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		// Cancel form
		return m.cancelForm()
		
	case "tab":
		// Next field
		m.nextField()
		
	case "shift+tab":
		// Previous field
		m.prevField()
		
	case "enter":
		// Try to submit
		if err := m.validateForm(); err != nil {
			// Show validation error (could improve this)
			m.form.Fields[m.form.CurrentField].Help = "⚠️  " + err.Error()
			return m, nil
		}
		return m.submitForm()
		
	case "backspace":
		// Delete character
		m.deleteCharFromCurrentField()
		
	case "ctrl+u":
		// Clear current field
		m.setCurrentFieldValue("")
		
	default:
		// Add character to current field
		if len(msg.String()) == 1 {
			m.addCharToCurrentField(rune(msg.String()[0]))
		}
	}
	
	return m, nil
}

func (m Model) executeForgeCommand(command string, args ...string) (Model, tea.Cmd) {
	m.currentScreen = ExecutingScreen
	m.executing = true
	
	output, err := executeForge(command, args...)
	
	if err != nil {
		m.lastError = err
		m.output = output
	} else {
		m.lastError = nil
		m.output = output
		if m.output == "" {
			m.output = "Command executed successfully!"
		}
	}
	
	m.form.ReturnScreen = ForgeScreen
	return m, nil
}

func (m Model) executeAnvilCommand(command string, args ...string) (Model, tea.Cmd) {
	m.currentScreen = ExecutingScreen
	m.executing = true
	
	output, err := executeAnvil(command, args...)
	
	if err != nil {
		m.lastError = err
		m.output = output
	} else {
		m.lastError = nil
		m.output = output
	}
	
	m.form.ReturnScreen = AnvilScreen
	return m, nil
}
