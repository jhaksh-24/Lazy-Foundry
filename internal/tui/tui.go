package tui

import (
	"fmt"
	
	tea "github.com/charmbracelet/bubbletea"
)

func Run() error {
	m := NewModel()
	
	p := tea.NewProgram(
		m,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)
	
	if _, err := p.Run(); err != nil {
		return fmt.Errorf("error running TUI: %w", err)
	}
	
	return nil
}
