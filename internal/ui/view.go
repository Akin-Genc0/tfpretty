package ui

import (
	tea "charm.land/bubbletea/v2"
)

func renderFooter() string {
	return `
───────────────────────────────────────────────────────────────
↑↓ navigate       enter details       ? help       q quit
`
}

//helper function for the bottom screen render
//helper for rendering the diffrent screens

func (m Model) View() tea.View {
	switch m.Screen {
	case PlanScreen:
		return renderPlanScreen(m)
	default:
		return renderPlanScreen(m)
	}
}
