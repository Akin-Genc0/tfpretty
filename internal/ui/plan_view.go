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

// have helper function to list resources have the cursor pointer so it can move
func listResources(plan terraform.Plan, cursor int) {
	//loop over all reoases then check whether the cruse is eqaul to its index iof so and entr is pressed then take them ot the deatl screen
}

func renderPlanScreen(m Model) tea.View {
	// plan layout
	return tea.NewView("")
}
