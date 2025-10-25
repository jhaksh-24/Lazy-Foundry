package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Screen int

const (
	HomeScreen Screen = iota
	ForgeScreen
	AnvilScreen
	PresetsScreen
	ExecutingScreen
)

type Model struct {
	currentScreen Screen
	cursor        int
	input         string
	inputMode     bool
	inputPrompt   string
	executing     bool
	output        string
	lastError     error
	presetNames   []string
	width         int
	height        int
}

func NewModel() Model {
	return Model{
		currentScreen: HomeScreen,
		cursor:        0,
		executing:     false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
