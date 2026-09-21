package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/Akin-Genc0/tfpretty/internal/terraform"
)

type CountChange struct {
	CountCreate  int
	CountUpdate  int
	CountDelete  int
	CountReplace int
}

func countChanges(plan terraform.Plan) CountChange {
	counts := CountChange{}

	for _, resource := range plan.ResourceChanges {
		switch resource.Action {
		case terraform.ActionCreate:
			counts.CountCreate++
		case terraform.ActionUpdate:
			counts.CountUpdate++
		case terraform.ActionDelete:
			counts.CountDelete++
		case terraform.ActionReplace:
			counts.CountReplace++
		}
	}

	return counts
}

func (m Model) View() tea.View {
	header := "tfpretty Terraform Plan"

	return tea.NewView(header)
}
