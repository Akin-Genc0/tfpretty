package ui

import (
	tea "charm.land/bubbletea/v2"
)

type CountChange struct {
	CountCreate  int
	CountUpdate  int
	CountDelete  int
	CountReplace int
}

func (m Model) View() tea.View {
	header := "tfpretty Terraform Plan"

	return tea.NewView(header)
}
