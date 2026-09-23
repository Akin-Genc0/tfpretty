package ui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/Akin-Genc0/tfpretty/internal/terraform"
)

func resourceData(resource terraform.ResourceChange) string {
	var content string
	switch resource.Action {
	case terraform.ActionCreate:
		for attribute, value := range resource.Change.After {
			content += fmt.Sprintf("+  %-38s %v\n", attribute, value)
		}
	case terraform.ActionDelete:
		for attribute, value := range resource.Change.Before {
			content += fmt.Sprintf("-  %-38s %v\n", attribute, value)
		}

	default:
		seen := make(map[string]bool)
		for attribute, beforeValue := range resource.Change.Before {
			content += fmt.Sprintf("~  %-38s %v -> %v\n", attribute, beforeValue, resource.Change.After[attribute])
			seen[attribute] = true
		}

		for attribute, afterValue := range resource.Change.After {
			if !seen[attribute] {
				content += fmt.Sprintf("+  %-38s %v\n", attribute, afterValue)
			}
		}
	}

	return content
}

//helper function to render the row ui

func renderDetailScreen(m Model) tea.View {
	//render header part
	//render resoresdata stuff
	//render footer part
	return tea.NewView("")
}
