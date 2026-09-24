package ui

import (
	"fmt"
	"strings"

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
		content += "Changes\n"
		content += "───────────────────────────────────────────────────────────\n\n"
		content += "    ATTRIBUTE              BEFORE              AFTER\n\n"

		seen := make(map[string]bool)
		changed := 0
		for attribute, beforeValue := range resource.Change.Before {
			content += formatChange(attribute, beforeValue, resource.Change.After[attribute])
			seen[attribute] = true
			changed++
		}

		for attribute, afterValue := range resource.Change.After {
			if !seen[attribute] {
				content += fmt.Sprintf("+  %-38s %v\n", attribute, afterValue)
				changed++
			}
		}

		content += fmt.Sprintf("\n%d attributes changed\n", changed)
	}

	return content
}

func formatChange(attribute string, before, after any) string {
	return fmt.Sprintf("~  %-22s %-18v ->  %v\n", attribute, before, after)
}

func renderDetailScreen(m Model) tea.View {
	resource := m.Plan.ResourceChanges[m.Cursor]

	content := fmt.Sprintf(
		"Name    %s\nType    %s\n\n%s  %s\n\n%s%s",
		resource.Name,
		resource.Type,
		renderActionSymbol(resource.Action),
		strings.ToUpper(string(resource.Action)),
		resourceData(resource),
		renderFooter("esc back       r raw       ? help       q quit"),
	)

	return tea.NewView(content)
}
