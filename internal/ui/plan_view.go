package ui

import (
	"fmt"

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

func renderChangeSummary(counts CountChange) string {
	return fmt.Sprintf(
		"%s  %s  %s  %s",
		createStyle.Render(fmt.Sprintf("+  %d create", counts.CountCreate)),
		updateStyle.Render(fmt.Sprintf("~  %d update", counts.CountUpdate)),
		deleteStyle.Render(fmt.Sprintf("-  %d delete", counts.CountDelete)),
		replaceStyle.Render(fmt.Sprintf("+/-  %d replace", counts.CountReplace)),
	)
}

func listResources(plan terraform.Plan, cursor int) string {
	output := "NAME                                  TYPE\n"
	output += "───────────────────────────────────────────────────────────────\n"

	for index, resource := range plan.ResourceChanges {
		row := fmt.Sprintf("%-37s %s", resource.Name, resource.Type)

		if index == cursor {
			output += selectedStyle.Render(row) + "\n"
		} else {
			output += row + "\n"
		}
	}

	return output
}

func renderPlanScreen(m Model) tea.View {
	content := fmt.Sprintf(
		"tfpretty Terraform Plan\n\n%s\n\n%s%s",
		renderChangeSummary(countChanges(m.Plan)),
		listResources(m.Plan, m.Cursor),
		renderFooter(),
	)

	return tea.NewView(content)
}
