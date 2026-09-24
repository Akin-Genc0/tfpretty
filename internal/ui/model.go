package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/Akin-Genc0/tfpretty/internal/terraform"
)

type Screen int

const (
	PlanScreen Screen = iota
	DetailScreen
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

	case tea.KeyPressMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up":
			if m.Screen == DetailScreen {
				return m, nil
			}
			if m.Cursor > 0 {
				m.Cursor--
			}

		case "down":
			if m.Screen == DetailScreen {
				return m, nil
			}
			if m.Cursor < len(m.Plan.ResourceChanges)-1 {
				m.Cursor++
			}

		case "esc":
			m.Screen = PlanScreen
			return m, tea.ClearScreen

		case "?":
			m.Screen = Help
			return m, tea.ClearScreen

		case "r":
			m.Screen = Diff
			return m, tea.ClearScreen

		case "enter":
			m.Screen = DetailScreen
			return m, tea.ClearScreen
		}

	}

	return m, nil
}
