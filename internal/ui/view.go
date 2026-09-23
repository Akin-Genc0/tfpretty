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

func (m Model) View() tea.View {
	switch m.Screen {
	case Help:
		return renderHelpScreen(m)
	case PlanScreen:
		return renderPlanScreen(m)
	default:
		return renderPlanScreen(m)
	}
}
