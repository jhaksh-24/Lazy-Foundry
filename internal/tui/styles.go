package tui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// Primary colors
	primaryColor   = lipgloss.Color("#00D9FF")  // Cyan
	secondaryColor = lipgloss.Color("#FF6B9D")  // Pink
	accentColor    = lipgloss.Color("#C6FF00")  // Lime
	
	// UI colors
	backgroundColor = lipgloss.Color("#1a1b26")
	textColor       = lipgloss.Color("#c0caf5")
	dimColor        = lipgloss.Color("#565f89")
	successColor    = lipgloss.Color("#9ece6a")
	errorColor      = lipgloss.Color("#f7768e")
	warningColor    = lipgloss.Color("#e0af68")
)

var (
	// Header style - top banner
	headerStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(primaryColor).
		Background(backgroundColor).
		Padding(1, 2).
		Align(lipgloss.Center)
	
	// Title style - section titles
	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(secondaryColor).
		Padding(0, 1)
	
	// Menu item styles
	menuItemStyle = lipgloss.NewStyle().
		Foreground(textColor).
		Padding(0, 2)
	
	selectedMenuItemStyle = lipgloss.NewStyle().
		Foreground(backgroundColor).
		Background(primaryColor).
		Bold(true).
		Padding(0, 2)
	
	// Box styles
	boxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(primaryColor).
		Padding(1, 2).
		Margin(1, 2)
	
	// Status message styles
	successStyle = lipgloss.NewStyle().
		Foreground(successColor).
		Bold(true)
	
	errorStyle = lipgloss.NewStyle().
		Foreground(errorColor).
		Bold(true)
	
	dimStyle = lipgloss.NewStyle().
		Foreground(dimColor)
	
	// Help text at bottom
	helpStyle = lipgloss.NewStyle().
		Foreground(dimColor).
		Italic(true).
		Align(lipgloss.Center).
		Padding(1, 0)

	// Output/execution styles	
	outputStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(accentColor).
		Padding(1, 2).
		Margin(1, 2).
		Foreground(textColor)
)

func renderTitle(text string) string {
	return titleStyle.Render(text)
}

func renderMenuItem(text string, selected bool) string {
	if selected {
		return selectedMenuItemStyle.Render("▶ " + text)
	}
	return menuItemStyle.Render("  " + text)
}

func renderBox(content string) string {
	return boxStyle.Render(content)
}

func renderSuccess(text string) string {
	return successStyle.Render("✓ " + text)
}

func renderError(text string) string {
	return errorStyle.Render("✗ " + text)
}

func renderHelp(text string) string {
	return helpStyle.Render(text)
}
