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

	const visibleResources = 6
	start := 0
	if cursor >= visibleResources {
		start = cursor - visibleResources + 1
	}

	end := start + visibleResources
	if end > len(plan.ResourceChanges) {
		end = len(plan.ResourceChanges)
	}

	if start > 0 {
		output += fmt.Sprintf("... %d more above\n", start)
	}

	for index := start; index < end; index++ {
		resource := plan.ResourceChanges[index]
		row := fmt.Sprintf("%-37s %s", resource.Name, resource.Type)

		if index == cursor {
			output += selectedStyle.Render(row) + "\n"
		} else {
			output += row + "\n"
		}
	}

	if end < len(plan.ResourceChanges) {
		output += fmt.Sprintf("... %d more below\n", len(plan.ResourceChanges)-end)
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
