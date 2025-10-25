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
	HelpScreen
	// Form screens
	PresetFormScreen
	DeployFormScreen
	ScriptFormScreen
	InstallFormScreen
)

// InputField represents a single field in a form
type InputField struct {
	Label       string
	Value       string
	Placeholder string
	Secure      bool   // For password/private key fields
	Optional    bool   // Whether field can be empty
	Help        string // Help text shown below field
}

// FormState manages multi-field input forms
type FormState struct {
	Active       bool
	Title        string
	Fields       []InputField
	CurrentField int
	Submitted    bool
	OnSubmit     func([]string) (Model, tea.Cmd)
	OnCancel     func() (Model, tea.Cmd)
	ReturnScreen Screen // Which screen to return to on cancel
}

// Model is the main TUI state
type Model struct {
	currentScreen Screen
	cursor        int
	executing     bool
	output        string
	lastError     error
	width         int
	height        int
	
	// Form state
	form FormState
	
	// Preset management
	presetNames   []string
	selectedPreset string
}

func NewModel() Model {
	return Model{
		currentScreen: HomeScreen,
		cursor:        0,
		executing:     false,
		form: FormState{
			Active: false,
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

// Helper methods for form management
func (m *Model) startForm(title string, fields []InputField, returnScreen Screen, onSubmit func([]string) (Model, tea.Cmd)) {
	m.form = FormState{
		Active:       true,
		Title:        title,
		Fields:       fields,
		CurrentField: 0,
		OnSubmit:     onSubmit,
		ReturnScreen: returnScreen,
	}
	m.currentScreen = PresetFormScreen // Generic form screen
}

func (m *Model) cancelForm() (Model, tea.Cmd) {
	m.currentScreen = m.form.ReturnScreen
	m.form = FormState{Active: false}
	m.cursor = 0
	return *m, nil
}

func (m *Model) submitForm() (Model, tea.Cmd) {
	// Collect all field values
	values := make([]string, len(m.form.Fields))
	for i, field := range m.form.Fields {
		values[i] = field.Value
	}
	
	// Clear form state
	m.form.Active = false
	
	// Call the submit handler
	if m.form.OnSubmit != nil {
		return m.form.OnSubmit(values)
	}
	
	return *m, nil
}

func (m *Model) nextField() {
	if m.form.CurrentField < len(m.form.Fields)-1 {
		m.form.CurrentField++
	}
}

func (m *Model) prevField() {
	if m.form.CurrentField > 0 {
		m.form.CurrentField--
	}
}

func (m *Model) currentFieldValue() string {
	if m.form.CurrentField < len(m.form.Fields) {
		return m.form.Fields[m.form.CurrentField].Value
	}
	return ""
}

func (m *Model) setCurrentFieldValue(value string) {
	if m.form.CurrentField < len(m.form.Fields) {
		m.form.Fields[m.form.CurrentField].Value = value
	}
}

func (m *Model) deleteCharFromCurrentField() {
	value := m.currentFieldValue()
	if len(value) > 0 {
		m.setCurrentFieldValue(value[:len(value)-1])
	}
}

func (m *Model) addCharToCurrentField(char rune) {
	value := m.currentFieldValue()
	m.setCurrentFieldValue(value + string(char))
}
