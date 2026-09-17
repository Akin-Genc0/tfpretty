package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/Akin-Genc0/tfpretty/internal/terraform"
)

type Screen int

const (
	PlanScreen Screen = iota
	Json
	Help
	Diff
)

type Model struct {
	Plan   terraform.Plan
	Cursor int
	Screen Screen
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	// Is it a key press
	case tea.KeyPressMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}

		case "down":
			if m.Cursor < len(m.Plan.ResourceChanges)-1 {
				m.Cursor++
			}

		case "j":
			m.Screen = Json

		case "esc":
			m.Screen = PlanScreen

		case "?":
			m.Screen = Help

		case "d":
			m.Screen = Diff
		}
	}

	return m, nil
}
