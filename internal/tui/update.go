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
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
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
		case 0: //Forge
			m.currentScreen = ForgeScreen
			m.cursor = 0
		case 1:  //Anvil
			m.currentScreen = AnvilScreen
			m.cursor = 0
		case 2:
			// Will show help screen
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
		if m.cursor < 5 {
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
		case 4: // Deploy (needs input)
			// TODO: Add input mode for contract name
			m.output = "Deploy feature coming soon - use CLI for now"
			m.currentScreen = ExecutingScreen
		case 5: // Back
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
		if m.cursor < 3 { // 4 menu items
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
	case "esc":
		m.currentScreen = AnvilScreen
		m.cursor = 0
	}
	return m, nil
}

func (m Model) handleExecutingInput(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "enter", "esc", " ":
		m.currentScreen = ForgeScreen // Default back to forge
		m.cursor = 0
		m.output = ""
		m.lastError = nil
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
	
	return m, nil
}
